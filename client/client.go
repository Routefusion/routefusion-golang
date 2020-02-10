package client

import (
	"io"
	"net/http"
	"time"
)

const (
	maxDefaultRetries = 3
)

// These values are derived from the default values of DefaultTransport and
// Transport respectively from net/http/transport.go
var (
	defaultRequestTimeout      = 30 * time.Second
	defaultTLSHandshakeTimeout = 10 * time.Second
	defaultMaxIdleConns        = 100
	defaultMaxIdleConnsPerHost = 2
	defaultIdleConnTimeout     = 90 * time.Second
)

// Client implements the base client request and response handling
// used by all service clients.
type Client struct {
	Authorizer Authorizer
	Retryer    Retryer

	httpClient *http.Client
	baseURL    string
}

// NewClient returns a new instance of sdk.Client.
// The config gets sanitized by returning a copy of the passed configuration
// with default set if the given values are not set by the caller.
// The default values are derived from the default values of DefaultTransport
// and Transport respectively from net/http/transport.go
func NewClient(config Config) *Client {
	config = sanitize(config)

	httpClient := newHTTPClient(config)

	client := &Client{
		baseURL:    config.BaseURL,
		httpClient: httpClient,
		Retryer:    config.Retryer,
		Authorizer: config.Authorizer,
	}

	return client
}

// sanitize is some space shuttle code that enables configuration to be
// easy.
func sanitize(config Config) Config {
	sanitized := Config{
		Authorizer: config.Authorizer,
		BaseURL:    config.BaseURL,
	}
	sanitized.Retryer = config.Retryer
	if config.Retryer == nil {
		sanitized.Retryer = DefaultRetryer{NumMaxRetries: maxDefaultRetries}
	}

	sanitized.RequestTimeout = &defaultRequestTimeout
	if config.RequestTimeout != nil {
		sanitized.RequestTimeout = config.RequestTimeout
	}
	sanitized.TLSHandshakeTimeout = &defaultTLSHandshakeTimeout
	if config.TLSHandshakeTimeout != nil {
		sanitized.TLSHandshakeTimeout = config.TLSHandshakeTimeout
	}
	sanitized.MaxIdleConns = &defaultMaxIdleConns
	if config.MaxIdleConns != nil {
		sanitized.MaxIdleConns = config.MaxIdleConns
	}
	sanitized.MaxIdleConnsPerHost = &defaultMaxIdleConnsPerHost
	if config.MaxIdleConnsPerHost != nil {
		sanitized.MaxIdleConnsPerHost = config.MaxIdleConnsPerHost
	}
	sanitized.IdleConnTimeout = &defaultIdleConnTimeout
	if config.IdleConnTimeout != nil {
		sanitized.IdleConnTimeout = config.IdleConnTimeout
	}
	return sanitized
}

// newHTTPClient creates a HTTP client based on go's http.DefaultClient and
// http.DefaultTransport.
func newHTTPClient(config Config) *http.Client {
	defaultTransport, _ := http.DefaultTransport.(*http.Transport)

	defaultTransport.TLSHandshakeTimeout = *config.TLSHandshakeTimeout
	defaultTransport.MaxIdleConns = *config.MaxIdleConns
	defaultTransport.MaxIdleConnsPerHost = *config.MaxIdleConnsPerHost
	defaultTransport.IdleConnTimeout = *config.IdleConnTimeout

	client := http.DefaultClient
	client.Transport = defaultTransport
	client.Timeout = *config.RequestTimeout
	return client
}

// NewRequest is to get a new Request option tied to a client
func (c *Client) NewRequest(op Operation, output interface{},
	body io.ReadSeeker, paramsList ...map[string]string) (*Request, error) {
	var params map[string]string
	if len(paramsList) == 0 {
		params = nil
	} else {
		params = paramsList[0]
	}
	return NewRequest(c.httpClient, c.Retryer, c.Authorizer, c.baseURL, op,
		output, body, params)
}

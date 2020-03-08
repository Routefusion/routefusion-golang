//go:generate go run internal/cmd/genmethods/generator.go

// Package routefusion implements a go sdk for routefusion services.
package routefusion

import (
	"net/http"
	"time"

	"github.com/routefusion/routefusion-golang/client"
)

const (
	baseURL = "https://api.routefusion.co/"
	v1      = "v1"
)

// Routefusion is the master struct that implements routefusion.Client.
type Routefusion struct {
	Client
	cl      *client.Client
	baseURL string
}

// Config is an optional control struct you can give the client. If config
// is nil, the client defaults to preset constants.
type Config struct {
	URL string

	// used to fine-tune the underlying transport of the HTTP client.
	// see client/client.go and godoc for preset constants.
	RequestTimeout      *time.Duration
	TLSHandshakeTimeout *time.Duration
	MaxIdleConns        *int
	MaxIdleConnsPerHost *int
	IdleConnTimeout     *time.Duration
}

// New gives an instantiated and preset version of Routefusion.
func New(cfg Config) *Routefusion {
	url := cfg.URL
	if url == "" {
		url = baseURL
	}
	return &Routefusion{
		cl: client.NewClient(client.Config{
			BaseURL:             baseURL,
			RequestTimeout:      cfg.RequestTimeout,
			TLSHandshakeTimeout: cfg.TLSHandshakeTimeout,
			MaxIdleConns:        cfg.MaxIdleConns,
			MaxIdleConnsPerHost: cfg.MaxIdleConnsPerHost,
			IdleConnTimeout:     cfg.IdleConnTimeout,
		}),
	}
}

func (r *Routefusion) GetUser() (*UserDetails, error) {
	var output = &UserDetails{}
	op := client.Operation{
		HTTPPath:   r.baseURL + v1 + "/users/me",
		HTTPMethod: http.MethodGet,
	}

	req, err := r.cl.NewRequest(op, output, nil)
	if err != nil {
		return nil, err
	}

	if err := req.Send(); err != nil {
		return nil, err
	}

	return output, nil
}

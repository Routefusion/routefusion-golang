package client

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"path"
	"sync"
	"time"
)

const (
	requestHeaderKeyAccept = "Accept"
)

// A Request is the service request to be made.
type Request struct {
	sync.Mutex

	HTTPRequest  *http.Request
	HTTPResponse *http.Response
	Error        RequestFailureError
	Output       interface{}
	Retryer      Retryer
	RetryCount   int

	body   io.ReadSeeker
	client *http.Client
}

// An Operation is the service API operation to be made
type Operation struct {
	HTTPMethod string
	HTTPPath   string
}

// NewRequest returns a new request. It is intended to be a shoot once and forget
// object. While Send() is threadsafe, multiple calls in goroutines
// for a single Request will not make sense because it retries inherently.
func NewRequest(client *http.Client,
	retryer Retryer,
	authorizer Authorizer,
	baseURL string,
	op Operation,
	output interface{},
	body io.ReadSeeker,
	params map[string]string) (*Request, error) {

	finalURL, err := url.Parse(baseURL)
	if err != nil {
		return nil, fmt.Errorf("invalid endpoint or HTTPPath supplied: %s", err)
	}

	finalURL.Path = path.Join(finalURL.Path, op.HTTPPath)
	httpReq, err := http.NewRequest(op.HTTPMethod, finalURL.String(), nil)
	if err != nil {
		return nil, fmt.Errorf("error making new request: %s", err)
	}
	if authorizer != nil {
		authorizer.AuthorizeRequest(httpReq)
	}
	unpackParams(httpReq, params)

	return &Request{
		Output:      output,
		HTTPRequest: httpReq,
		body:        body,
		Retryer:     retryer,
		client:      client,
	}, nil
}

// Send makes the actual request. It's got a retryer built in that
// employs customizable retry logic for a configurable finite number
// of attempts.
// An error message of format `http request failed after 0 attempts: ...` means
// that at least one regular HTTP request but no (zero) further attempts were made.
func (r *Request) Send() (err error) {
	r.Lock()
	defer r.Unlock()

	for try := 0; ; try++ {
		if r.body != nil {
			r.HTTPRequest.Body = newOffsetReader(r.body, 0)
		}

		r.HTTPResponse, err = r.client.Do(r.HTTPRequest)

		if r.Retryer.ShouldRetry(r) && try < r.Retryer.MaxRetries() {
			r.RetryCount++
			time.Sleep(r.Retryer.RetryRules(r))
			continue
		}

		if err != nil {
			code := ErrCodeUndefined
			msg := fmt.Sprintf("http request failed after %d attempts", try)
			if uerr, ok := err.(*url.Error); ok && uerr.Timeout() {
				code = ErrCodeTimeout
			}
			return NewRequestFailureError(NewRFError(code, msg, err), 0, "")
		}

		if r.HTTPResponse.StatusCode < http.StatusOK ||
			r.HTTPResponse.StatusCode > http.StatusIMUsed {
			msg := fmt.Sprintf("http request failed after %d attempts", try)
			if r.HTTPResponse.StatusCode == http.StatusNotFound {
				if r.HTTPResponse != nil {
					r.HTTPResponse.Body.Close()
				}
				return NewRequestFailureError(NewRFError(
					ErrCodeNotFound, msg, nil),
					r.HTTPResponse.StatusCode,
					"")
			}

			return NewRequestFailureError(
				NewRFError(ErrCodeUndefined, msg, err),
				r.HTTPResponse.StatusCode,
				"")
		}

		if r.Output != nil {
			err = r.unmarshalBody()
			if err != nil {
				return NewRequestFailureError(
					NewRFError(ErrCodeUnmarshalFailed, "unmarshal failed", err),
					r.HTTPResponse.StatusCode,
					"")
			}
		}

		return nil
	}
}

func (r *Request) readBody() ([]byte, error) {
	p, err := ioutil.ReadAll(r.HTTPResponse.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response body: %s", err)
	}
	defer r.HTTPResponse.Body.Close()
	return p, nil
}

func (r *Request) unmarshalBody() error {
	defer r.HTTPResponse.Body.Close()

	acceptHeader := r.HTTPRequest.Header.Get(requestHeaderKeyAccept)

	unmarshaler, ok := unmarshalers[acceptHeader]
	if ok {
		return unmarshaler.Unmarshal(r.HTTPResponse.Body, r.Output)
	}

	return json.NewDecoder(r.HTTPResponse.Body).Decode(r.Output)
}

func unpackParams(r *http.Request, params map[string]string) {
	q := r.URL.Query()
	for paramName, paramValue := range params {
		q.Add(paramName, paramValue)
	}
	r.URL.RawQuery = q.Encode()
}

// MergeRequestHeader merges the given headers but ignores multi-values and
// therefore overwrites the slice of header values with the new slice.
func MergeRequestHeader(h1, h2 http.Header) http.Header {
	headers := make(http.Header, len(h1))
	for k, v1 := range h1 {
		headers[k] = v1
		if v2, ok := h2[k]; ok {
			headers[k] = v2
		}
	}

	for k, v2 := range h2 {
		headers[k] = v2
	}

	return headers
}

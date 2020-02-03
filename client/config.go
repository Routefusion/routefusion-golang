package client

import (
	"time"
)

// Config holds any all extensible and configurable features.
type Config struct {
	Retryer    Retryer
	Authorizer Authorizer
	BaseURL    string

	// used to fine-tune the underlying transport of the HTTP client.
	RequestTimeout      *time.Duration
	TLSHandshakeTimeout *time.Duration
	MaxIdleConns        *int
	MaxIdleConnsPerHost *int
	IdleConnTimeout     *time.Duration
}

// Int is a helper to set the pointer based config values above.
func Int(i int) *int {
	return &i
}

// Duration is a helper to set the pointer based config values above.
func Duration(t time.Duration) *time.Duration {
	return &t
}

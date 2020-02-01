package client

import (
	"fmt"
	"net/http"
)

const (
	bearerTokenAuthorization = "Authorization"
)

// Authorizer defines a way to retrieve Credentials
type Authorizer interface {
	// AuthorizeRequest is an implementation provided to make
	// credentials signing pluggable and customizable.
	//
	// A default bearer token implementation is provided.
	AuthorizeRequest(r *http.Request)
}

// BearerTokenAuthorizer represents the basic inputs for token authorization.
type BearerTokenAuthorizer struct {
	Token string
}

// AuthorizeRequest sets a  pre-configured token value on the request header.
func (b *BearerTokenAuthorizer) AuthorizeRequest(r *http.Request) {
	if b.Token != "" {
		r.Header.Set(bearerTokenAuthorization, fmt.Sprintf("Bearer %s", b.Token))
	}
}

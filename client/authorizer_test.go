package client

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_AuthorizeRequest(t *testing.T) {
	b := &BearerTokenAuthorizer{
		Token: "password",
	}

	assert := assert.New(t)
	r, _ := http.NewRequest("GET", "", nil)
	b.AuthorizeRequest(r)
	token := r.Header.Get(bearerTokenAuthorization)
	assert.Equal(token, "Bearer password")
}

package routefusion

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha512"
	"encoding/base64"
	"io/ioutil"
	"net/http"
)

const (
	clientID  = "client-id"
	signature = "signature"
)

// ClientAuthorizer is built to the specifications of https://developer.routefusion.co/#authentication
type ClientAuthorizer struct {
	ClientID  string
	SecretKey string
}

// AuthorizeRequest sets a  pre-configured token value on the request header.
func (c *ClientAuthorizer) AuthorizeRequest(r *http.Request) {
	mac := hmac.New(sha512.New, []byte(c.SecretKey))
	if r.Body == nil {
		mac.Write([]byte(r.URL.Path))
	} else {
		bodyBytes, _ := ioutil.ReadAll(r.Body)
		r.Body.Close()
		mac.Write(bodyBytes)
		r.Body = ioutil.NopCloser(bytes.NewBuffer(bodyBytes))
	}
	h := mac.Sum(nil)
	encoded := base64.StdEncoding.EncodeToString(h)
	r.Header.Set(clientID, c.ClientID)
	r.Header.Set(signature, string(encoded))
}

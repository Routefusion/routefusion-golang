package client

import (
	"net/http"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

type mockRetryer struct{}

func (m *mockRetryer) RetryRules(*Request) time.Duration {
	return 1 * time.Second
}

func (m *mockRetryer) ShouldRetry(*Request) bool {
	return true
}

func (m *mockRetryer) MaxRetries() int {
	return 42
}

type mockAuthorizer struct{}

func (m *mockAuthorizer) AuthorizeRequest(r *http.Request) {
	r.Header.Set("Authorization", "Yes")
}

func TestSanitizeConfig(t *testing.T) {
	mr := &mockRetryer{}
	ma := &mockAuthorizer{}

	tests := []struct {
		desc     string
		config   Config
		expected Config
	}{
		{
			desc: "minimal config gets set with default values",
			config: Config{
				BaseURL: "http://base.url",
			},
			expected: Config{
				Retryer:             DefaultRetryer{NumMaxRetries: maxDefaultRetries},
				Authorizer:          nil,
				BaseURL:             "http://base.url",
				RequestTimeout:      &defaultRequestTimeout,
				TLSHandshakeTimeout: &defaultTLSHandshakeTimeout,
				MaxIdleConns:        &defaultMaxIdleConns,
				MaxIdleConnsPerHost: &defaultMaxIdleConnsPerHost,
				IdleConnTimeout:     &defaultIdleConnTimeout,
			},
		},
		{
			desc: "fully set config respects the values",
			config: Config{
				Retryer:             mr,
				Authorizer:          ma,
				BaseURL:             "http://base.url",
				RequestTimeout:      Duration(2 * time.Second),
				TLSHandshakeTimeout: Duration(3 * time.Second),
				MaxIdleConns:        Int(42),
				MaxIdleConnsPerHost: Int(73),
				IdleConnTimeout:     Duration(4 * time.Second),
			},
			expected: Config{
				Retryer:             mr,
				Authorizer:          ma,
				BaseURL:             "http://base.url",
				RequestTimeout:      Duration(2 * time.Second),
				TLSHandshakeTimeout: Duration(3 * time.Second),
				MaxIdleConns:        Int(42),
				MaxIdleConnsPerHost: Int(73),
				IdleConnTimeout:     Duration(4 * time.Second),
			},
		},
	}
	for _, test := range tests {
		t.Run(test.desc, func(t *testing.T) {
			actual := sanitize(test.config)
			assert.Equal(t, test.expected, actual)
		})
	}
}

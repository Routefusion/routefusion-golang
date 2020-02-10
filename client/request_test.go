package client

import (
	"errors"
	"io"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"reflect"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

type basicOutputType struct {
	ID      int    `json:"id"`
	Created string `json:"created"`
}

func Test_Send(t *testing.T) {
	testCases := []struct {
		desc               string
		statusCode         int
		method             string
		path               string
		params             map[string]string
		response           []byte
		body               io.ReadSeeker
		expectedURL        string
		expectedErr        RequestFailureError
		expectedReqBody    []byte
		expectedOP         *basicOutputType
		forceHTTPClientErr bool
	}{
		{
			desc:        "basic 200 returns no error and the response message",
			statusCode:  200,
			method:      "GET",
			path:        "/test/path",
			response:    []byte(`{"id" : 123, "created" : "true"}`),
			expectedErr: nil,
			expectedURL: "/test/path",
			expectedOP:  &basicOutputType{ID: 123, Created: "true"},
		},
		{
			desc:       "503 error returns an error message after retrying",
			statusCode: 503,
			method:     "GET",
			path:       "/test/path",
			response:   []byte(`{"error_code": 503, "error_message": "service unavailable"}`),
			expectedErr: &requestError{RFError: newBaseError(ErrCodeUndefined, "http request failed after 3 attempts", nil),
				statusCode: 503, requestID: ""},
			body:            strings.NewReader(`{"somebody": "bodyval"}`),
			expectedReqBody: []byte(`{"somebody": "bodyval"}`),
			expectedURL:     "/test/path",
			expectedOP:      &basicOutputType{},
		},
		{
			desc:            "400 error returns an error message after 0 retries",
			statusCode:      400,
			method:          "GET",
			path:            "/test/path",
			response:        []byte(`{"error_code": 400, "error_message": "bad request"}`),
			expectedErr:     &requestError{RFError: newBaseError(ErrCodeUndefined, "http request failed after 0 attempts", nil), statusCode: 400, requestID: ""},
			body:            nil,
			expectedReqBody: nil,
			expectedURL:     "/test/path",
			expectedOP:      &basicOutputType{},
		},
		{
			desc:        "basic 200 returns no error and the response message with Params",
			statusCode:  200,
			method:      "GET",
			path:        "/test/path",
			response:    []byte(`{"id" : 223, "created" : "true"}`),
			params:      map[string]string{"query-1": "123", "query-2": "something-else"},
			expectedErr: nil,
			expectedURL: "/test/path?query-1=123&query-2=something-else",
			expectedOP:  &basicOutputType{ID: 223, Created: "true"},
		},
		{
			desc:               "http client returns error",
			statusCode:         0,
			method:             "GET",
			path:               "/test/path",
			response:           nil,
			expectedErr:        &requestError{RFError: newBaseError(ErrCodeUndefined, "http request failed after 0 attempts", errors.New("Get : http: nil Request.URL")), statusCode: 0, requestID: ""},
			body:               nil,
			expectedReqBody:    nil,
			expectedURL:        "/test/path",
			expectedOP:         &basicOutputType{},
			forceHTTPClientErr: true,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.desc, func(t *testing.T) {
			var op = &basicOutputType{}

			ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter,
				r *http.Request) {
				bdy, err := ioutil.ReadAll(r.Body)
				if err != nil {
					t.Fatal(err)
				}
				if !reflect.DeepEqual(testCase.expectedReqBody, bdy) {
					if string(testCase.expectedReqBody) == string(bdy) {
						// seems to be the cleanest way to check for
						// nils while allowing the tester to skip
						// byte arrays when not checking.
					} else {
						t.Errorf("expectedBody: %s but actual body: %s",
							string(testCase.expectedReqBody), string(bdy))
					}
				}
				w.WriteHeader(testCase.statusCode)
				w.Write(testCase.response)
			}))
			defer ts.Close()

			cl := NewClient(Config{BaseURL: ts.URL})
			req, err := cl.NewRequest(Operation{HTTPMethod: testCase.method,
				HTTPPath: testCase.path}, op, testCase.body, testCase.params)
			if err != nil {
				t.Fatal(err)
			}

			if testCase.forceHTTPClientErr {
				req.HTTPRequest.URL = nil
			}

			err = req.Send()
			if err != nil {
				if !isEqual(testCase.expectedErr, err) {
					t.Errorf("expected :%v, actual: %v", testCase.expectedErr,
						err)
				}
			}

			actualOP := req.Output.(*basicOutputType)
			assert.Equal(t, testCase.expectedOP, actualOP)

		})
	}
}

func isEqual(expectedErr error, actualErr error) bool {
	if expectedErr == nil && actualErr == nil {
		return true
	}

	if expectedErr == nil {
		return false
	}

	if actualErr == nil {
		return false
	}

	return expectedErr.Error() == actualErr.Error()

}

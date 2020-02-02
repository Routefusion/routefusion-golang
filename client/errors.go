package client

import (
	"fmt"
)

type ErrorCode string

const (
	ErrCodeTimeout         = "timeout"
	ErrCodeNotFound        = "not_found"
	ErrCodeUnmarshalFailed = "unmarshal_failed"

	// ErrCodeUndefined is for generic unknown/unexpected errors.
	ErrCodeUndefined = "unknown"
)

// RFError wraps the lower level errors wit a code, message and original error.
// RF stands for RouteFusion.
type RFError interface {
	// Satisfy the generic error interface.
	error

	// Returns the short phrase depicting the classification of the error.
	Code() string

	// Returns the error details message.
	Message() string

	// Returns the original error if one was set.  Nil is returned if not set.
	OrigErr() error
}

func NewRFError(code string, message string, origErr error) RFError {
	return newBaseError(code, message, origErr)
}

type RequestFailureError interface {
	RFError

	// The status code of the HTTP response.
	StatusCode() int

	// The request ID returned by the service for a request failure. This will
	// be empty if no request ID is available such as the request failed due
	// to a connection error.
	RequestID() string
}

func NewRequestFailureError(err RFError, statusCode int, reqID string) RequestFailureError {
	return newRequestError(err, statusCode, reqID)
}

func printError(code, message, extra string, origErr error) string {
	msg := fmt.Sprintf("%s: %s", code, message)
	if extra != "" {
		msg = fmt.Sprintf("%s\n\t%s", msg, extra)
	}

	if origErr != nil {
		msg = fmt.Sprintf("%s\n caused by: %s", msg, origErr.Error())
	}
	return msg
}

type baseError struct {
	code    string
	message string
	origErr error
}

func newBaseError(code string, message string, origErr error) *baseError {
	return &baseError{
		code:    code,
		message: message,
		origErr: origErr,
	}
}

func (b baseError) Error() string {
	return printError(b.code, b.message, "", nil)
}

func (b baseError) String() string {
	return b.Error()
}

func (b baseError) Code() string {
	return b.code
}

func (b baseError) Message() string {
	return b.message
}

func (b baseError) OrigErr() error {
	return b.origErr
}

type requestError struct {
	RFError
	statusCode int
	requestID  string
}

func newRequestError(err RFError, statusCode int, requestID string) *requestError {
	return &requestError{
		RFError:    err,
		statusCode: statusCode,
		requestID:  requestID,
	}
}

func (r requestError) Error() string {
	extra := fmt.Sprintf("status code: %d, request id: %s",
		r.statusCode, r.requestID)
	return printError(r.Code(), r.Message(), extra, r.OrigErr())
}

func (r requestError) String() string {
	return r.Error()
}

func (r requestError) StatusCode() int {
	return r.statusCode
}

func (r requestError) RequestID() string {
	return r.requestID
}

func (r requestError) OrigErrs() []error {
	return []error{r.OrigErr()}
}

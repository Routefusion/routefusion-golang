package client

import (
	"math/rand"
	"strconv"
	"sync"
	"time"
)

// Retryer is an interface to control retry logic for a given service.
type Retryer interface {
	RetryRules(*Request) time.Duration
	ShouldRetry(*Request) bool
	MaxRetries() int
}

// DefaultRetryer implements basic retry logic using exponential backoff for
// most services.
type DefaultRetryer struct {
	NumMaxRetries     int
	MaxRetryThreshold int
}

// MaxRetries returns the number of maximum returns the service will use to make
// an individual API
func (d DefaultRetryer) MaxRetries() int {
	return d.NumMaxRetries
}

var seededRand = rand.New(&lockedSource{src: rand.NewSource(time.Now().UnixNano())})

// RetryRules returns the delay duration before retrying this request again
func (d DefaultRetryer) RetryRules(r *Request) time.Duration {
	// Set the upper limit of delay in retrying at ~five minutes
	minTime := 30
	maxRetryThreshold := d.MaxRetryThreshold
	if maxRetryThreshold == 0 {
		maxRetryThreshold = 8
	}

	retryCount := r.RetryCount
	if retryCount > maxRetryThreshold {
		retryCount = maxRetryThreshold
	}

	delay := (1 << uint(retryCount)) * (seededRand.Intn(minTime) + minTime)
	return time.Duration(delay) * time.Millisecond
}

// ShouldRetry returns true if the request should be retried.
func (d DefaultRetryer) ShouldRetry(r *Request) bool {
	if r.HTTPResponse == nil {
		return false
	}
	if r.HTTPResponse.StatusCode >= 500 {
		return true
	}
	return false
}

// This will look in the Retry-After header, RFC 7231, for how long
// it will wait before attempting another request
func getRetryDelay(r *Request) (time.Duration, bool) {
	if !canUseRetryAfterHeader(r) {
		return 0, false
	}

	delayStr := r.HTTPResponse.Header.Get("Retry-After")
	if len(delayStr) == 0 {
		return 0, false
	}

	delay, err := strconv.Atoi(delayStr)
	if err != nil {
		return 0, false
	}

	return time.Duration(delay) * time.Second, true
}

// Will look at the status code to see if the retry header pertains to
// the status code.
func canUseRetryAfterHeader(r *Request) bool {
	switch r.HTTPResponse.StatusCode {
	case 429:
	case 503:
	default:
		return false
	}

	return true
}

// lockedSource is a thread-safe implementation of rand.Source
type lockedSource struct {
	lk  sync.Mutex
	src rand.Source
}

func (r *lockedSource) Int63() (n int64) {
	r.lk.Lock()
	n = r.src.Int63()
	r.lk.Unlock()
	return
}

func (r *lockedSource) Seed(seed int64) {
	r.lk.Lock()
	r.src.Seed(seed)
	r.lk.Unlock()
}

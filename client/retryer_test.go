package client

import (
	"net/http"
	"testing"
	"time"
)

func Test_RetryRules(t *testing.T) {
	testCases := []struct {
		desc                          string
		response                      *http.Response
		retryCount                    int
		expectedSleepTimeMinThreshold time.Duration
		expectedSleepTimeMaxThreshold time.Duration
	}{
		{
			desc: "backoff algorithm works for increased retries = 1",
			response: &http.Response{
				StatusCode: 429,
			},
			retryCount:                    5,
			expectedSleepTimeMinThreshold: 500 * time.Millisecond,
			expectedSleepTimeMaxThreshold: 4 * time.Second,
		},
		{
			desc: "backoff algorithm works for increased retries = 2",
			response: &http.Response{
				StatusCode: 429,
			},
			retryCount:                    8,
			expectedSleepTimeMinThreshold: 6 * time.Second,
			expectedSleepTimeMaxThreshold: 20 * time.Second,
		},
	}
	for _, testCase := range testCases {
		t.Run(testCase.desc, func(t *testing.T) {
			d := DefaultRetryer{NumMaxRetries: 3}
			r := &Request{HTTPResponse: testCase.response,
				RetryCount: testCase.retryCount}

			timeSecs := d.RetryRules(r)

			if timeSecs < testCase.expectedSleepTimeMinThreshold {
				t.Errorf("timesecs %v below expected min threshold: %v",
					timeSecs, testCase.expectedSleepTimeMinThreshold)
			}

			if timeSecs > testCase.expectedSleepTimeMaxThreshold {
				t.Errorf("timesecs %v above expected max threshold: %v",
					timeSecs, testCase.expectedSleepTimeMaxThreshold)
			}
		})
	}
}

func TestCanUseRetryAfter(t *testing.T) {
	testCases := []struct {
		r *Request
		e bool
	}{
		{
			&Request{
				HTTPResponse: &http.Response{StatusCode: 200},
			},
			false,
		},
		{
			&Request{
				HTTPResponse: &http.Response{StatusCode: 500},
			},
			false,
		},
		{
			&Request{
				HTTPResponse: &http.Response{StatusCode: 429},
			},
			true,
		},
		{
			&Request{
				HTTPResponse: &http.Response{StatusCode: 503},
			},
			true,
		},
	}

	for i, c := range testCases {
		a := canUseRetryAfterHeader(c.r)
		if c.e != a {
			t.Errorf("%d: expected %v, but received %v", i, c.e, a)
		}
	}
}

func TestGetRetryDelay(t *testing.T) {
	testCases := []struct {
		r     *Request
		e     time.Duration
		equal bool
		ok    bool
	}{
		{
			&Request{
				HTTPResponse: &http.Response{StatusCode: 429, Header: http.Header{"Retry-After": []string{"3600"}}},
			},
			3600 * time.Second,
			true,
			true,
		},
		{
			&Request{
				HTTPResponse: &http.Response{StatusCode: 503, Header: http.Header{"Retry-After": []string{"120"}}},
			},
			120 * time.Second,
			true,
			true,
		},
		{
			&Request{
				HTTPResponse: &http.Response{StatusCode: 503, Header: http.Header{"Retry-After": []string{"120"}}},
			},
			1 * time.Second,
			false,
			true,
		},
		{
			&Request{
				HTTPResponse: &http.Response{StatusCode: 503, Header: http.Header{"Retry-After": []string{""}}},
			},
			0 * time.Second,
			true,
			false,
		},
	}

	for i, c := range testCases {
		a, ok := getRetryDelay(c.r)
		if c.ok != ok {
			t.Errorf("%d: expected %v, but received %v", i, c.ok, ok)
		}

		if (c.e != a) == c.equal {
			t.Errorf("%d: expected %v, but received %v", i, c.e, a)
		}
	}
}

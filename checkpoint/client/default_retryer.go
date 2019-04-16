package client

import (
	"time"

	"github.com/janschumann/checkpoint-go-sdk/checkpoint/client/request"
)

// DefaultRetryer implements basic retry logic using exponential backoff for
// most services.
type DefaultRetryer struct {
	NumMaxRetries int
}

// MaxRetries returns the number of maximum returns the service will use to make
// an individual API request.
func (d DefaultRetryer) MaxRetries() int {
	return d.NumMaxRetries
}

// RetryRules returns the delay duration before retrying this request again
func (d DefaultRetryer) RetryRules(r *request.Request) time.Duration {
	// Set the upper limit of delay in retrying at ~five minutes
	minTime := 30

	retryCount := r.RetryCount

	delay := (1 << uint(retryCount)) * minTime
	return time.Duration(delay) * time.Millisecond
}

// ShouldRetry returns true if the request should be retried.
func (d DefaultRetryer) ShouldRetry(r *request.Request) bool {
	return r.Retryable && r.IsErrorRetryable()
}

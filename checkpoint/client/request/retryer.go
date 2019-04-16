package request

import (
	"time"

	"github.com/janschumann/checkpoint-go-sdk/checkpoint/checkpointerror"
)

// Retryer is an interface to control retry logic for a given service.
// The default implementation used by most services is the client.DefaultRetryer
// structure, which contains basic retry logic using exponential backoff.
type Retryer interface {
	RetryRules(*Request) time.Duration
	ShouldRetry(*Request) bool
	MaxRetries() int
}

// retryableCodes is a collection of service response codes which are retry-able
// without any further action.
var retryableCodes = map[string]struct{}{
	"validation_error":     {},
	"generic_server_error": {},
}

var retryableStatusCodes = map[int]struct{}{
	500: {},
	501: {},
}

func isCodeRetryable(code string) bool {
	if _, ok := retryableCodes[code]; ok {
		return true
	}

	return false
}

func isStatusCodeRetryable(status int) bool {
	if _, ok := retryableStatusCodes[status]; ok {
		return true
	}

	return false
}

// IsErrorRetryable returns whether the error is retryable, based on its Code.
// Returns false if error is nil.
func IsErrorRetryable(err error) bool {
	if err != nil {
		if cerr, ok := err.(checkpointerror.Error); ok {
			return isCodeRetryable(cerr.Code())
		}
	}
	return false
}

// IsErrorRetryable returns whether the error is retryable, based on its Code.
// Returns false if the request has no Error set.
//
// Alias for the utility function IsErrorRetryable
func (r *Request) IsErrorRetryable() bool {
	return IsErrorRetryable(r.Error)
}

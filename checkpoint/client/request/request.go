package request

import (
	"fmt"
	"github.com/janschumann/checkpoint-go-sdk/checkpoint/checkpointerror"
	"net/http"
	"time"
)

type ResponseValidateFunc func(interface{}) checkpointerror.Error

// A Request describes a API request and contains retry rules
type Request struct {
	HTTPMethod string
	HTTPPath   string
	Body       interface{}
	Retryer
	RetryCount   int
	Retryable    bool
	RetryDelay   time.Duration
	Error        checkpointerror.Error
	AttemptTime  time.Time
	ValidateFunc ResponseValidateFunc
}

// MessageResponses describes responses that only return a message
type MessageResponse struct {
	Message string `json:"message"`
}

// ErrorResponse describes a API error
type ErrorResponse struct {
	Message        string                      `json:"message"`
	Code           string                      `json:"code"`
	Warnings       []ValidationMessageResponse `json:"warnings,omitempty"`
	Errors         []ValidationMessageResponse `json:"errors,omitempty"`
	BlockingErrors []ValidationMessageResponse `json:"blocking-errors,omitempty"`
}

type ValidationMessageResponse struct {
	IsCurrentSession bool   `json:"current-session,omitempty"`
	Message          string `json:"message"`
}

// Cretaes a new POST request, which is used for almost all the cp api requests
func NewPostRequest(op string, input interface{}) *Request {
	return NewRequest(op, "POST", input, true)
}

// Create a new Request object
func NewRequest(op, method string, input interface{}, retryable bool) *Request {
	return &Request{
		HTTPMethod: method,
		HTTPPath:   op,
		Body:       input,
		Retryable:  retryable,
	}
}

// Weather a request should be retried
func (r *Request) WillRetry() bool {
	if r.Retryer.ShouldRetry(r) {
		r.RetryDelay = r.Retryer.RetryRules(r)
		r.RetryCount++
		r.Error = nil
		return true

	}
	return false
}

// Validate the response
func (r *Request) Validate(response interface{}) {
	if r.Error != nil {
		fmt.Println("Debug: Will not validate response due to request error.", r)
		return
	}
	if r.ValidateFunc == nil {
		fmt.Println("Warning: no response validation function set.", r)
		return
	}

	r.Error = r.ValidateFunc(response)
}

// Populates an error to the request object, if an error can be found within the response
func (r *Request) UnmarshalError(res *http.Response, err error, errorOut *ErrorResponse) {
	if err != nil || errorOut.Code != "" {
		r.unmarshalError(errorOut, err, res)
	}
}

func (r *Request) unmarshalError(err *ErrorResponse, origErr error, res *http.Response) {
	var theError checkpointerror.Error
	var errs []error
	if nil != err.Warnings {
		for _, e := range err.Warnings {
			errs = append(errs, checkpointerror.New("checkpoint_warning", e.Message, origErr))
		}
	}
	if nil != err.Errors {
		for _, e := range err.Errors {
			errs = append(errs, checkpointerror.New("checkpoint_error", e.Message, origErr))
		}
	}
	if nil != err.BlockingErrors {
		for _, e := range err.BlockingErrors {
			errs = append(errs, checkpointerror.New("checkpoint_blocking_error", e.Message, origErr))
		}
	}

	if len(errs) == 0 && origErr != nil {
		errs = append(errs, origErr)
		theError = checkpointerror.NewBatchError("checkpoint_batch_error", origErr.Error(), errs)
	} else {
		theError = checkpointerror.NewBatchError(err.Code, err.Message, errs)
	}

	if res != nil {
		theError = checkpointerror.NewRequestFailure(theError, res)
	}

	r.Error = theError
}

func EnsureMessageOK(output interface{}) checkpointerror.Error {
	if out, ok := output.(*MessageResponse); ok {
		if out.Message == "OK" {
			return nil
		}
		return checkpointerror.New("response_validation_error", "Message did not return OK", nil)
	}

	return checkpointerror.New("response_validation_error", "Request did not return request.MessageResponse", nil)
}

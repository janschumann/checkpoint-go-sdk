package checkpointerror

import (
	"fmt"
	"net/http"
)

// SprintError returns a string of the formatted error code.
func SprintError(code, message, extra string, origErr error) string {
	msg := fmt.Sprintf("%s: %s", code, message)
	if extra != "" {
		msg = fmt.Sprintf("%s\n\t%s", msg, extra)
	}
	if origErr != nil {
		msg = fmt.Sprintf("%s\ncaused by: %s", msg, origErr.Error())
	}
	return msg
}

// A baseError wraps the code and message which defines an error. It also
// can be used to wrap an original error object.
type baseError struct {
	// Classification of error
	code string

	// Detailed information about error
	message string

	// Optional original error this error is based off of. Allows building
	// chained errors.
	errs []error
}

// So that the Error interface type can be included as an anonymous field
// in the requestError struct and not conflict with the error.Error() method.
type checkpointError Error

// A requestError wraps a request or service error.
type requestError struct {
	checkpointError
	response http.Response
}

// newRequestError returns a wrapped error with additional information for
// request status code
func newRequestError(err Error, response *http.Response) *requestError {
	return &requestError{
		checkpointError: err,
		response:        *response,
	}
}

// Error returns the string representation of the error.
// Satisfies the error interface.
func (r requestError) Error() string {
	extra := fmt.Sprintf("status code: %d",
		r.StatusCode())
	return SprintError(r.Code(), r.Message(), extra, r.OrigErr())
}

// String returns the string representation of the error.
// Alias for Error to satisfy the stringer interface.
func (r requestError) String() string {
	return r.Error()
}

// StatusCode returns the wrapped status code for the error
func (r requestError) StatusCode() int {
	return r.response.StatusCode
}

// OrigErrs returns the original errors if one was set. An empty slice is
// returned if no error was set.
func (r requestError) OrigErrs() []error {
	if b, ok := r.checkpointError.(BatchedErrors); ok {
		return b.OrigErrs()
	}
	return []error{r.OrigErr()}
}

// newBaseError returns an error object for the code, message, and errors.
func newBaseError(code, message string, origErrs []error) *baseError {
	b := &baseError{
		code:    code,
		message: message,
		errs:    origErrs,
	}

	return b
}

// Error returns the string representation of the error.
//
// Satisfies the error interface.
func (b baseError) Error() string {
	size := len(b.errs)
	if size > 0 {
		return SprintError(b.code, b.message, "", errorList(b.errs))
	}

	return SprintError(b.code, b.message, "", nil)
}

// String returns the string representation of the error.
// Alias for Error to satisfy the stringer interface.
func (b baseError) String() string {
	return b.Error()
}

// Code returns the short phrase depicting the classification of the error.
func (b baseError) Code() string {
	return b.code
}

// Message returns the error details message.
func (b baseError) Message() string {
	return b.message
}

// OrigErr returns the original error if one was set. Nil is returned if no
// error was set. This only returns the first element in the list. If the full
// list is needed, use BatchedErrors.
func (b baseError) OrigErr() error {
	switch len(b.errs) {
	case 0:
		return nil
	case 1:
		return b.errs[0]
	default:
		if err, ok := b.errs[0].(Error); ok {
			return NewBatchError(err.Code(), err.Message(), b.errs[1:])
		}
		return NewBatchError("BatchedErrors",
			"multiple errors occurred", b.errs)
	}
}

// OrigErrs returns the original errors if one was set. An empty slice is
// returned if no error was set.
func (b baseError) OrigErrs() []error {
	return b.errs
}

// An error list that satisfies the golang interface
type errorList []error

// Error returns the string representation of the error.
//
// Satisfies the error interface.
func (e errorList) Error() string {
	msg := ""
	if size := len(e); size > 0 {
		for i := 0; i < size; i++ {
			msg += fmt.Sprintf("%s", e[i].Error())
			// We check the next index to see if it is within the slice.
			// If it is, then we append a newline. We do this, because unit tests
			// could be broken with the additional '\n'
			if i+1 < size {
				msg += "\n"
			}
		}
	}
	return msg
}

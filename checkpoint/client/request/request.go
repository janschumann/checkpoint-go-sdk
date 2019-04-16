package request

type Request struct {
	HTTPMethod string
	HTTPPath   string
	Body       interface{}
}

type Response struct {
	Success *interface{}
	Error   *ErrorResponse
}

type MessageResponse struct {
	Message string `json:"message"`
}

type ValidationMessageResponse struct {
	IsCurrentSession bool   `json:"current-session,omitempty"`
	Message          string `json:"message"`
}

type ErrorResponse struct {
	Message        string                      `json:"message"`
	Code           string                      `json:"code"`
	Warnings       []ValidationMessageResponse `json:"warnings,omitempty"`
	Errors         []ValidationMessageResponse `json:"errors,omitempty"`
	BlockingErrors []ValidationMessageResponse `json:"blocking-errors,omitempty"`
}

func NewPostRequest(op string, input interface{}) *Request {
	return NewRequest(op, "POST", input)
}

func NewRequest(op, method string, input interface{}) *Request {
	return &Request{
		HTTPMethod: method,
		HTTPPath:   op,
		Body:       input,
	}
}

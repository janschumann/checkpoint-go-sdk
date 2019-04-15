package client

import (
	"github.com/janschumann/checkpoint-go-sdk/checkpoint"
)

// A Client implements the base client request and response handling
// used by all service clients.
type Client struct {
	session     *Session
	sessionData *SessionData
}

type Request struct {
	HTTPMethod string
	HTTPPath   string
	Body       interface{}
}

type Response struct {
	Success *interface{}
	Error   *Error
}

type ValidationMessage struct {
	IsCurrentSession int    `json:"current-session"`
	Message          string `json:"message"`
}

type Error struct {
	Message        string               `json:"message"`
	Code           string               `json:"code"`
	Warnings       *[]ValidationMessage `json:"warnings"`
	Errors         *[]ValidationMessage `json:"warnings"`
	BlockingErrors *[]ValidationMessage `json:"blocking-errors"`
}

func New(cfg *checkpoint.Config) *Client {
	s := Must(NewSession(cfg))
	c := &Client{
		session: s,
	}

	return c
}

func (c *Client) NewPostRequest(op string, input interface{}) *Request {
	return c.NewRequest(op, "POST", input)
}

func (c *Client) NewRequest(op, method string, input interface{}) *Request {
	return &Request{
		HTTPMethod: method,
		HTTPPath: op,
		Body: input,
	}
}

func (c *Client) Send(req *Request, successOut interface{}) error {
	err := c.ensureSession()
	if err != nil {
		return err
	}

	return c.session.Send(req, successOut)
}

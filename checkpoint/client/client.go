package client

import (
	"github.com/janschumann/checkpoint-go-sdk/checkpoint"
	"time"
)

// A Client implements the base client request and response handling
// used by all service clients.
type Client struct {
	session        *Session
	sessionExpires int64
}

type Request struct {
	HTTPMethod  string
	HTTPPath    string
	RequestBody interface{}
	SuccessOut  interface{}
	ErrorOut    interface{}
}

type LoginInput struct {
	User     string `json:"user"`
	Password string `json:"password"`
}

type LastLogin struct {
	Posix int    `json:"posix"`
	Iso   string `json:"iso-8601"`
}

type LoginOutput struct {
	Uid       string     `json:"uid"`
	Sid       string     `json:"sid"`
	Url       string     `json:"url"`
	ExpiresIn int        `json:"session-timeout"`
	LastLogin *LastLogin `json:"last-login-was-at"`
	Version   string     `json:"api-server-version"`
}

type LoginError struct {
	Message string `json:"message"`
	Code    string `json:"code"`
}

// New will return a pointer to a new initialized service client.
func New(cfg *checkpoint.Config) *Client {
	s := Must(NewSession(cfg))
	c := &Client{
		session: s,
	}

	return c
}

func (c *Client) Send(r *Request, out, err interface{}) error {
	e := c.ensureSession()
	if e != nil {
		return e
	}

	e = c.session.Send(r, out, err)

	return e
}

func (c *Client) ensureSession() error {
	if time.Now().Unix() >= c.sessionExpires {
		creds, err := c.session.Config.Credentials.Get()
		if err != nil {
			return err
		}

		successOut := &LoginOutput{}
		errorOut := &LoginError{}

		r := &Request{
			HTTPMethod: "POST",
			HTTPPath:   "login",
			RequestBody: &LoginInput{
				User:     creds.User,
				Password: creds.Password,
			},
		}

		err = c.session.Send(r, successOut, errorOut)
		if err != nil {
			return err
		}

		c.sessionExpires = time.Now().Unix() + int64(successOut.ExpiresIn)
		c.session.HTTPClient.Set("X-chkp-sid", successOut.Sid)
	}

	return nil
}

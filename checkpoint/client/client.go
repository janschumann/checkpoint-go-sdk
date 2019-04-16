package client

import (
	"github.com/janschumann/checkpoint-go-sdk/checkpoint"
	"github.com/janschumann/checkpoint-go-sdk/checkpoint/client/http"
	"github.com/janschumann/checkpoint-go-sdk/checkpoint/client/request"
	"time"
)

const sessionIdHeader = "X-chkp-sid"

type Client struct {
	httpClient  *http.Client
	sessionData *LoginResponse
}

func New(cfgs ...*checkpoint.Config) (*Client, error) {
	cfg, err := newConfig(cfgs...)
	if err != nil {
		return nil, err
	}

	return &Client{httpClient: http.NewClient(cfg)}, nil
}

func Must(c *Client, err error) *Client {
	if err != nil {
		panic(err)
	}

	return c
}

func (c *Client) Send(req *request.Request, res interface{}) error {
	err := c.EnsureSession()
	if err != nil {
		return err
	}

	return c.httpClient.Send(req, res)
}

type LoginInput struct {
	User                string `json:"user"`
	Password            string `json:"password"`
	SessionName         string `json:"session-name"`
	ContinueLastSession bool   `json:"continue-last-session"`
}

type LastLoginResponse struct {
	Posix int64  `json:"posix"`
	Iso   string `json:"iso-8601"`
}

type LoginResponse struct {
	Uid       string             `json:"uid"`
	Sid       string             `json:"sid"`
	Url       string             `json:"url"`
	ExpiresIn int64              `json:"session-timeout"`
	LastLogin *LastLoginResponse `json:"last-login-was-at"`
	Version   string             `json:"api-server-version"`
}

func (s *LoginResponse) isValid() bool {
	if s != nil {
		return time.Now().Unix() < s.LastLogin.Posix+s.ExpiresIn
	}
	return false
}

func (c *Client) EnsureSession() error {
	if !c.sessionData.isValid() {
		out, err := c.login()
		if err != nil {
			return err
		}
		c.sessionData = out
		c.httpClient.SetHeader(sessionIdHeader, c.sessionData.Sid)
	}

	return nil
}

func (c *Client) login() (*LoginResponse, error) {
	creds, err := c.httpClient.Config.Credentials.Get()
	if err != nil {
		return nil, err
	}
	input := &LoginInput{
		User:                creds.User,
		Password:            creds.Password,
		SessionName:         *c.httpClient.Config.SessionName,
		ContinueLastSession: true,
	}
	out := &LoginResponse{}
	req := request.NewPostRequest("login", input)

	// use the underlying implementation,
	// as EnsureSession must not be called
	// for the login process
	return out, c.httpClient.Send(req, out)
}

type PublishInput struct{}

type PublishResponse struct {
	TaskId string `json:"task-id"`
}

func (c *Client) Publish() (*PublishResponse, error) {
	input := &PublishInput{}
	out := &PublishResponse{}
	req := request.NewPostRequest("publish", input)

	return out, c.Send(req, out)
}

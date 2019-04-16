package client

import (
	"github.com/janschumann/checkpoint-go-sdk/checkpoint"
	"github.com/janschumann/checkpoint-go-sdk/checkpoint/client/http"
	"github.com/janschumann/checkpoint-go-sdk/checkpoint/client/request"
	"sync"
	"time"
)

const sessionIdHeader = "X-chkp-sid"

// The main api client
type Client struct {
	httpClient  *http.Client
	sessionData *LoginResponse

	m                   sync.RWMutex
	continueLastSession bool
}

// Create a new client from a list of config objects
func New(cfgs ...*checkpoint.Config) (*Client, error) {
	cfg, err := newConfig(cfgs...)
	if err != nil {
		return nil, err
	}

	return &Client{
		httpClient:          http.NewClient(cfg),
		sessionData:         &LoginResponse{},
		continueLastSession: true,
	}, nil
}

func Must(c *Client, err error) *Client {
	if err != nil {
		panic(err)
	}

	return c
}

// Reset a session and ensure that the session will not be reused
func (c *Client) ResetSession() {
	c.httpClient.Reset()
	c.sessionData = &LoginResponse{}
	c.continueLastSession = false
}

// Ensures that the session is valid and sends the request
func (c *Client) Send(req *request.Request, res interface{}) error {
	if err := c.EnsureSession(); err != nil {
		return err
	}

	return c.doSend(req, res)
}

func (c *Client) doSend(req *request.Request, res interface{}) error {
	req.Retryer = DefaultRetryer{NumMaxRetries: 10}

	for {
		req.Error = nil
		req.AttemptTime = time.Now()

		c.httpClient.Send(req, res)
		req.Validate(res)

		if req.Error == nil {
			return nil
		} else {
			if req.WillRetry() {
				time.Sleep(req.RetryDelay)
				continue
			}
		}

		return req.Error
	}
}

// Describes the input parameters for the login request
type LoginInput struct {
	User                string `json:"user"`
	Password            string `json:"password"`
	SessionName         string `json:"session-name"`
	SessionTimeout      int64  `json:"session-timeout"`
	ContinueLastSession bool   `json:"continue-last-session"`
}

// Describes the response parameters returned by a login request
type LoginResponse struct {
	Uid       string             `json:"uid"`
	Sid       string             `json:"sid"`
	Url       string             `json:"url"`
	ExpiresIn int64              `json:"session-timeout"`
	LastLogin *LastLoginResponse `json:"last-login-was-at"`
	Version   string             `json:"api-server-version"`
}

type LastLoginResponse struct {
	Posix int64  `json:"posix"`
	Iso   string `json:"iso-8601"`
}

// Describes the input parameters for the logout request
type LogoutInput struct{}

// Weather the session is expired
func (s *LoginResponse) isValid() bool {
	if s.LastLogin != nil {
		return time.Now().Unix() < s.LastLogin.Posix+s.ExpiresIn
	}
	return false
}

// Re-new the session if needed
func (c *Client) EnsureSession() error {
	c.m.RLock()
	if c.sessionData.isValid() {
		defer c.m.RUnlock()
	} else {
		c.m.RUnlock()
		c.m.Lock()
		defer c.m.Unlock()

		// reset the session first
		c.ResetSession()

		out, err := c.login()
		if err != nil {
			return err
		}
		c.sessionData = out
		c.httpClient.SetHeader(sessionIdHeader, c.sessionData.Sid)
		c.continueLastSession = true
	}

	return nil
}

// perform the login. Do not check the session.
func (c *Client) login() (*LoginResponse, error) {
	creds, err := c.httpClient.Config.Credentials.Get()
	if err != nil {
		return nil, err
	}
	input := &LoginInput{
		User:                creds.User,
		Password:            creds.Password,
		SessionName:         *c.httpClient.Config.SessionName,
		SessionTimeout:      3600,
		ContinueLastSession: c.continueLastSession,
	}
	out := &LoginResponse{}
	req := request.NewPostRequest("login", input)

	return out, c.doSend(req, out)
}

// Describes the input parameters for the publish request
type PublishInput struct{}

// Describes the response parameters returned by a publish request
type PublishResponse struct {
	TaskId string `json:"task-id"`
}

// Publish changes
func (c *Client) Publish() (*PublishResponse, error) {
	input := &PublishInput{}
	out := &PublishResponse{}
	req := request.NewPostRequest("publish", input)

	return out, c.Send(req, out)
}

// Logout the user. Do not check the session.
func (c *Client) Logout() (*request.MessageResponse, error) {
	input := &LogoutInput{}
	out := &request.MessageResponse{}
	req := request.NewPostRequest("logout", input)
	req.Retryable = false

	err := c.doSend(req, out)

	if err == nil {
		c.ResetSession()
	}

	return out, err
}

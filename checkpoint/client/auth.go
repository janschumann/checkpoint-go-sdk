package client

import (
	"time"
)

const sessionIdHeader = "X-chkp-sid"

type LoginInput struct {
	User     string `json:"user"`
	Password string `json:"password"`
}

type LastLogin struct {
	Posix int64  `json:"posix"`
	Iso   string `json:"iso-8601"`
}

type SessionData struct {
	Uid       string     `json:"uid"`
	Sid       string     `json:"sid"`
	Url       string     `json:"url"`
	ExpiresIn int64      `json:"session-timeout"`
	LastLogin *LastLogin `json:"last-login-was-at"`
	Version   string     `json:"api-server-version"`
}

func (s *SessionData) isValid() bool {
	return time.Now().Unix() < s.LastLogin.Posix+s.ExpiresIn
}

func (c *Client) ensureSession() error {
	if !c.sessionData.isValid() {
		creds, err := c.session.Config.Credentials.Get()
		if err != nil {
			return err
		}
		input := &LoginInput{
			User:     creds.User,
			Password: creds.Password,
		}
		out := &SessionData{}
		req := c.NewPostRequest("login", input)

		err = c.session.Send(req, out)

		c.sessionData = out
		c.session.HTTPClient.Set(sessionIdHeader, c.sessionData.Sid)
	}

	return nil
}

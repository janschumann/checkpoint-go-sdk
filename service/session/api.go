package session

import (
	"github.com/janschumann/checkpoint-go-sdk/checkpoint/client/request"
)

type ShowSessionsInput struct {
	Offset                int  `json:"offset"`
	Limit                 int  `json:"limit"`
	ViewPublishedSessions bool `json:"view-published-sessions"`
}

type ShowSessionsResponse struct {
	From    int                `json:"from"`
	Objects *[]SessionResponse `json:"objects"`
	To      int                `json:"to"`
	Total   int                `json:"total"`
}

type LastLoginResponse struct {
	Posix int64  `json:"posix"`
	Iso   string `json:"iso-8601"`
}

type SessionResponse struct {
	Uid       string             `json:"uid"`
	Sid       string             `json:"sid"`
	Url       string             `json:"url"`
	ExpiresIn int64              `json:"httpClient-timeout"`
	LastLogin *LastLoginResponse `json:"last-login-was-at"`
	Version   string             `json:"api-server-version"`
}

type DisconnectInput struct {
	Uid     string `json:"uid"`
	Discard bool   `json:"discard"`
}

type DisconnectOutput struct {
	Message string `json:"message"`
}

type TakeOverSessionInput struct {
	Uid                     string `json:"uid"`
	DisconnectActiveSession bool   `json:"disconnect-active-httpClient"`
}

type TakeOverSessionOutput struct {
}

func (c *SessionService) ShowSessions() (*ShowSessionsResponse, error) {
	//@todo implement paging
	input := &ShowSessionsInput{
		Limit:                 500,
		Offset:                0,
		ViewPublishedSessions: true,
	}
	out := &ShowSessionsResponse{}
	req := request.NewPostRequest("show-sessions", input)

	return out, c.Send(req, out)
}

func (c *SessionService) Disconnect(input *DisconnectInput) (*DisconnectOutput, error) {
	out := &DisconnectOutput{}
	req := request.NewPostRequest("disconnect", input)

	return out, c.Send(req, out)
}

func (c *SessionService) TakeOverSession(input *TakeOverSessionInput) (*TakeOverSessionOutput, error) {
	out := &TakeOverSessionOutput{}
	req := request.NewPostRequest("take-over-httpClient", input)

	return out, c.Send(req, out)
}

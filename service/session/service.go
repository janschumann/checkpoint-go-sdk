package session

import (
	"github.com/janschumann/checkpoint-go-sdk/checkpoint/client"
)

type SessionService struct {
	*client.Client
}

func New(c *client.Client) *SessionService {
	return &SessionService{
		Client: c,
	}
}

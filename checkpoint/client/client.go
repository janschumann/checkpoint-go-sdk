package client

import (
	"github.com/janschumann/checkpoint-go-sdk/checkpoint"
)

// A Client implements the base client request and response handling
// used by all service clients.
type Client struct {
	Session *Session
}

// New will return a pointer to a new initialized service client.
func New(cfg *checkpoint.Config) *Client {
	s := Must(NewSession(cfg))
	svc := &Client{
		Session: s,
	}

	return svc
}

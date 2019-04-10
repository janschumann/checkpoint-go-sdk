package host

import (
	"github.com/janschumann/checkpoint-go-sdk/checkpoint/client"
)

type HostService struct {
	*client.Client
}

func New(c *client.Client) *HostService {
	return &HostService{
		Client: c,
	}
}

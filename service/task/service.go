package task

import "github.com/janschumann/checkpoint-go-sdk/checkpoint/client"

type TaskService struct {
	*client.Client
}

func New(c *client.Client) *TaskService {
	return &TaskService{
		Client: c,
	}
}

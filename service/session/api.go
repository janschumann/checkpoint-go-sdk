package session

import (
	"github.com/janschumann/checkpoint-go-sdk/checkpoint/client"
)

type PublishInput struct{}

type PublishOutput struct {
	TaskId string `json:"task-id"`
}

type PublishError struct {
	Message string `json:"message"`
	Code    string `json:"code"`
}

func (c *SessionService) PublishRequest(input *PublishInput) (*client.Request, *PublishOutput, *PublishError) {
	if input == nil {
		input = &PublishInput{}
	}

	op := &client.Request{
		HTTPMethod:  "POST",
		HTTPPath:    "publish",
		RequestBody: input,
	}

	successOut := &PublishOutput{}
	errorOut := &PublishError{}

	return op, successOut, errorOut
}

func (c *SessionService) Publish() (*PublishOutput, *PublishError, error) {
	op, successOut, errorOut := c.PublishRequest(nil)
	err := c.Client.Send(op, successOut, errorOut)
	return successOut, errorOut, err
}

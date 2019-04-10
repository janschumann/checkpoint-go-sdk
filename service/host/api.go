package host

import (
	"github.com/janschumann/checkpoint-go-sdk/checkpoint/client"
)

type AddHostInput struct {
	Name       string `json:"name"`
	IpAddress  string `json:"ip-address,omitempty"`
	Ip4Address string `json:"ipv4-address,omitempty"`
	Ip6Address string `json:"ipv6-address,omitempty"`
}

type AddHostOutput struct {
	Name       string `json:"name"`
	Uid        string `json:"uid"`
	Ip4Address string `json:"ipv4-address"`
	Ip6Address string `json:"ipv6-address"`
}

type ValidationMessage struct {
	IsCurrentSession int    `json:"current-session"`
	Message          string `json:"message"`
}
type AddHostError struct {
	Message        string `json:"message"`
	Code           string `json:"code"`
	Warnings       *[]ValidationMessage
	Errors         *[]ValidationMessage
	BlockingErrors *[]ValidationMessage
}

func (c *HostService) AddHost(input *AddHostInput) (*AddHostOutput, *AddHostError, error) {
	op, successOut, errorOut := c.AddHostRequest(input)
	err := c.Client.Send(op, successOut, errorOut)
	return successOut, errorOut, err
}

func (c *HostService) AddHostRequest(input *AddHostInput) (req *client.Request, output *AddHostOutput, error *AddHostError) {
	if input == nil {
		input = &AddHostInput{}
	}

	op := &client.Request{
		HTTPMethod:  "POST",
		HTTPPath:    "add-host",
		RequestBody: input,
	}

	output = &AddHostOutput{}
	error = &AddHostError{}

	return op, output, error
}

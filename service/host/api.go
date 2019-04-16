package host

import (
	"fmt"
	"github.com/janschumann/checkpoint-go-sdk/checkpoint/checkpointerror"
	"github.com/janschumann/checkpoint-go-sdk/checkpoint/client/request"
)

//
// Re-Used types
//

type HostResponse struct {
	Uid        string `json:"uid"`
	Name       string `json:"name"`
	Ip4Address string `json:"ipv4-address"`
	Ip6Address string `json:"ipv6-address"`
}

type HostsResponse struct {
	From    int             `json:"from"`
	Objects *[]HostResponse `json:"objects"`
	To      int             `json:"to"`
	Total   int             `json:"total"`
}

//
// add-host
//

type AddHostInput struct {
	Name       string `json:"name"`
	IpAddress  string `json:"ip-address,omitempty"`
	Ip4Address string `json:"ipv4-address,omitempty"`
	Ip6Address string `json:"ipv6-address,omitempty"`
}

func (c *HostService) AddHost(input *AddHostInput) (*HostResponse, error) {
	out := &HostResponse{}
	req := request.NewPostRequest("add-host", input)
	req.ValidateFunc = validateAddHost

	return out, c.Client.Send(req, out)
}

func validateAddHost(output interface{}) checkpointerror.Error {
	if out, ok := output.(*HostResponse); ok {
		if out.Uid != "" {
			return nil
		}
		return checkpointerror.New("response_validation_error", fmt.Sprintf("Host response is empty."), nil)
	}

	return checkpointerror.New("response_validation_error", "Request did not return HostResponse", nil)
}

//
// set-host
//

type SetHostInput struct {
	Uid        string `json:"uid"`
	Name       string `json:"name,omitempty"`
	IpAddress  string `json:"ip-address,omitempty"`
	Ip4Address string `json:"ipv4-address,omitempty"`
	Ip6Address string `json:"ipv6-address,omitempty"`
}

func (c *HostService) SetHost(input *SetHostInput) (*HostResponse, error) {
	out := &HostResponse{}
	req := request.NewPostRequest("set-host", input)

	return out, c.Client.Send(req, out)
}

//
// show-host
//

type ShowHostInput struct {
	Uid string `json:"uid"`
}

func (c *HostService) ShowHost(input *ShowHostInput) (*HostResponse, error) {
	out := &HostResponse{}
	req := request.NewPostRequest("show-host", input)

	return out, c.Client.Send(req, out)
}

//
// show-hosts
//

type ShowHostsInput struct {
	Limit  int `json:"limit"`
	Offset int `json:"offset"`
}

func (c *HostService) ShowHosts() (*HostsResponse, error) {
	out := &HostsResponse{}

	//@todo implement paging
	input := &ShowHostsInput{
		Limit:  500,
		Offset: 0,
	}

	req := request.NewPostRequest("show-hosts", input)

	return out, c.Client.Send(req, out)
}

//
// delete-host
//

type DeleteHostInput struct {
	Uid string `json:"uid"`
}

func (c *HostService) DeleteHost(input *ShowHostInput) (*request.MessageResponse, error) {
	out := &request.MessageResponse{}
	req := request.NewPostRequest("delete-host", input)
	req.ValidateFunc = request.EnsureMessageOK

	return out, c.Client.Send(req, out)
}

func (c *HostService) DeleteHosts() error {
	var errs []error

	hsr, err := c.ShowHosts()
	if err != nil {
		errs = append(errs, checkpointerror.New("show_hosts_failed", "Failed to retrieve hosts", err))
	} else {
		for _, o := range *hsr.Objects {
			_, err := c.DeleteHost(&ShowHostInput{
				Uid: o.Uid,
			})
			if err != nil {
				errs = append(errs, checkpointerror.New("delete_hosts_failed", fmt.Sprintf("Failed to delete host %s", o.Uid), err))
			}
		}
	}

	_, err = c.Client.Publish()
	if err != nil {
		errs = append(errs, checkpointerror.New("delete_hosts_failed", "Failed to publish", err))
	}

	if len(errs) > 0 {
		return checkpointerror.NewBatchError("delete_hosts_failed", "Failed to delete hosts", errs)
	}

	return nil
}

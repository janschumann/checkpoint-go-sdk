package host

import "github.com/janschumann/checkpoint-go-sdk/checkpoint/client/request"

//
// Re-Used types
//

type HostResponse struct {
	Name       string `json:"name"`
	Uid        string `json:"uid"`
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

	return out, c.Client.Send(req, out)
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

	return out, c.Client.Send(req, out)
}

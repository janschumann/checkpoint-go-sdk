package host

type AddHostInput struct {
	Name       string `json:"name"`
	IpAddress  string `json:"ip-address,omitempty"`
	Ip4Address string `json:"ipv4-address,omitempty"`
	Ip6Address string `json:"ipv6-address,omitempty"`
}

type HostData struct {
	Name       string `json:"name"`
	Uid        string `json:"uid"`
	Ip4Address string `json:"ipv4-address,omitempty"`
	Ip6Address string `json:"ipv6-address,omitempty"`
}

type HostsOutput struct {
	From    int         `json:"from"`
	Objects *[]HostData `json:"objects"`
	To      int         `json:"to"`
	Total   int         `json:"total"`
}

type ShowHostInput struct {
	Uid string `json:"uid"`
}

type ShowHostsInput struct {
	Limit  int `json:"limit"`
	Offset int `json:"offset"`
}

type DeleteHostOutput struct {
	Message string `json:"message"`
}

func (c *HostService) AddHost(input *AddHostInput) (*HostData, error) {
	out := &HostData{}
	req := c.NewPostRequest("add-host", input)

	return out, c.Client.Send(req, out)
}

func (c *HostService) ShowHost(input *ShowHostInput) (*HostData, error) {
	out := &HostData{}
	req := c.NewPostRequest("show-host", input)

	return out, c.Client.Send(req, out)
}

func (c *HostService) ShowHosts() (*HostsOutput, error) {
	out := &HostsOutput{}

	//@todo implement paging
	input := &ShowHostsInput{
		Limit:  500,
		Offset: 0,
	}

	req := c.NewPostRequest("show-hosts", input)

	return out, c.Client.Send(req, out)
}

func (c *HostService) DeleteHost(input *ShowHostInput) (*DeleteHostOutput, error) {
	out := &DeleteHostOutput{}
	req := c.NewPostRequest("delete-host", input)

	return out, c.Client.Send(req, out)
}

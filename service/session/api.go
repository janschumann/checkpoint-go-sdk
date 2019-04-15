package session

type PublishInput struct{}

type PublishOutput struct {
	TaskId string `json:"task-id"`
}

type ShowSessionsInput struct {
	Offset                int  `json:"offset"`
	Limit                 int  `json:"limit"`
	ViewPublishedSessions bool `json:"view-published-sessions"`
}

type ShowObjectsInput struct {
	Offset int `json:"offset"`
	Limit  int `json:"limit"`
}

type SessionOutput struct {
	Name string `json:"name"`
	Uid  string `json:"uid"`
	Type string `json:"type"`
}

type ShowSessionsOutput struct {
	From    int `json:"from"`
	Objects *[]SessionOutput
	To      int `json:"to"`
	Total   int `json:"total"`
}

type ShowObjectsOutput struct {
	From    int `json:"from"`
	Objects *[]SessionOutput
	To      int `json:"to"`
	Total   int `json:"total"`
}

type DisconnectInput struct {
	Uid     string `json:"uid"`
	Discard bool   `json:"discard"`
}

type DisconnectOutput struct {
	Message string `json:"message"`
}

func (c *SessionService) Publish() (*PublishOutput, error) {
	input := &PublishInput{}
	out := &PublishOutput{}
	req := c.NewPostRequest("publish", input)

	return out, c.Client.Send(req, out)
}

func (c *SessionService) ShowSessions() (*ShowSessionsOutput, error) {
	//@todo implement paging
	input := &ShowSessionsInput{
		Limit:                 500,
		Offset:                0,
		ViewPublishedSessions: true,
	}
	out := &ShowSessionsOutput{}
	req := c.NewPostRequest("show-sessions", input)

	return out, c.Client.Send(req, out)
}

func (c *SessionService) ShowObjects() (*ShowObjectsOutput, error) {
	input := &ShowObjectsInput{
		Limit:  500,
		Offset: 0,
	}
	out := &ShowObjectsOutput{}
	req := c.NewPostRequest("show-objects", input)

	return out, c.Client.Send(req, out)
}

func (c *SessionService) Disconnect(input *DisconnectInput) (*DisconnectOutput, error) {
	out := &DisconnectOutput{}
	req := c.NewPostRequest("disconnect", input)

	return out, c.Client.Send(req, out)
}

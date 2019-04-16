package http

import (
	"fmt"
	"github.com/dghubble/sling"
	"github.com/janschumann/checkpoint-go-sdk/checkpoint"
	"github.com/janschumann/checkpoint-go-sdk/checkpoint/checkpointerror"
	"github.com/janschumann/checkpoint-go-sdk/checkpoint/client/request"
	"net/url"
)

type Client struct {
	Config *checkpoint.Config
	Sling  *sling.Sling
}

func NewClient(config *checkpoint.Config) *Client {
	return &Client{
		Config: config,
		Sling:  newSling(config),
	}
}

func newSling(c *checkpoint.Config) *sling.Sling {
	u := &url.URL{
		Host:   fmt.Sprintf("%s:%d", checkpoint.StringValue(c.ApiHost), checkpoint.IntValue(c.ApiPort)),
		Scheme: fmt.Sprintf("%s", checkpoint.StringValue(c.ApiScheme)),
		Path:   fmt.Sprintf("%s/%s", checkpoint.StringValue(c.ApiContext), checkpoint.StringValue(c.ApiVersion)),
	}

	return sling.New().
		Client(c.HTTPClient).
		Base(u.String()).
		Set("Content-Type", "application/json")
}

// Set a header to the base sling object
func (c *Client) SetHeader(k, v string) {
	c.Sling.Set(k, v)
}

// Reset the sling object. Removes the session token.
func (c *Client) Reset() {
	c.Sling = newSling(c.Config)
}

// Send the request vi Sling.Receive()
func (c *Client) Send(req *request.Request, successOutput interface{}) {
	errorOut := &request.ErrorResponse{}

	s := c.Sling.New()
	s.BodyJSON(req.Body)

	path := fmt.Sprintf("v%s/%s", checkpoint.StringValue(c.Config.ApiVersion), req.HTTPPath)
	switch req.HTTPMethod {
	case "HEAD":
		s.Head(path)
	case "GET":
		s.Get(path)
	case "POST":
		s.Post(path)
	case "PUT":
		s.Put(path)
	case "PATCH":
		s.Path(path)
	case "DELETE":
		s.Delete(path)
	case "OPTIONS":
		s.Options(path)
	case "TRACE":
		s.Trace(path)
	case "CONNECT":
		s.Connect(path)
	default:
		req.Error = checkpointerror.New("request_error", fmt.Sprintf("Unknown operation '%s'", req.HTTPMethod), nil)
		return
	}

	res, err := s.Receive(successOutput, errorOut)
	req.UnmarshalError(res, err, errorOut)
}

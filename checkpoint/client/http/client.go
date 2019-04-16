package http

import (
	"errors"
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

func (c *Client) SetHeader(k, v string) {
	c.Sling.Set(k, v)
}

func (c *Client) Send(req *request.Request, successOutput interface{}) error {
	s := c.Sling.New()

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
		return errors.New(fmt.Sprintf("Unknown operation '%s'", req.HTTPMethod))
	}

	s.BodyJSON(req.Body)

	errorOut := &request.ErrorResponse{}

	// @todo implement a ResponseError object which allows to set the original httpresponse?
	// for now we ignote the original http response
	_, err := s.Receive(successOutput, errorOut)
	if err != nil || errorOut.Code != "" {
		return newErrorFromErrorResponse(errorOut, err)
	}

	return nil
}

func newErrorFromErrorResponse(err *request.ErrorResponse, origErr error) checkpointerror.Error {
	var errs []error
	if nil != err.Warnings {
		for _, e := range err.Warnings {
			errs = append(errs, checkpointerror.New("warning", e.Message, origErr))
		}
	}
	if nil != err.Errors {
		for _, e := range err.Errors {
			errs = append(errs, checkpointerror.New("error", e.Message, origErr))
		}
	}
	if nil != err.BlockingErrors {
		for _, e := range err.BlockingErrors {
			errs = append(errs, checkpointerror.New("blocking error", e.Message, origErr))
		}
	}

	return checkpointerror.NewBatchError(err.Code, err.Message, errs)
}

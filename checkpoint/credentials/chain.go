// borrowed from aws-sdk-go

package credentials

import (
	"github.com/janschumann/checkpoint-go-sdk/checkpoint/checkpointerror"
)

type ChainProvider struct {
	Providers []Provider
	curr      Provider
}

func NewChainCredentials(providers []Provider) *Credentials {
	return NewCredentials(&ChainProvider{
		Providers: append([]Provider{}, providers...),
	})
}

func (c *ChainProvider) Retrieve() (Value, error) {
	var errs []error
	for _, p := range c.Providers {
		creds, err := p.Retrieve()
		if err == nil {
			c.curr = p
			return creds, nil
		}
		errs = append(errs, err)
	}
	c.curr = nil

	return Value{}, checkpointerror.NewBatchError("NoCredentialProviders",
		"no valid providers in chain", errs)
}

func (c *ChainProvider) IsExpired() bool {
	if c.curr != nil {
		return c.curr.IsExpired()
	}

	return true
}

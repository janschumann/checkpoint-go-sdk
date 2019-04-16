// borrowed from aws-sdk-go

package credentials

import (
	"github.com/janschumann/checkpoint-go-sdk/checkpoint/checkpointerror"
)

// A ChainProvider will search for a provider which returns credentials
// and cache that provider until Retrieve is called again.
//
// The ChainProvider provides a way of chaining multiple providers together
// which will pick the first available using priority order of the Providers
// in the list.
//
// If none of the Providers retrieve valid credentials Value, ChainProvider's
// Retrieve() will return the error ErrNoValidProvidersFoundInChain.
//
// If a Provider is found which returns valid credentials Value ChainProvider
// will cache that Provider for all calls to IsExpired(), until Retrieve is
// called again.
type ChainProvider struct {
	Providers []Provider
	curr      Provider
}

// NewChainCredentials returns a pointer to a new Credentials object
// wrapping a chain of providers.
func NewChainCredentials(providers []Provider) *Credentials {
	return NewCredentials(&ChainProvider{
		Providers: append([]Provider{}, providers...),
	})
}

// Retrieve returns the credentials value or error if no provider returned
// without error.
//
// If a provider is found it will be cached and any calls to IsExpired()
// will return the expired state of the cached provider.
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

// IsExpired will returned the expired state of the currently cached provider
// if there is one.  If there is no current provider, true will be returned.
func (c *ChainProvider) IsExpired() bool {
	if c.curr != nil {
		return c.curr.IsExpired()
	}

	return true
}

package defaults

import (
	"github.com/janschumann/checkpoint-go-sdk/checkpoint"
	"github.com/janschumann/checkpoint-go-sdk/checkpoint/credentials"
	"net/http"
)

const defaultScheme = "https"
const defaultPort = 443
const defaultApiPath = "web_api"
const defaultApiVersion = "1.3"

// Config returns the default configuration without credentials.
// To retrieve a config with credentials also included use
// `defaults.Get().Config` instead.
//
// Generally you shouldn't need to use this method directly, but
// is available if you need to reset the configuration of an
// existing service client or session.
func Config() *checkpoint.Config {
	return checkpoint.NewConfig().
		WithCredentials(credentials.AnonymousCredentials).
		WithHTTPClient(http.DefaultClient).
		WithApiScheme(defaultScheme).
		WithApiPort(defaultPort).
		WithApiContext(defaultApiPath).
		WithApiVersion(defaultApiVersion)
}

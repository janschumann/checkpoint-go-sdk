package checkpoint

import (
	"fmt"
	"github.com/janschumann/checkpoint-go-sdk/checkpoint/credentials"
	"io"
	"net/http"
)

type Config struct {
	Credentials *credentials.Credentials

	// The name of the session to easily identify it
	SessionName *string

	// The HTTP client to use when sending requests. Defaults to
	// `http.DefaultClient`. @see defaults package
	HTTPClient *http.Client

	ApiScheme        *string
	ApiHost          *string
	ApiPort          *int
	ApiContext       *string
	ApiVersion       *string
	FingerprintFile  *string
	CheckFingerprint *bool

	// ssl settings
	CustomCABundle *io.Reader
	Insecure       *bool
}

func NewConfig() *Config {
	return &Config{}
}

func (c *Config) WithCredentials(creds *credentials.Credentials) *Config {
	c.Credentials = creds
	return c
}

func (c *Config) WithInsecure(insecure bool) *Config {
	c.Insecure = &insecure
	return c
}

func (c *Config) WithCustomCABundle(bundle *io.Reader) *Config {
	c.CustomCABundle = bundle
	return c
}

func (c *Config) WithHTTPClient(client *http.Client) *Config {
	c.HTTPClient = client
	return c
}

func (c *Config) WithApiScheme(scheme string) *Config {
	c.ApiScheme = &scheme
	return c
}

func (c *Config) WithApiHost(host string) *Config {
	c.ApiHost = &host
	return c
}

func (c *Config) WithApiPort(port int) *Config {
	c.ApiPort = &port
	return c
}

func (c *Config) WithApiContext(context string) *Config {
	c.ApiContext = &context
	return c
}

func (c *Config) WithApiVersion(version string) *Config {
	c.ApiVersion = &version
	return c
}

func (c *Config) WithSessionName(name string) *Config {
	c.SessionName = String(fmt.Sprintf("Terraform%s", name))
	return c
}

func (c *Config) MergeIn(cfgs ...*Config) {
	for _, other := range cfgs {
		mergeInConfig(c, other)
	}
}

func mergeInConfig(dst *Config, other *Config) {
	if other == nil {
		return
	}

	if other.Credentials != nil {
		dst.Credentials = other.Credentials
	}

	if other.HTTPClient != nil {
		dst.HTTPClient = other.HTTPClient
	}

	if other.ApiScheme != nil {
		dst.ApiScheme = other.ApiScheme
	}

	if other.ApiHost != nil {
		dst.ApiHost = other.ApiHost
	}

	if other.ApiPort != nil {
		dst.ApiPort = other.ApiPort
	}

	if other.ApiContext != nil {
		dst.ApiContext = other.ApiContext
	}

	if other.ApiVersion != nil {
		dst.ApiVersion = other.ApiVersion
	}

	if other.SessionName != nil {
		dst.SessionName = other.SessionName
	}

	if other.Insecure != nil {
		dst.Insecure = other.Insecure
	}
}

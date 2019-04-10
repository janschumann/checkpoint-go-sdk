package checkpoint

import (
	"github.com/janschumann/checkpoint-go-sdk/checkpoint/credentials"
	"io"
	"net/http"
)

type Config struct {
	// Enables verbose error printing of all credential chain errors.
	// Should be used when wanting to see all errors while attempting to
	// retrieve credentials.
	CredentialsChainVerboseErrors *bool

	// The credentials object to use when signing requests. Defaults to a
	// chain of credential providers to search for credentials in environment
	// variables, shared credential file, and EC2 Instance Roles.
	Credentials *credentials.Credentials

	// The HTTP client to use when sending requests. Defaults to
	// `http.DefaultClient`. @see defaults package
	HTTPClient *http.Client

	ApiScheme  *string
	ApiHost    *string
	ApiPort    *int
	ApiContext *string
	ApiVersion *string

	FingerprintFile  *string
	CheckFingerprint *bool

	Proxy *string

	// Set this to `true` to disable SSL when sending requests. Defaults
	// to `false`.
	DisableSSL *bool

	CustomCABundle *io.Reader

	// An integer value representing the logging level. The default log level
	// is zero (LogOff), which represents no logging. To enable logging set
	// to a LogLevel Value.
	LogLevel *LogLevelType

	// The logger writer interface to write logging messages to. Defaults to
	// standard out.
	Logger Logger
}

// NewConfig returns a new Config pointer that can be chained with builder
// methods to set multiple configuration values inline without using pointers.
//
//     // Create Session with MaxRetry configuration to be shared by multiple
//     // service clients.
//     sess := session.Must(session.NewSession(aws.NewConfig().
//         WithMaxRetries(3),
//     ))
//
//     // Create S3 service client with a specific Region.
//     svc := s3.New(sess, aws.NewConfig().
//         WithRegion("us-west-2"),
//     )
func NewConfig() *Config {
	return &Config{}
}

// WithCredentialsChainVerboseErrors sets a config verbose errors boolean and returning
// a Config pointer.
func (c *Config) WithCredentialsChainVerboseErrors(verboseErrs bool) *Config {
	c.CredentialsChainVerboseErrors = &verboseErrs
	return c
}

// WithCredentials sets a config Credentials value returning a Config pointer
// for chaining.
func (c *Config) WithCredentials(creds *credentials.Credentials) *Config {
	c.Credentials = creds
	return c
}

// WithDisableSSL sets a config DisableSSL value returning a Config pointer
// for chaining.
func (c *Config) WithDisableSSL(disable bool) *Config {
	c.DisableSSL = &disable
	return c
}

func (c *Config) WithCustomCABundle(bundle *io.Reader) *Config {
	c.CustomCABundle = bundle
	return c
}

// WithHTTPClient sets a config HTTPClient value returning a Config pointer
// for chaining.
func (c *Config) WithHTTPClient(client *http.Client) *Config {
	c.HTTPClient = client
	return c
}

// WithLogLevel sets a config LogLevel value returning a Config pointer for
// chaining.
func (c *Config) WithLogLevel(level LogLevelType) *Config {
	c.LogLevel = &level
	return c
}

// WithLogger sets a config Logger value returning a Config pointer for
// chaining.
func (c *Config) WithLogger(logger Logger) *Config {
	c.Logger = logger
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

// MergeIn merges the passed in configs into the existing config object.
func (c *Config) MergeIn(cfgs ...*Config) {
	for _, other := range cfgs {
		mergeInConfig(c, other)
	}
}

func mergeInConfig(dst *Config, other *Config) {
	if other == nil {
		return
	}

	if other.CredentialsChainVerboseErrors != nil {
		dst.CredentialsChainVerboseErrors = other.CredentialsChainVerboseErrors
	}

	if other.Credentials != nil {
		dst.Credentials = other.Credentials
	}

	if other.DisableSSL != nil {
		dst.DisableSSL = other.DisableSSL
	}

	if other.HTTPClient != nil {
		dst.HTTPClient = other.HTTPClient
	}

	if other.LogLevel != nil {
		dst.LogLevel = other.LogLevel
	}

	if other.Logger != nil {
		dst.Logger = other.Logger
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
}

// Copy will return a shallow copy of the Config object. If any additional
// configurations are provided they will be merged into the new config returned.
func (c *Config) Copy(cfgs ...*Config) *Config {
	dst := &Config{}
	dst.MergeIn(c)

	for _, cfg := range cfgs {
		dst.MergeIn(cfg)
	}

	return dst
}

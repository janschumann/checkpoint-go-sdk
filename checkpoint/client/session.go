package client

import (
	"crypto/tls"
	"crypto/x509"
	"errors"
	"fmt"
	"github.com/dghubble/sling"
	"github.com/janschumann/checkpoint-go-sdk/checkpoint"
	"github.com/janschumann/checkpoint-go-sdk/checkpoint/checkpointerror"
	"github.com/janschumann/checkpoint-go-sdk/checkpoint/credentials"
	"github.com/janschumann/checkpoint-go-sdk/checkpoint/defaults"
	"io"
	"io/ioutil"
	"net"
	"net/http"
	"net/url"
	"os"
	"time"
)

type Session struct {
	Config     *checkpoint.Config
	HTTPClient *sling.Sling
}

// NewSession returns a new Session created from SDK defaults, config files,
// environment, and user provided config files. Once the Session is created
// it can be mutated to modify the Config or Handlers. The Session is safe to
// be read concurrently, but it should not be written to concurrently.
//
// If the AWS_SDK_LOAD_CONFIG environment variable is set to a truthy value
// the shared config file (~/.aws/config) will also be loaded in addition to
// the shared credentials file (~/.aws/credentials). Values set in both the
// shared config, and shared credentials will be taken from the shared
// credentials file. Enabling the Shared Config will also allow the Session
// to be built with retrieving credentials with AssumeRole set in the config.
//
// See the NewSessionWithOptions func for information on how to override or
// control through code how the Session will be created. Such as specifying the
// config profile, and controlling if shared config is enabled or not.
func NewSession(cfgs ...*checkpoint.Config) (*Session, error) {
	cfg := defaults.Config()

	// Get a merged version of the user provided config to determine if
	// credentials were.
	userCfg := &checkpoint.Config{}
	userCfg.MergeIn(cfgs...)

	envCfg := loadEnvConfig()
	if err := mergeConfigSrcs(cfg, userCfg, envCfg); err != nil {
		return nil, err
	}

	s := &Session{
		Config:     cfg,
		HTTPClient: newHttpClient(cfg),
	}

	if checkpoint.BoolValue(cfg.DisableSSL) {
		if err := disableSSL(cfg); err != nil {
			return nil, err
		}
	} else if len(envCfg.CustomCABundle) != 0 && cfg.CustomCABundle == nil {
		f, err := os.Open(envCfg.CustomCABundle)
		if err != nil {
			return nil, checkpointerror.New("LoadCustomCABundleError",
				"failed to open custom CA bundle PEM file", err)
		}
		defer f.Close()
		if err := loadCustomCABundle(cfg, f); err != nil {
			return nil, err
		}
	}

	return s, nil
}

func newHttpClient(c *checkpoint.Config) *sling.Sling {
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

// Must is a helper function to ensure the Session is valid and there was no
// error when calling a NewSession function.
//
// This helper is intended to be used in variable initialization to load the
// Session and configuration at startup. Such as:
//
//     var sess = session.Must(session.NewSession())
func Must(sess *Session, err error) *Session {
	if err != nil {
		panic(err)
	}

	return sess
}

func getTransport(client *http.Client) (*http.Transport, error) {
	var t *http.Transport
	switch v := client.Transport.(type) {
	case *http.Transport:
		t = v
	default:
		if client.Transport != nil {
			return nil, checkpointerror.New("CustomizeTransportError",
				"unable to customize HTTPClient's transport: unsupported type", nil)
		}
	}
	if t == nil {
		// Nil transport implies `http.DefaultTransport` should be used. Since
		// the SDK cannot modify, nor copy the `DefaultTransport` specifying
		// the values the next closest behavior.
		t = getCABundleTransport()
	}

	if t.TLSClientConfig == nil {
		t.TLSClientConfig = &tls.Config{}
	}

	return t, nil
}

func disableSSL(c *checkpoint.Config) error {
	t, err := getTransport(c.HTTPClient)
	if err != nil {
		return err
	}

	t.TLSClientConfig.InsecureSkipVerify = true

	c.HTTPClient.Transport = t

	return nil
}

func loadCustomCABundle(c *checkpoint.Config, bundle io.Reader) error {
	t, err := getTransport(c.HTTPClient)
	if err != nil {
		return err
	}

	p, err := loadCertPool(bundle)
	if err != nil {
		return err
	}
	t.TLSClientConfig.RootCAs = p

	c.HTTPClient.Transport = t

	return nil
}

func getCABundleTransport() *http.Transport {
	return &http.Transport{
		Proxy: http.ProxyFromEnvironment,
		DialContext: (&net.Dialer{
			Timeout:   30 * time.Second,
			KeepAlive: 30 * time.Second,
			DualStack: true,
		}).DialContext,
		MaxIdleConns:          100,
		IdleConnTimeout:       90 * time.Second,
		TLSHandshakeTimeout:   10 * time.Second,
		ExpectContinueTimeout: 1 * time.Second,
	}
}

func loadCertPool(r io.Reader) (*x509.CertPool, error) {
	b, err := ioutil.ReadAll(r)
	if err != nil {
		return nil, checkpointerror.New("LoadCustomCABundleError",
			"failed to read custom CA bundle PEM file", err)
	}

	p := x509.NewCertPool()
	if !p.AppendCertsFromPEM(b) {
		return nil, checkpointerror.New("LoadCustomCABundleError",
			"failed to load custom CA bundle PEM file", err)
	}

	return p, nil
}

func mergeConfigSrcs(cfg, userCfg *checkpoint.Config, envCfg envConfig) error {
	// Merge in user provided configuration
	cfg.MergeIn(userCfg)

	// Configure credentials if not already set
	if cfg.Credentials == credentials.AnonymousCredentials && userCfg.Credentials == nil {

		if len(envCfg.Credentials.User) > 0 {
			cfg.Credentials = credentials.NewStaticCredentialsFromCreds(
				envCfg.Credentials,
			)
		} else {
			// Fallback to default credentials provider, include mock errors
			// for the credential chain so user can identify why credentials
			// failed to be retrieved.
			cfg.Credentials = credentials.NewCredentials(&credentials.ChainProvider{
				VerboseErrors: checkpoint.BoolValue(cfg.CredentialsChainVerboseErrors),
				Providers: []credentials.Provider{
					&credProviderError{Err: checkpointerror.New("EnvAccessKeyNotFound", "failed to find credentials in the environment.", nil)},
				},
			})
		}
	}

	return nil
}

type credProviderError struct {
	Err error
}

var emptyCreds = credentials.Value{}

func (c credProviderError) Retrieve() (credentials.Value, error) {
	return credentials.Value{}, c.Err
}
func (c credProviderError) IsExpired() bool {
	return true
}

func (s *Session) Send(req *Request, successOutput interface{}) error {
	httpClient := s.HTTPClient.New()
	httpClient.BodyJSON(req.Body)
	path := fmt.Sprintf("%s/%s", checkpoint.StringValue(s.Config.ApiVersion), req.HTTPPath)

	switch req.HTTPMethod {
	case "HEAD":
		httpClient.Head(path)
	case "GET":
		httpClient.Get(path)
	case "POST":
		httpClient.Post(path)
	case "PUT":
		httpClient.Put(path)
	case "PATCH":
		httpClient.Path(path)
	case "DELETE":
		httpClient.Delete(path)
	case "OPTIONS":
		httpClient.Options(path)
	case "TRACE":
		httpClient.Trace(path)
	case "CONNECT":
		httpClient.Connect(path)
	default:
		return errors.New(fmt.Sprintf("Unknown operation '%s'", req.HTTPMethod))
	}

	errorOutput := &Error{}

	// @todo implement a ResponseError object which allows to set the original httpresponse
	_, err := httpClient.Receive(successOutput, errorOutput)

	if err != nil || errorOutput.Code != "" {
		return checkpointerror.New(errorOutput.Code, errorOutput.Message, err)
	}

	return nil
}

package client

import (
	"crypto/tls"
	"crypto/x509"
	"github.com/janschumann/checkpoint-go-sdk/checkpoint"
	"github.com/janschumann/checkpoint-go-sdk/checkpoint/checkpointerror"
	"github.com/janschumann/checkpoint-go-sdk/checkpoint/credentials"
	"github.com/janschumann/checkpoint-go-sdk/checkpoint/defaults"
	"io"
	"io/ioutil"
	"net"
	"net/http"
	"os"
	"time"
)

func newConfig(cfgs ...*checkpoint.Config) (*checkpoint.Config, error) {
	cfg := defaults.Config()

	// Get a merged version of the user provided config to determine if
	// credentials were.
	userCfg := &checkpoint.Config{}
	userCfg.MergeIn(cfgs...)

	envCfg := loadEnvConfig()
	if err := mergeConfigSrcs(cfg, userCfg, envCfg); err != nil {
		return nil, err
	}

	if checkpoint.BoolValue(cfg.Insecure) {
		if err := disableCertVerification(cfg); err != nil {
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

	return cfg, nil
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

func disableCertVerification(c *checkpoint.Config) error {
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
			// Fallback to default credentials provider
			cfg.Credentials = credentials.NewCredentials(&credentials.ChainProvider{
				Providers: []credentials.Provider{
					&credProviderError{Err: checkpointerror.New("EnvCredentialsNotFound",
						"failed to find credentials in the environment.",
						nil)},
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

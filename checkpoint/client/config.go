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
	"strconv"
	"time"
)

// merges a list of config objects
func newConfig(cfgs ...*checkpoint.Config) (*checkpoint.Config, error) {
	cfg := defaults.Config()

	userCfg := &checkpoint.Config{}
	userCfg.MergeIn(cfgs...)

	envCfg := loadEnvConfig()
	mergeConfigSrcs(cfg, userCfg, envCfg)

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

func mergeConfigSrcs(cfg, userCfg *checkpoint.Config, envCfg envConfig) {
	// Merge in user provided configuration
	cfg.MergeIn(userCfg)

	// Merge in env config values if not provided by user
	if userCfg.ApiScheme == nil && len(envCfg.ApiScheme) > 0 {
		cfg.WithApiScheme(envCfg.ApiScheme)
	}
	if userCfg.ApiHost == nil && len(envCfg.ApiHost) > 0 {
		cfg.WithApiHost(envCfg.ApiHost)
	}
	if userCfg.ApiPort == nil && len(envCfg.ApiPort) > 0 {
		if v, err := strconv.Atoi(envCfg.ApiPort); err == nil {
			cfg.WithApiPort(v)
		}
	}
	if userCfg.ApiContext == nil && len(envCfg.ApiContext) > 0 {
		cfg.WithApiContext(envCfg.ApiContext)
	}
	if userCfg.ApiVersion == nil && len(envCfg.ApiVersion) > 0 {
		cfg.WithApiVersion(envCfg.ApiVersion)
	}
	if userCfg.Insecure == nil && len(envCfg.Insecure) > 0 {
		if v, err := strconv.ParseBool(envCfg.Insecure); err == nil {
			cfg.WithInsecure(v)
		}
	}
	if userCfg.SessionName == nil && len(envCfg.SessionName) > 0 {
		cfg.WithSessionName(envCfg.SessionName)
	}

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
					&credentials.EnvProvider{},
					&credProviderError{Err: checkpointerror.New("NoCredentialsFound",
						"failed to find credentials in the chain.",
						nil)},
				},
			})
		}
	}
}

// A error provider to use as last resort in a provider chain
type credProviderError struct {
	Err error
}

func (c credProviderError) Retrieve() (credentials.Value, error) {
	return credentials.Value{}, c.Err
}
func (c credProviderError) IsExpired() bool {
	return true
}

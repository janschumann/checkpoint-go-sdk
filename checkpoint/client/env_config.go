package client

import (
	"os"
	"github.com/janschumann/checkpoint-go-sdk/checkpoint/credentials"
	"github.com/janschumann/checkpoint-go-sdk/checkpoint/defaults"
)

// EnvProviderName provides a name of the provider when config is loaded from environment.
const EnvProviderName = "EnvConfigCredentials"

// envConfig is a collection of environment values the SDK will read
// setup config from. All environment values are optional. But some values
// such as credentials require multiple values to be complete or the values
// will be ignored.
type envConfig struct {
	// The credentials object to use when signing requests. Defaults to a
	// chain of credential providers to search for credentials in environment
	// variables, shared credential file, and EC2 Instance Roles.
	Credentials credentials.Value

	ApiScheme string
	ApiHost string
	ApiPort int
	ApiContext string
	ApiVersion string

	// Set this to `true` to disable SSL when sending requests. Defaults
	// to `false`.
	DisableSSL bool

	CustomCABundle string
}

var (
	credUserEnvKey = []string{
		"CHECKPOINT_USER",
	}
	credPasswordEnvKey = []string{
		"CHECKPOINT_PASSWORD",
	}
)
)

// loadEnvConfig retrieves the SDK's environment configuration.
// See `envConfig` for the values that will be retrieved.
//
// If the environment variable `CHECKPOINT_SDK_LOAD_CONFIG` is set to a truthy value
// the shared SDK config will be loaded in addition to the SDK's specific
// configuration values.
func loadEnvConfig() envConfig {
	cfg := envConfig{}

	setFromEnvVal(&cfg.Credentials.User, credUserEnvKey)
	setFromEnvVal(&cfg.Credentials.Password, credPasswordEnvKey)

	// Require logical grouping of credentials
	if len(cfg.Credentials.User) == 0 || len(cfg.Credentials.Password) == 0 {
		cfg.Credentials = credentials.Value{}
	} else {
		cfg.Credentials.ProviderName = EnvProviderName
	}

	cfg.CustomCABundle = os.Getenv("CHECKPOINT_CA_BUNDLE")

	return cfg
}

func setFromEnvVal(dst *string, keys []string) {
	for _, k := range keys {
		if v := os.Getenv(k); len(v) > 0 {
			*dst = v
			break
		}
	}
}

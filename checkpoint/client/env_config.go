package client

import (
	"github.com/janschumann/checkpoint-go-sdk/checkpoint/credentials"
	"os"
)

// EnvProviderName provides a name of the provider when config is loaded from environment.
const EnvProviderName = "EnvConfigCredentials"

type envConfig struct {
	Credentials credentials.Value

	ApiScheme  string
	ApiHost    string
	ApiPort    int
	ApiContext string
	ApiVersion string

	Insecure bool

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

func loadEnvConfig() envConfig {
	cfg := envConfig{}

	setFromEnvVal(&cfg.Credentials.User, credUserEnvKey)
	setFromEnvVal(&cfg.Credentials.Password, credPasswordEnvKey)

	if len(cfg.Credentials.User) == 0 || len(cfg.Credentials.Password) == 0 {
		cfg.Credentials = credentials.Value{}
	} else {
		cfg.Credentials.ProviderName = EnvProviderName
	}

	cfg.CustomCABundle = os.Getenv("CHECKPOINT_CA_BUNDLE")
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

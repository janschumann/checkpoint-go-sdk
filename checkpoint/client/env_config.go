package client

import (
	"github.com/janschumann/checkpoint-go-sdk/checkpoint/credentials"
	"os"
	"strings"
)

const EnvProviderName = "EnvConfigCredentials"

type envConfig struct {
	Credentials credentials.Value

	ApiScheme  string
	ApiHost    string
	ApiPort    string
	ApiContext string
	ApiVersion string

	SessionName string
	Insecure    string

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

	cfg.CustomCABundle = strings.Trim(os.Getenv("CHECKPOINT_CA_BUNDLE"), "\"")
	cfg.ApiScheme = strings.Trim(os.Getenv("CHECKPOINT_API_SCHEME"), "\"")
	cfg.ApiHost = strings.Trim(os.Getenv("CHECKPOINT_API_HOST"), "\"")
	cfg.ApiPort = strings.Trim(os.Getenv("CHECKPOINT_API_PORT"), "\"")
	cfg.ApiContext = strings.Trim(os.Getenv("CHECKPOINT_API_CONTEXT"), "\"")
	cfg.ApiVersion = strings.Trim(os.Getenv("CHECKPOINT_API_VERSION"), "\"")
	cfg.Insecure = strings.Trim(os.Getenv("CHECKPOINT_API_INSECURE"), "\"")
	cfg.SessionName = strings.Trim(os.Getenv("CHECKPOINT_SESSION_NAME"), "\"")

	return cfg
}

func setFromEnvVal(dst *string, keys []string) {
	for _, k := range keys {
		if v := os.Getenv(k); len(v) > 0 {
			*dst = strings.Trim(v, "\"")
			break
		}
	}
}

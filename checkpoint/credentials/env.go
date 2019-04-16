// borrowed from aws-sdk-go

package credentials

import (
	"os"

	"github.com/janschumann/checkpoint-go-sdk/checkpoint/checkpointerror"
)

const EnvProviderName = "EnvProvider"

var (
	ErrEnvUserNotFound     = checkpointerror.New("EnvUserNotFound", "CHECKPOINT_USER not found in environment", nil)
	ErrEnvPasswordNotFound = checkpointerror.New("EnvPasswordNotFound", "CHECKPOINT_PASSWORD not found in environment", nil)
)

// A EnvProvider retrieves credentials from the environment variables of the
// running process. Environment credentials never expire.
//
// Environment variables used:
//
// * User:     CHECKPOINT_USER
//
// * Password: CHECKPOINT_PASSWORD
type EnvProvider struct {
	retrieved bool
}

func NewEnvCredentials() *Credentials {
	return NewCredentials(&EnvProvider{})
}

// Retrieve retrieves the keys from the environment.
func (e *EnvProvider) Retrieve() (Value, error) {
	e.retrieved = false

	user := os.Getenv("CHECKPOINT_USER")
	if user == "" {
		return Value{ProviderName: EnvProviderName}, ErrEnvUserNotFound
	}
	password := os.Getenv("CHECKPOINT_PASSWORD")
	if password == "" {
		return Value{ProviderName: EnvProviderName}, ErrEnvPasswordNotFound
	}

	e.retrieved = true
	return Value{
		User:         user,
		Password:     password,
		ProviderName: EnvProviderName,
	}, nil
}

// IsExpired returns false if the credentials have been retrieved.
func (e *EnvProvider) IsExpired() bool {
	return !e.retrieved
}

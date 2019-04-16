// borrowed from aws-sdk-go

package credentials

import "github.com/janschumann/checkpoint-go-sdk/checkpoint/checkpointerror"

const StaticProviderName = "StaticProvider"

var (
	ErrStaticCredentialsEmpty = checkpointerror.New("EmptyStaticCreds", "static credentials are empty", nil)
)

type StaticProvider struct {
	Value
}

func NewStaticCredentials(user, password string) *Credentials {
	return NewCredentials(&StaticProvider{Value: Value{
		User:     user,
		Password: password,
	}})
}

func NewStaticCredentialsFromCreds(creds Value) *Credentials {
	return NewCredentials(&StaticProvider{Value: creds})
}

func (s *StaticProvider) Retrieve() (Value, error) {
	if s.User == "" || s.Password == "" {
		return Value{ProviderName: StaticProviderName}, ErrStaticCredentialsEmpty
	}

	if len(s.Value.ProviderName) == 0 {
		s.Value.ProviderName = StaticProviderName
	}
	return s.Value, nil
}

// For StaticProvider, the credentials never expire.
func (s *StaticProvider) IsExpired() bool {
	return false
}

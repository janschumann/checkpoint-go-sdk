package credentials

import (
	"reflect"
	"testing"

	"github.com/janschumann/checkpoint-go-sdk/checkpoint/checkpointerror"
)

type secondStubProvider struct {
	creds   Value
	expired bool
	err     error
}

func (s *secondStubProvider) Retrieve() (Value, error) {
	s.expired = false
	s.creds.ProviderName = "secondStubProvider"
	return s.creds, s.err
}
func (s *secondStubProvider) IsExpired() bool {
	return s.expired
}

func TestChainProviderWithNames(t *testing.T) {
	p := &ChainProvider{
		Providers: []Provider{
			&stubProvider{err: checkpointerror.New("FirstError", "first provider error", nil)},
			&stubProvider{err: checkpointerror.New("SecondError", "second provider error", nil)},
			&secondStubProvider{
				creds: Value{
					User:     "USER",
					Password: "NOPASSWORD",
				},
			},
			&stubProvider{
				creds: Value{
					User:     "USER",
					Password: "PASSWORD",
				},
			},
		},
	}

	creds, err := p.Retrieve()
	if err != nil {
		t.Errorf("Expect no error, got %v", err)
	}
	if e, a := "secondStubProvider", creds.ProviderName; e != a {
		t.Errorf("Expect provider name to match, %v got, %v", e, a)
	}

	// Also check credentials
	if e, a := "USER", creds.User; e != a {
		t.Errorf("Expect access key ID to match, %v got %v", e, a)
	}
	if e, a := "NOPASSWORD", creds.Password; e != a {
		t.Errorf("Expect secret access key to match, %v got %v", e, a)
	}

}

func TestChainProviderGet(t *testing.T) {
	p := &ChainProvider{
		Providers: []Provider{
			&stubProvider{err: checkpointerror.New("FirstError", "first provider error", nil)},
			&stubProvider{err: checkpointerror.New("SecondError", "second provider error", nil)},
			&stubProvider{
				creds: Value{
					User:     "USER",
					Password: "PASSWORD",
				},
			},
		},
	}

	creds, err := p.Retrieve()
	if err != nil {
		t.Errorf("Expect no error, got %v", err)
	}
	if e, a := "USER", creds.User; e != a {
		t.Errorf("Expect user to match, %v got %v", e, a)
	}
	if e, a := "PASSWORD", creds.Password; e != a {
		t.Errorf("Expect password to match, %v got %v", e, a)
	}
}

func TestChainProviderIsExpired(t *testing.T) {
	stubProvider := &stubProvider{expired: true}
	p := &ChainProvider{
		Providers: []Provider{
			stubProvider,
		},
	}

	if !p.IsExpired() {
		t.Errorf("Expect expired to be true before any Retrieve")
	}
	_, err := p.Retrieve()
	if err != nil {
		t.Errorf("Expect no error, got %v", err)
	}
	if p.IsExpired() {
		t.Errorf("Expect not expired after retrieve")
	}

	stubProvider.expired = true
	if !p.IsExpired() {
		t.Errorf("Expect return of expired provider")
	}

	_, err = p.Retrieve()
	if p.IsExpired() {
		t.Errorf("Expect not expired after retrieve")
	}
}

func TestChainProviderWithNoProvider(t *testing.T) {
	p := &ChainProvider{
		Providers: []Provider{},
	}

	if !p.IsExpired() {
		t.Errorf("Expect expired with no providers")
	}
	_, err := p.Retrieve()
	if err == nil {
		t.Errorf("Expect no providers error returned")
	}
}

func TestChainProviderWithNoValidProvider(t *testing.T) {
	errs := []error{
		checkpointerror.New("FirstError", "first provider error", nil),
		checkpointerror.New("SecondError", "second provider error", nil),
	}
	p := &ChainProvider{
		Providers: []Provider{
			&stubProvider{err: errs[0]},
			&stubProvider{err: errs[1]},
		},
	}

	if !p.IsExpired() {
		t.Errorf("Expect expired with no providers")
	}
	_, err := p.Retrieve()

	if err == nil {
		t.Errorf("Expect no providers error returned")
	}
}

func TestChainProviderWithNoValidProviderWithVerboseEnabled(t *testing.T) {
	errs := []error{
		checkpointerror.New("FirstError", "first provider error", nil),
		checkpointerror.New("SecondError", "second provider error", nil),
	}
	p := &ChainProvider{
		Providers: []Provider{
			&stubProvider{err: errs[0]},
			&stubProvider{err: errs[1]},
		},
	}

	if !p.IsExpired() {
		t.Errorf("Expect expired with no providers")
	}
	_, err := p.Retrieve()

	expectErr := checkpointerror.NewBatchError("NoCredentialProviders", "no valid providers in chain", errs)
	if e, a := expectErr, err; !reflect.DeepEqual(e, a) {
		t.Errorf("Expect no providers error returned, %v, got %v", e, a)
	}
}

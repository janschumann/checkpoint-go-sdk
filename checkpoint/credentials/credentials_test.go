package credentials

import (
	"github.com/janschumann/checkpoint-go-sdk/checkpoint/checkpointerror"
	"testing"
	"time"
)

type stubProvider struct {
	creds   Value
	expired bool
	err     error
}

func (s *stubProvider) Retrieve() (Value, error) {
	s.expired = false
	s.creds.ProviderName = "stubProvider"
	return s.creds, s.err
}
func (s *stubProvider) IsExpired() bool {
	return s.expired
}

func TestCredentialsGet(t *testing.T) {
	c := NewCredentials(&stubProvider{
		creds: Value{
			User:     "USER",
			Password: "PASSWORD",
		},
		expired: true,
	})

	creds, err := c.Get()
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if e, a := "USER", creds.User; e != a {
		t.Errorf("Expect user to match, %v got %v", e, a)
	}
	if e, a := "PASSWORD", creds.Password; e != a {
		t.Errorf("Expect password to match, %v got %v", e, a)
	}
}

func TestCredentialsGetWithError(t *testing.T) {
	c := NewCredentials(&stubProvider{err: checkpointerror.New("provider error", "", nil), expired: true})

	_, err := c.Get()
	if e, a := "provider error", err.(checkpointerror.Error).Code(); e != a {
		t.Errorf("Expected provider error, %v got %v", e, a)
	}
}

func TestCredentialsExpire(t *testing.T) {
	stub := &stubProvider{}
	c := NewCredentials(stub)

	stub.expired = false
	if !c.IsExpired() {
		t.Errorf("Expected to start out expired")
	}

	if !c.IsExpired() {
		t.Errorf("Expected to be expired")
	}

	c.forceRefresh = false
	if c.IsExpired() {
		t.Errorf("Expected not to be expired")
	}

	stub.expired = true
	if !c.IsExpired() {
		t.Errorf("Expected to be expired")
	}
}

type MockProvider struct {
}

func (*MockProvider) Retrieve() (Value, error) {
	return Value{}, nil
}

func (*MockProvider) IsExpired() bool {
	return false
}

func TestCredentialsGetWithProviderName(t *testing.T) {
	stub := &stubProvider{}

	c := NewCredentials(stub)

	creds, err := c.Get()
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if e, a := creds.ProviderName, "stubProvider"; e != a {
		t.Errorf("Expected provider name to match, %v got %v", e, a)
	}
}

func TestCredentialsIsExpired_Race(t *testing.T) {
	creds := NewChainCredentials([]Provider{&MockProvider{}})

	starter := make(chan struct{})
	for i := 0; i < 10; i++ {
		go func() {
			<-starter
			for {
				creds.IsExpired()
			}
		}()
	}
	close(starter)

	time.Sleep(10 * time.Second)
}

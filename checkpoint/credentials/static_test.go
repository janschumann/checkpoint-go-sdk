package credentials

import (
	"testing"
)

func TestStaticProviderGet(t *testing.T) {
	s := StaticProvider{
		Value: Value{
			User:     "USER",
			Password: "PASSWORD",
		},
	}

	creds, err := s.Retrieve()
	if err != nil {
		t.Errorf("expect nil, got %v", err)
	}
	if e, a := "USER", creds.User; e != a {
		t.Errorf("expect %v, got %v", e, a)
	}
	if e, a := "PASSWORD", creds.Password; e != a {
		t.Errorf("expect %v, got %v", e, a)
	}
}

func TestStaticProviderIsExpired(t *testing.T) {
	s := StaticProvider{
		Value: Value{
			User:     "USER",
			Password: "PASSWORD",
		},
	}

	if s.IsExpired() {
		t.Errorf("Expect static credentials to never expire")
	}
}

package credentials

import (
	"os"
	"testing"
)

func TestEnvProviderRetrieve(t *testing.T) {
	os.Clearenv()
	os.Setenv("CHECKPOINT_USER", "user")
	os.Setenv("CHECKPOINT_PASSWORD", "password")

	e := EnvProvider{}
	creds, err := e.Retrieve()
	if err != nil {
		t.Errorf("expect nil, got %v", err)
	}

	if e, a := "user", creds.User; e != a {
		t.Errorf("expect %v, got %v", e, a)
	}
	if e, a := "password", creds.Password; e != a {
		t.Errorf("expect %v, got %v", e, a)
	}
}

func TestEnvProviderIsExpired(t *testing.T) {
	os.Clearenv()
	os.Setenv("CHECKPOINT_USER", "user")
	os.Setenv("CHECKPOINT_PASSWORD", "password")

	e := EnvProvider{}

	if !e.IsExpired() {
		t.Errorf("Expect creds to be expired before retrieve.")
	}

	_, err := e.Retrieve()
	if err != nil {
		t.Errorf("expect nil, got %v", err)
	}

	if e.IsExpired() {
		t.Errorf("Expect creds to not be expired after retrieve.")
	}
}

func TestEnvProviderNoUser(t *testing.T) {
	os.Clearenv()
	os.Setenv("CHECKPOINT_PASSWORD", "password")

	e := EnvProvider{}
	_, err := e.Retrieve()
	if e, a := ErrEnvUserNotFound, err; e != a {
		t.Errorf("expect %v, got %v", e, a)
	}
}

func TestEnvProviderNoPassword(t *testing.T) {
	os.Clearenv()
	os.Setenv("CHECKPOINT_USER", "user")

	e := EnvProvider{}
	_, err := e.Retrieve()
	if e, a := ErrEnvPasswordNotFound, err; e != a {
		t.Errorf("expect %v, got %v", e, a)
	}
}

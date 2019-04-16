package checkpoint

import (
	"github.com/janschumann/checkpoint-go-sdk/checkpoint/credentials"
	"net/http"
	"reflect"
	"testing"
)

var testCredentials = credentials.NewStaticCredentials("USER", "PASSWORD")

var mergeTestZeroValueConfig = Config{}

var mergeTestConfig = Config{
	Credentials: testCredentials,
	HTTPClient:  http.DefaultClient,
	ApiVersion:  String("MergeTestEndpoint"),
	ApiScheme:   String("MergeTestEndpoint"),
	ApiPort:     Int(80),
	ApiHost:     String("MergeTestEndpoint"),
	ApiContext:  String("MergeTestEndpoint"),
	Insecure:    Bool(false),
	SessionName: String("foo"),
}

var mergeTests = []struct {
	cfg  *Config
	in   *Config
	want *Config
}{
	{&Config{}, nil, &Config{}},
	{&Config{}, &mergeTestZeroValueConfig, &Config{}},
	{&Config{}, &mergeTestConfig, &mergeTestConfig},
}

func TestMerge(t *testing.T) {
	for i, tt := range mergeTests {
		got := tt.cfg
		got.MergeIn(tt.in)
		if !reflect.DeepEqual(got, tt.want) {
			t.Errorf("Config %d %+v", i, tt.cfg)
			t.Errorf("   Merge(%+v)", tt.in)
			t.Errorf("     got %+v", got)
			t.Errorf("    want %+v", tt.want)
		}
	}
}

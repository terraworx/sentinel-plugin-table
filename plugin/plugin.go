package table

import (
	"fmt"
	"os"
	"strings"

	sdk "github.com/hashicorp/sentinel-sdk"
	framework "github.com/hashicorp/sentinel-sdk/framework"
)

func New() sdk.Plugin {
	return &framework.Plugin{
		Root: &root{},
	}
}

type root struct{}

// framework.Root impl.
func (m *root) Configure(raw map[string]interface{}) error {
	return nil
}

// framework.Namespace impl.
func (m *root) Get(key string) (interface{}, error) {
	return nil, nil
}

// framework.Call impl.
func (m *root) Func(key string) interface{} {
	switch key {
	case "get":
		return func(key string) (interface{}, error) {
			val, ok := os.LookupEnv(key)
			if !ok {
				return nil, fmt.Errorf("%s is not a valid environment variable", key)
			}
			return val, nil
		}
	case "list":
		return func() (map[string]string, error) {
			env := make(map[string]string)

			for _, envVar := range os.Environ() {
				parts := strings.SplitN(envVar, "=", 2)
				if len(parts) == 2 {
					env[parts[0]] = parts[1]
				}
			}

			return env, nil
		}
	}
	return nil
}

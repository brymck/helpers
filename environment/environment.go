package validation

import (
	"fmt"
	"os"
)

// MustGetEnv retrieves the value of the environment variable named by the key.
// It returns the value and panics if the value is unset or blank.
func MustGetEnv(key string) string {
	value := os.Getenv(key)
	if value == "" {
		panic(fmt.Sprintf("environment variable %q not set", key))
	}
	return value
}

// MustMapEnv retrieves the value of the environment variable named by the key
// and assigns it to the target string reference. It panics if the value is
// unset or blank.
func MustMapEnv(target *string, key string) {
	value := MustGetEnv(key)
	*target = value
}

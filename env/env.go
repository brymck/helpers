package env

import (
	"fmt"
	"os"
)

// MustGetEnv retrieves the value of the env variable named by the key.
// It returns the value and panics if the value is unset or blank.
func MustGetEnv(key string) string {
	value := os.Getenv(key)
	if value == "" {
		panic(fmt.Sprintf("environment variable %q not set", key))
	}
	return value
}

// MustMapEnv retrieves the value of the env variable named by the key
// and assigns it to the target string reference. It panics if the value is
// unset or blank.
func MustMapEnv(target *string, key string) {
	value := MustGetEnv(key)
	*target = value
}

// GetPort retrieves the port from the PORT environment variable, falling back
// to the provided default value if that value is unset.
func GetPort(defaultPort string) string {
	envPort := os.Getenv("PORT")
	if envPort == "" {
		return defaultPort
	} else {
		return envPort
	}
}

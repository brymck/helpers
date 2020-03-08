package auth

import (
	"context"
	"os"
)

type apiKeyAuth struct {
	apiKey   string
	metadata map[string]string
}

func newApiKeyAuth() *apiKeyAuth {
	apiKey := os.Getenv("BRYMCK_IO_API_KEY")
	if apiKey == "" {
		panic("no value for environment variable BRYMCK_IO_API_KEY")
	}
	metadata := map[string]string{"x-api-key": apiKey, "x-goog-api-key": apiKey}
	return &apiKeyAuth{apiKey: apiKey, metadata: metadata}
}

func (a *apiKeyAuth) GetRequestMetadata(context.Context, ...string) (map[string]string, error) {
	return a.metadata, nil
}

func (*apiKeyAuth) RequireTransportSecurity() bool {
	return true
}

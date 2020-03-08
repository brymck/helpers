package auth

import (
	"os"

	"google.golang.org/grpc/credentials"
)

type auth struct {
	credentials.PerRPCCredentials
}

func isOnCloudRun() bool {
	return os.Getenv("K_SERVICE") != ""
}

func NewAuth(tokenUrl string) *auth {
	if isOnCloudRun() {
		return &auth{newMetadataTokenAuth(tokenUrl)}
	} else {
		if os.Getenv("BRYMCK_IO_API_KEY") != "" {
			return &auth{newApiKeyAuth()}
		} else {
			return &auth{newLocalTokenAuth()}
		}
	}
}

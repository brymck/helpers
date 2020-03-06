package services

import (
	"crypto/x509"
	"fmt"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"

	"github.com/brymck/helpers/internal/auth"
)

const (
	serviceAddressTemplate = "%s-4tt23pryoq-an.a.run.app:443"
	serviceUrlTemplate     = "https://%s-4tt23pryoq-an.a.run.app"
	tokenUrlTemplate       = "/instance/service-accounts/default/identity?audience=%s"
)

func Connect(serviceName string) (*grpc.ClientConn, error) {
	pool, _ := x509.SystemCertPool()
	ce := credentials.NewClientTLSFromCert(pool, "")

	audience := fmt.Sprintf(serviceUrlTemplate, serviceName)
	tokenUrl := fmt.Sprintf(tokenUrlTemplate, audience)
	creds := auth.NewAuth(tokenUrl)

	serviceAddress := fmt.Sprintf(serviceAddressTemplate, serviceName)
	conn, err := grpc.Dial(
		serviceAddress,
		grpc.WithTransportCredentials(ce),
		grpc.WithPerRPCCredentials(creds),
	)
	if err != nil {
		return nil, err
	}
	return conn, nil
}

func MustConnect(serviceName string) *grpc.ClientConn {
	conn, err := Connect(serviceName)
	if err != nil {
		panic(err)
	}
	return conn
}

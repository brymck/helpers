package servers

import (
	"fmt"
	"net"

	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	"github.com/brymck/helpers/env"
)

type GrpcServer struct {
	listener *net.Listener
	Server   *grpc.Server
}

func Listen() *GrpcServer {
	// Start server
	port := env.GetPort("8080")
	listener, err := net.Listen("tcp", fmt.Sprintf(":%s", port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	log.Infof("listening for gRPC on port %s", port)
	server := grpc.NewServer()
	return &GrpcServer{listener: &listener, Server: server}
}

func (ls *GrpcServer) Serve() {
	reflection.Register(ls.Server)
	err := ls.Server.Serve(*ls.listener)
	if err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

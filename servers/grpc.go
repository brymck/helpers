package servers

import (
	"fmt"
	"net"

	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	"github.com/brymck/helpers/env"
)

type ListenerServer struct {
	Listener *net.Listener
	Server *grpc.Server
}

func Listen() *ListenerServer {
	// Start server
	port := env.GetPort("8080")
	listener, err := net.Listen("tcp", fmt.Sprintf(":%s", port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	log.Infof("listening for gRPC on port %s", port)
	server := grpc.NewServer()
	return &ListenerServer{Listener: &listener, Server: server}
}

func Serve(ls *ListenerServer) {
	reflection.Register(ls.Server)
	err := ls.Server.Serve(*ls.Listener)
	if err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

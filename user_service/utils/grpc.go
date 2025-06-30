package utils

import (
	"net"

	"google.golang.org/grpc"
)

func GrpcServer(port string, options []grpc.ServerOption) (*grpc.Server, net.Listener, error) {
	// gServer, listener, error
	gServer := grpc.NewServer(options...)
	listener, err := net.Listen("tcp", port)

	return gServer, listener, err
}

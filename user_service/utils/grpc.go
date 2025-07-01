package utils

import (
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func GrpcServer(port string, options []grpc.ServerOption) (*grpc.Server, net.Listener, error) {
	// gServer, listener, error
	gServer := grpc.NewServer(options...)
	listener, err := net.Listen("tcp", port)

	return gServer, listener, err
}

func GrpcDial(port string) (*grpc.ClientConn, error) {
	conn, err := grpc.NewClient(port, grpc.WithTransportCredentials(insecure.NewCredentials()))
	return conn, err
}

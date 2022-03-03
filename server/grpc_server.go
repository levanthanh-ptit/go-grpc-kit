package server

import (
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"
)

// RegisterFunc registering handler type
type RegisterFunc func(server *grpc.Server) error

// GrpcServer server object
type GrpcServer struct {
	name string

	registerHandler RegisterFunc

	Server     *grpc.Server
	serverOpts []grpc.ServerOption
}

// NewGrpcServer constructor
func NewGrpcServer(name string) *GrpcServer {
	return &GrpcServer{
		name: name,
	}
}

// WithRegister injection method for gRPC server registration
func (s *GrpcServer) WithRegister(handler RegisterFunc) *GrpcServer {
	s.registerHandler = handler
	return s
}

// WithOptions injection method for gRPC server options, middlewares,...
func (s *GrpcServer) WithOptions(options ...grpc.ServerOption) *GrpcServer {
	s.serverOpts = append(s.serverOpts, options...)
	return s
}

// makeServer prepare gRPC server
func (s *GrpcServer) makeServer() (err error) {
	s.Server = grpc.NewServer(s.serverOpts...)
	// attach gRPC
	err = s.registerHandler(s.Server)
	return
}

// ServeTCP run server in TCP
func (s *GrpcServer) ServeTCP(host string, port int) (err error) {
	lis, err := net.Listen("tcp", fmt.Sprintf("%s:%v", host, port))
	if err != nil {
		log.Printf("%s gRPC Server - Failed to listen on: %s:%v", s.name, host, port)
		return
	}
	err = s.makeServer()
	if err != nil {
		log.Printf("%s gRPC Server - Failed to create gRPC server", s.name)
		return
	}
	log.Printf("%s gRPC Server - Started on %s:%v", s.name, host, port)
	err = s.Server.Serve(lis)
	return
}

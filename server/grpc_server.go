package server

import (
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"
)

type RegisterGrpcFunc func(server *grpc.Server) error

// GrpcServer server object
type GrpcServer struct {
	name string
	host string
	port string

	gprpcRegisterHandler RegisterGrpcFunc

	server     *grpc.Server
	serverOpts []grpc.ServerOption
}

// NewGrpcServer constructor
func NewGrpcServer(name string) *GrpcServer {
	return &GrpcServer{
		name: name,
	}
}

// WithHost add host
func (s *GrpcServer) WithHost(host string) *GrpcServer {
	s.host = host
	return s
}

// WithPort add host
func (s *GrpcServer) WithPort(port string) *GrpcServer {
	s.port = port
	return s
}

// WithGprpcRegister add host
func (s *GrpcServer) WithGprpcRegister(handler RegisterGrpcFunc) *GrpcServer {
	s.gprpcRegisterHandler = handler
	return s
}

// registerGrpc attach gRPC
func (s *GrpcServer) registerGrpc() (err error) {
	err = s.gprpcRegisterHandler(s.server)
	return
}

// makeServer prepare gRPC server
func (s *GrpcServer) makeServer() (err error) {
	s.server = grpc.NewServer(s.serverOpts...)
	err = s.registerGrpc()
	return
}

// ServeTCP run server in TCP
func (s *GrpcServer) ServeTCP() (err error) {
	lis, err := net.Listen("tcp", fmt.Sprintf("%s:%s", s.host, s.port))
	if err != nil {
		log.Printf("%s gRPC Server - Failed to listen on: %s:%s", s.name, s.host, s.port)
		return
	}
	err = s.makeServer()
	if err != nil {
		log.Println("gRPC Server - Failed to create gRPC server")
		return
	}
	log.Printf("%s gRPC Server - Started on %s:%s", s.name, s.host, s.port)
	err = s.server.Serve(lis)
	return
}

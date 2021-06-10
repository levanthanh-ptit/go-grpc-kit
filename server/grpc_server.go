package server

import (
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"
)

type RegisterGrpcFunc func(server *grpc.Server)

// GrpcServer server object
type GrpcServer struct {
	name string
	host string
	port string

	gprpcRegisterHandler RegisterGrpcFunc

	server *grpc.Server
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

// registerGrpc attach gRPC
func (s *GrpcServer) registerGrpc() {
	s.gprpcRegisterHandler(s.server)
}

// makeServer prepare gRPC server
func (s *GrpcServer) makeServer() {
	s.server = grpc.NewServer()
	s.registerGrpc()
}

// ServerTCP run server in TCP
func (s *GrpcServer) ServerTCP() {
	lis, err := net.Listen("tcp", fmt.Sprintf("%s:%s", s.host, s.port))
	if err != nil {
		log.Fatalln("User gRPC - Failed to listen:", err)
	}
	s.makeServer()
	log.Printf("User gRPC - Started on %s:%s", s.host, s.port)
	log.Fatalln(s.server.Serve(lis))
}

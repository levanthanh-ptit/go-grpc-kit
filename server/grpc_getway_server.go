package server

import (
	"fmt"
	"log"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/runtime"
)

// GrpcGetwayServer server object
type GrpcGetwayServer struct {
	name string
	host string
	port string

	gwmux  *runtime.ServeMux
	server *http.Server
}

// NewGrpcGetwayServer constructor
func NewGrpcGetwayServer(name string) *GrpcGetwayServer {
	return &GrpcGetwayServer{
		name: name,
	}
}

// WithHost add host
func (s *GrpcGetwayServer) WithHost(host string) *GrpcGetwayServer {
	s.host = host
	return s
}

// WithPort add host
func (s *GrpcGetwayServer) WithPort(port string) *GrpcGetwayServer {
	s.port = port
	return s
}

// RegisterGrpcClient attach gRPC
func (s *GrpcGetwayServer) RegisterGrpcClient() {
	log.Fatalln("Must register an gRPC client")
}

// makeServer prepare gRPC server
func (s *GrpcGetwayServer) makeServer() {
	s.RegisterGrpcClient()
	s.server = &http.Server{
		Addr:    fmt.Sprintf("%s:%s", s.host, s.port),
		Handler: s.gwmux,
	}
}

// Server run server
func (s *GrpcGetwayServer) Server() {
	s.makeServer()
	log.Printf("User gRPC-Gateway - Started on http://%s:%s", s.host, s.port)
	log.Fatalln(s.server.ListenAndServe())
}

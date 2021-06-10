package server

import (
	"fmt"
	"log"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
)

// GrpcGetwayServer server object
type GrpcGetwayServer struct {
	name string
	host string
	port string

	handler http.Handler
	server  *http.Server
}

// NewGrpcGetwayServer constructor
func NewGrpcGetwayServer(name string) *GrpcGetwayServer {
	gwmux := runtime.NewServeMux()
	return &GrpcGetwayServer{
		name:    name,
		handler: gwmux,
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

// GetServer getter
func (s *GrpcGetwayServer) GetServer() *http.Server {
	return s.server
}

// AddHandlerFunc type for add handler
type AddHandlerFunc func(http.Handler) http.Handler

// WithHandler add handler
func (s *GrpcGetwayServer) WithHandler(handlerFunc AddHandlerFunc) *GrpcGetwayServer {
	s.handler = handlerFunc(s.handler)
	return s
}

// WithChainHandler add handler
func (s *GrpcGetwayServer) WithChainHandler(handlerFuncs []AddHandlerFunc) *GrpcGetwayServer {
	for _, handlerFunc := range handlerFuncs {
		s.handler = handlerFunc(s.handler)
	}
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
		Handler: s.handler,
	}
}

// Server run server
func (s *GrpcGetwayServer) Server() {
	s.makeServer()
	log.Printf("User gRPC-Gateway - Started on http://%s:%s", s.host, s.port)
	log.Fatalln(s.server.ListenAndServe())
}

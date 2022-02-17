package server

import (
	"errors"
	"fmt"
	"log"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
)

// ClientRegisterFunc ...
type ClientRegisterFunc func(gwmux *runtime.ServeMux) error

// GrpcGatewayServer server object
type GrpcGatewayServer struct {
	name string
	host string
	port string

	clientRegisterHandler ClientRegisterFunc

	gwmux   *runtime.ServeMux
	handler http.Handler
	server  *http.Server
}

// NewGrpcGatewayServer constructor
func NewGrpcGatewayServer(name string) *GrpcGatewayServer {
	gwmux := runtime.NewServeMux()
	return &GrpcGatewayServer{
		name:    name,
		gwmux:   gwmux,
		handler: gwmux,
	}
}

// WithHost add host
func (s *GrpcGatewayServer) WithHost(host string) *GrpcGatewayServer {
	s.host = host
	return s
}

// WithPort add host
func (s *GrpcGatewayServer) WithPort(port string) *GrpcGatewayServer {
	s.port = port
	return s
}

// WithClientRegister add client register handler
func (s *GrpcGatewayServer) WithClientRegister(handler ClientRegisterFunc) *GrpcGatewayServer {
	s.clientRegisterHandler = handler
	return s
}

// GetServer getter
func (s *GrpcGatewayServer) GetServer() *http.Server {
	return s.server
}

// AddHandlerFunc type for add handler
type AddHandlerFunc func(http.Handler) http.Handler

// WithHTTPHandler add handler
func (s *GrpcGatewayServer) WithHTTPHandler(handlerFunc AddHandlerFunc) *GrpcGatewayServer {
	s.handler = handlerFunc(s.handler)
	return s
}

// WithChainHTTPHandler add handler
func (s *GrpcGatewayServer) WithChainHTTPHandler(handlerFuncs []AddHandlerFunc) *GrpcGatewayServer {
	for _, handlerFunc := range handlerFuncs {
		s.handler = handlerFunc(s.handler)
	}
	return s
}

// registerGrpcClient attach gRPC
func (s *GrpcGatewayServer) registerGrpcClient() (err error) {
	if s.clientRegisterHandler == nil {
		err = errors.New("must implement client register handler")
		return
	}
	err = s.clientRegisterHandler(s.gwmux)
	return
}

// makeServer prepare gRPC server
func (s *GrpcGatewayServer) makeServer() (err error) {
	err = s.registerGrpcClient()
	if err != nil {
		return
	}
	s.server = &http.Server{
		Addr:    fmt.Sprintf("%s:%s", s.host, s.port),
		Handler: s.handler,
	}
	return
}

// Serve run server
func (s *GrpcGatewayServer) Serve() (err error) {
	err = s.makeServer()
	if err != nil {
		return
	}
	log.Printf("%s gRPC Gateway - Started on %s:%s", s.name, s.host, s.port)
	err = s.server.ListenAndServe()
	return
}

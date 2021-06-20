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

// GrpcGetwayServer server object
type GrpcGetwayServer struct {
	name string
	host string
	port string

	clientRegisterHandler ClientRegisterFunc

	gwmux   *runtime.ServeMux
	handler http.Handler
	server  *http.Server
}

// NewGrpcGetwayServer constructor
func NewGrpcGetwayServer(name string) *GrpcGetwayServer {
	gwmux := runtime.NewServeMux()
	return &GrpcGetwayServer{
		name:    name,
		gwmux:   gwmux,
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

// WithClientRegister add client register handler
func (s *GrpcGetwayServer) WithClientRegister(handler ClientRegisterFunc) *GrpcGetwayServer {
	s.clientRegisterHandler = handler
	return s
}

// GetServer getter
func (s *GrpcGetwayServer) GetServer() *http.Server {
	return s.server
}

// AddHandlerFunc type for add handler
type AddHandlerFunc func(http.Handler) http.Handler

// WithHTTPHandler add handler
func (s *GrpcGetwayServer) WithHTTPHandler(handlerFunc AddHandlerFunc) *GrpcGetwayServer {
	s.handler = handlerFunc(s.handler)
	return s
}

// WithChainHTTPHandler add handler
func (s *GrpcGetwayServer) WithChainHTTPHandler(handlerFuncs []AddHandlerFunc) *GrpcGetwayServer {
	for _, handlerFunc := range handlerFuncs {
		s.handler = handlerFunc(s.handler)
	}
	return s
}

// registerGrpcClient attach gRPC
func (s *GrpcGetwayServer) registerGrpcClient() (err error) {
	if s.clientRegisterHandler == nil {
		err = errors.New("must implement client register handler")
		return
	}
	err = s.clientRegisterHandler(s.gwmux)
	return
}

// makeServer prepare gRPC server
func (s *GrpcGetwayServer) makeServer() (err error) {
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
func (s *GrpcGetwayServer) Serve() (err error) {
	err = s.makeServer()
	if err != nil {
		return
	}
	log.Printf("%s gRPC Gateway - Started on %s:%s", s.name, s.host, s.port)
	err = s.server.ListenAndServe()
	return
}

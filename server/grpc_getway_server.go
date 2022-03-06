package server

import (
	"errors"
	"fmt"
	"log"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
)

// ClientRegisterFunc injector function for runtime.ServeMux register
type ClientRegisterFunc func(gwmux *runtime.ServeMux) error

// GrpcGatewayServer server object
type GrpcGatewayServer struct {
	name string

	clientRegisterHandler ClientRegisterFunc

	gwmux   *runtime.ServeMux
	Handler http.Handler
	Server  *http.Server
}

// NewGrpcGatewayServer constructor
func NewGrpcGatewayServer(name string) *GrpcGatewayServer {
	gwmux := runtime.NewServeMux()
	return &GrpcGatewayServer{
		name:    name,
		gwmux:   gwmux,
		Handler: gwmux,
	}
}

// WithClientRegister injection method for gRPC client registration
func (s *GrpcGatewayServer) WithClientRegister(handler ClientRegisterFunc) *GrpcGatewayServer {
	s.clientRegisterHandler = handler
	return s
}

// AddHandlerFunc type for add handler
type AddHandlerFunc func(http.Handler) http.Handler

// WithHandlers add handler
func (s *GrpcGatewayServer) WithHandlers(handlerFuncs ...AddHandlerFunc) *GrpcGatewayServer {
	for _, handlerFunc := range handlerFuncs {
		s.Handler = handlerFunc(s.Handler)
	}
	return s
}

// registerGrpcClient attach gRPC
func (s *GrpcGatewayServer) registerGrpcClient() (err error) {
	if s.clientRegisterHandler == nil {
		err = errors.New("Must implement client register handler")
		return
	}
	err = s.clientRegisterHandler(s.gwmux)
	return
}

// makeServer prepare gRPC server
func (s *GrpcGatewayServer) makeServer(host string, port int) (err error) {
	err = s.registerGrpcClient()
	if err != nil {
		return
	}
	s.Server = &http.Server{
		Addr:    fmt.Sprintf("%s:%v", host, port),
		Handler: s.Handler,
	}
	return
}

// Serve run server
func (s *GrpcGatewayServer) Serve(host string, port int) (err error) {
	err = s.makeServer(host, port)
	if err != nil {
		return
	}
	log.Printf("%s gRPC Gateway - Started on %s:%v", s.name, host, port)
	err = s.Server.ListenAndServe()
	return
}

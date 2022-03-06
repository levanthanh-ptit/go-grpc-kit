# go-grpc-kit

## gRPC Server

An lightweight `gRPC Server` for embedding to lift off declaring basic of gRPC server.

### `GrpcServer` struct

- `name`: Name of the server.
- `registerHandler`: Injected function accepting a `*grpc.Server` as parameter to register.
- `Server`: `*grpc.Server` object.
- `serverOpts`: Options to pass to `server`.

### `GrpcServer` decorators methods

- `WithRegister`: For inject `registerHandler` function.
- `WithOptions`: Append options or middleware to `Server`.

### `GrpcServer` methods

- `ServeTCP`: start the server by HTTP2.

## gRPC Gateway

An lightweight `gRPC Gateway` for embedding to lift off declaring basic of gRPC getaway server.

### `GrpcGatewayServer` struct

- `name`: Name of the server.
- `clientRegisterHandler`: Injected function accepting a `*runtime.ServeMux` as parameter to gRPC client registration.
- `gwmux`: `*runtime.ServeMux` object.
- `Handler`: `http.Handler` HTTP handler, commonly to add middleware.
- `Server`: `http.Server` object.

### `GrpcGatewayServer` decorators methods

- `WithClientRegister`: For inject `clientRegisterHandler` function.
- `WithHandlers`: To add handlers to `Server`.

### `GrpcGatewayServer` methods

- `Serve`: start the server.

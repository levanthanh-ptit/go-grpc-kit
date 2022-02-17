# go-grpc-kit

## gRPC Server

An lightweight `gRPC Server` for embedding to lift off declaring basic of gRPC server.

### `GrpcServer` struct

- `name`: Name of the server.
- `host`: Host to bind on.
- `port`: Port to listen on
- `grpcRegisterHandler`: Provide an `*grpc.Server` as parameter to register.
- `server`: An `*grpc.Server` object.
- `serverOpts`: Options to pass to `server`.

### `GrpcServer` decorators methods

- `WithHost`
- `WithPort`
- `WithGrpcRegister`

### `GrpcServer` methods

- `ServeTCP`: start the server by HTTP2.

## gRPC Gateway

An lightweight `gRPC Gateway` for embedding to lift off declaring basic of gRPC getaway server.

### `GrpcGatewayServer` struct

- `name`: Name of the server.
- `host`: Host to bind on.
- `port`: Port to listen on
- `clientRegisterHandler`: Provide an `*runtime.ServeMux` as parameter to register.
- `gwmux`: An `*runtime.ServeMux` object.
- `handler`: An `http.Handler` HTTP handler, commonly to add middleware.
- `server`: An `http.Server` object.

### `GrpcGatewayServer` decorators methods

- `WithHost`
- `WithPort`
- `WithHTTPHandler`: To add an HTTP handler to server.
- `WithChainHTTPHandler`: To add multiple HTTP handlers to server.

### `GrpcGatewayServer` methods

- `Serve`: start the server.

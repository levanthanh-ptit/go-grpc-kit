# go-grpc-kit

## gRPC Server

An lightweight `gRPC Server` for embeding to lift off declaring basic of gRPC server.

### `GrpcServer` struct

- `name`: Name of the server.
- `host`: Host to bind on.
- `port`: Port to listen on
- `gprpcRegisterHandler`: Provide an `*grpc.Server` as parameter to register.
- `server`: An `*grpc.Server` object.
- `serverOpts`: Options to pass to `server`.

### `GrpcServer` decorators methods

- `WithHost`
- `WithPort`
- `WithGprpcRegister`

### `GrpcServer` methods

- `ServeTCP`: start the server by HTTP2.

## gRPC Getway

An lightweight `gRPC Getway` for embeding to lift off declaring basic of gRPC getway server.

### `GrpcGetwayServer` struct

- `name`: Name of the server.
- `host`: Host to bind on.
- `port`: Port to listen on
- `clientRegisterHandler`: Provide an `*runtime.ServeMux` as parameter to register.
- `gwmux`: An `*runtime.ServeMux` object.
- `handler`: An `http.Handler` HTTP handler, commonly to add middleware.
- `server`: An `http.Server` object.

### `GrpcGetwayServer` decorators methods

- `WithHost`
- `WithPort`
- `WithHTTPHandler`: To add an HTTP hadler to server.
- `WithChainHTTPHandler`: To add mutiple HTTP hadlers to server.

### `GrpcGetwayServer` methods

- `Serve`: start the server.

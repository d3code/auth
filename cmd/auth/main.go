package main

import (
    "context"
    "github.com/d3code/auth/generated/protobuf/v1/auth"
    "github.com/d3code/auth/internal/middleware"
    "github.com/d3code/auth/internal/service/auth_service"
    "github.com/d3code/auth/pkg/server"
    "github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
    "google.golang.org/grpc"
    "net/http"
)

func main() {
    host := "localhost"

    httpPort := "8080"
    grpcPort := "8081"

    grpcServer := server.GrpcServer{
        Host: host,
        Port: grpcPort,
        RegisterServices: func(server *grpc.Server) {
            auth.RegisterAuthServiceServer(server, &auth_service.AuthService{})
        },
    }

    httpGateway := server.HttpGateway{
        Port: httpPort,
        GrpcConnections: map[string]server.GrpcConnection{
            "/auth/": {
                Host:   host,
                Port:   grpcPort,
                Secure: false,
                GrpcHandlers: []func(ctx context.Context, mux *runtime.ServeMux, conn *grpc.ClientConn) error{
                    auth.RegisterAuthServiceHandler,
                },
            },
        },
        HttpHandlers: map[string]http.Handler{
            "/health": middleware.ServerHealth(),
        },
    }

    ctx, cancel := context.WithCancel(context.Background())
    defer cancel()

    go httpGateway.Run(ctx)
    grpcServer.Run()
}

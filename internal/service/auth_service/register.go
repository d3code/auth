package auth_service

import (
    "context"
    "github.com/d3code/auth/generated/protobuf/v1/auth"
    "google.golang.org/grpc/codes"
    "google.golang.org/grpc/status"
)

func (AuthService) Register(context.Context, *auth.RegisterRequest) (*auth.JwtToken, error) {
    return nil, status.Errorf(codes.Unimplemented, "method Register not implemented")
}

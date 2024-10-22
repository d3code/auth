package service

import (
    "context"
    "github.com/d3code/auth/generated/protobuf/v1/auth"
    "google.golang.org/grpc/codes"
    "google.golang.org/grpc/status"
)

func (AuthService) ForgotPassword(context.Context, *auth.ForgotPasswordRequest) (*auth.ForgotPasswordResponse, error) {
    return nil, status.Errorf(codes.Unimplemented, "Not implemented")
}

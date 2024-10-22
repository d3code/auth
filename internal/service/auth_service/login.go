package auth_service

import (
    "context"
    "database/sql"
    "errors"
    "github.com/d3code/auth/generated/protobuf/v1/auth"
    "github.com/d3code/auth/internal/db"
    "github.com/d3code/auth/internal/util"
    "github.com/d3code/zlog"
    "google.golang.org/grpc/codes"
    "google.golang.org/grpc/status"
    "net/http"
)

type AuthService struct {
    auth.AuthServiceServer
}

func (AuthService) Login(c context.Context, loginRequest *auth.LoginRequest) (*auth.JwtToken, error) {
    zlog.Log.Infof("Login request: %+v", loginRequest)

    // Get User from database
    user, loginError := db.GetUser(loginRequest.GetUsername())
    if loginError != nil {
        if errors.Is(loginError, sql.ErrNoRows) {
            zlog.Log.Warnf("No account [%s]", loginRequest.GetUsername())
            return nil, status.Errorf(codes.PermissionDenied, "Invalid username or password")
        }
        zlog.Log.Error(loginError)
        return nil, status.Errorf(http.StatusInternalServerError, "Unknown error")
    }

    // Match password
    if !util.PasswordMatch(loginRequest.GetPassword(), user.Password) {
        zlog.Log.Warnf("Password incorrect [%s]", user.Username)
        return nil, status.Errorf(codes.PermissionDenied, "Invalid username or password")
    }

    // Create token from account
    response, hasError := db.CreateTokenForUser(user)
    if hasError != nil {
        zlog.Log.Error(hasError)
        return nil, status.Errorf(codes.Internal, "Unknown error")
    }

    zlog.Log.Infof("Login successful [%s]", user.Username)

    return response, nil
}

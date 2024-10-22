package service

import (
    "context"
    "database/sql"
    "encoding/json"
    "errors"
    "github.com/MicahParks/keyfunc"
    "github.com/d3code/auth/generated/protobuf/v1/auth"
    "github.com/d3code/auth/internal/db"
    "github.com/d3code/zlog"
    "github.com/golang-jwt/jwt/v4"
    "google.golang.org/grpc/codes"
    "google.golang.org/grpc/status"
    "net/http"
)

func (AuthService) Refresh(c context.Context, r *auth.RefreshRequest) (*auth.JwtToken, error) {
    headerToken := r.GetRefreshToken()
    jwks := GetJwks()
    if jwks == nil {
        zlog.Log.Error("Missing JWKS")
        return nil, status.Errorf(codes.Internal, "Missing JWKS")
    }

    jwksString, err := json.Marshal(jwks)
    if err != nil {
        zlog.Log.Error(err)
        return nil, err
    }

    jwksRaw := json.RawMessage(jwksString)
    jwksJson, jsonError := keyfunc.NewJSON(jwksRaw)
    if jsonError != nil {
        zlog.Log.Error(jsonError)
        return nil, jsonError
    }

    token, jwtParseError := jwt.Parse(headerToken, jwksJson.Keyfunc)
    if jwtParseError != nil {
        zlog.Log.Error(jwtParseError)
        return nil, status.Errorf(codes.PermissionDenied, jwtParseError.Error())
    }

    claims, claimsError := db.GetClaims(*token)
    if claimsError != nil {
        zlog.Log.Error(claimsError)
        return nil, claimsError
    }

    user, loginError := db.GetUserById(claims["sub"].(string))
    if loginError != nil {
        if errors.Is(loginError, sql.ErrNoRows) {
            zlog.Log.Warnf("No account [%s]", r.GetRefreshToken())
            return nil, status.Errorf(codes.PermissionDenied, "Invalid username or password")
        }
        zlog.Log.Error(loginError)
        return nil, status.Errorf(http.StatusInternalServerError, "Unknown error")
    }

    // Create token from account
    response, hasError := db.CreateTokenForUser(user)
    if hasError != nil {
        zlog.Log.Error(hasError)
        return nil, status.Errorf(codes.Internal, "Unknown error")
    }

    return response, nil
}

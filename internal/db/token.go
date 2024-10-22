package db

import (
    "crypto/rsa"
    "errors"
    "github.com/d3code/auth/generated/protobuf/v1/auth"
    "github.com/d3code/auth/internal/config"
    "github.com/d3code/auth/internal/model"
    "github.com/d3code/auth/pkg/encrypt"
    "github.com/d3code/zlog"
    "github.com/golang-jwt/jwt/v4"
    "github.com/google/uuid"
    "google.golang.org/protobuf/types/known/timestamppb"
    "time"
)

func GetClaims(token jwt.Token) (jwt.MapClaims, error) {
    if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
        return claims, nil
    }
    return nil, errors.New("token not valid")
}

func CreateTokenForUser(account *model.User) (*auth.JwtToken, error) {
    secrets, getSecretsError := GetSecrets()
    if getSecretsError != nil {
        return nil, getSecretsError
    }

    if secrets == nil || len(secrets) == 0 {
        createSecret, createSecretError := CreateSecret()
        if createSecretError != nil {
            zlog.Log.Errorf("Error creating secret [%s]", createSecretError)
            return nil, createSecretError
        }
        secrets = createSecret
    }

    now := time.Now()
    secret := secrets[0]
    privateKey := encrypt.RsaPrivateFromString(secret.KeyPrivate)

    token, err := createToken(account, false, now, secret, privateKey)
    if err != nil {
        return nil, err
    }
    tokenRefresh, err := createToken(account, true, now, secret, privateKey)
    if err != nil {
        return nil, err
    }

    expiration := time.Duration(config.Environment().Token.Expiration)
    var expiresAt = now.Add(time.Second * expiration)

    timestamp := timestamppb.Timestamp{
        Seconds: expiresAt.Unix(),
        Nanos:   int32(expiresAt.Nanosecond()),
    }

    response := auth.JwtToken{
        AccessToken:  *token,
        RefreshToken: *tokenRefresh,
        TokenType:    "Bearer",
        ExpiresIn:    int32(expiration),
        ExpiresAt:    &timestamp,
    }

    // TODO: Store token in database
    return &response, nil
}

func createToken(account *model.User, refresh bool, now time.Time, secret model.Secret, privateKey *rsa.PrivateKey) (*string, error) {
    scope := account.Scope
    expiration := time.Duration(config.Environment().Token.Expiration)

    if refresh {
        scope = "refresh"
        expiration = time.Duration(config.Environment().Token.ExpirationRefresh)
    }

    var expiresAt = now.Add(time.Second * expiration)
    claims := &CustomClaims{
        RegisteredClaims: &jwt.RegisteredClaims{
            ExpiresAt: jwt.NewNumericDate(expiresAt),
            ID:        uuid.New().String(),
            IssuedAt:  jwt.NewNumericDate(now),
            NotBefore: jwt.NewNumericDate(now),
            Issuer:    config.Environment().Token.Issuer,
            Subject:   account.Id,
            Audience:  []string{config.Environment().Token.Audience},
        },
        Scope: scope,
    }

    token := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)
    token.Header["kid"] = secret.Id
    signedString, signedKeyError := token.SignedString(privateKey)
    if signedKeyError != nil {
        zlog.Log.Error(signedKeyError)
        return nil, signedKeyError
    }

    return &signedString, nil
}

type CustomClaims struct {
    *jwt.RegisteredClaims
    Scope string `json:"scope"`
}

type AccessTokenResponse struct {
    AccessToken string        `json:"access_token"`
    TokenType   string        `json:"token_type"`
    ExpiresIn   time.Duration `json:"expires_in"`
    ExpiresAt   int64         `json:"expires_at"`
    Scope       string        `json:"scope"`
}

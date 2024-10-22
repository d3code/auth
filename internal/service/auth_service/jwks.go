package auth_service

import (
    "context"
    "encoding/json"
    "github.com/d3code/auth/generated/protobuf/v1/auth"
    "github.com/d3code/auth/internal/db"
    "github.com/d3code/auth/pkg/encrypt"
    "github.com/d3code/zlog"
    "github.com/go-jose/go-jose/v3"
    "google.golang.org/grpc/codes"
    "google.golang.org/grpc/status"
)

type JWKSJson struct {
    Keys []struct {
        N   string `json:"n"`
        Kid string `json:"kid"`
        E   string `json:"e"`
        Kty string `json:"kty"`
    } `json:"keys"`
}

func (AuthService) Jwks(c context.Context, r *auth.EmptyRequest) (*auth.JwksResponse, error) {
    jwks := GetJwks()
    if jwks == nil {
        return nil, status.Errorf(codes.Internal, "unable to retrieve JWKS")
    }

    jwksBytes, _ := json.Marshal(jwks)

    var jwksJson JWKSJson
    err := json.Unmarshal(jwksBytes, &jwksJson)
    if err != nil {
        zlog.Log.Error(err)
        return nil, err
    }

    var keys []*auth.Jwks
    for _, key := range jwksJson.Keys {

        keys = append(keys, &auth.Jwks{
            Kty: key.Kty,
            Kid: key.Kid,
            N:   key.N,
            E:   key.E,
        })
    }

    return &auth.JwksResponse{
        Keys: keys,
    }, nil
}

func GetJwks() *jose.JSONWebKeySet {
    secrets, err := db.GetSecrets()
    if err != nil {
        zlog.Log.Errorf("Error getting secrets [%s]", err)
        return nil
    }

    if secrets == nil || len(secrets) == 0 {
        createSecret, createSecretError := db.CreateSecret()
        if createSecretError != nil {
            zlog.Log.Errorf("Error creating secret [%s]", createSecretError)
            return nil
        }
        secrets = createSecret
    }

    var jwkList []jose.JSONWebKey
    for _, secret := range secrets {
        privateKey := encrypt.RsaPrivateFromString(secret.KeyPrivate)
        publicKey := privateKey.PublicKey

        key := jose.JSONWebKey{
            Key:   &publicKey,
            KeyID: secret.Id,
        }

        jwkList = append(jwkList, key)
    }

    jsonWebKeySet := jose.JSONWebKeySet{
        Keys: jwkList,
    }

    return &jsonWebKeySet
}

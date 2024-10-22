package db

import (
    _ "embed"
    "github.com/d3code/auth/internal/config"
    "github.com/d3code/auth/internal/model"
    "github.com/d3code/pkg/encrypt"
    "github.com/d3code/zlog"
    "github.com/google/uuid"
)

//go:embed sql/secret/secret-insert.sql
var SqlSecretInsert string

func CreateSecret() ([]model.Secret, error) {
    connection := config.DatabaseConnection()

    privateKey, _ := encrypt.RsaGenerate()
    privateKeyString := encrypt.RsaPrivateToString(privateKey)

    id := uuid.New().String()
    _, insertError := connection.Exec(SqlSecretInsert, id, privateKeyString)
    if insertError != nil {
        zlog.Log.Error(insertError)
        return nil, insertError
    }

    zlog.Log.Info("Created secret")
    return GetSecrets()
}

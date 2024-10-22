package db

import (
    _ "embed"
    "github.com/d3code/auth/internal/config"
    "github.com/d3code/auth/internal/model"
    "github.com/d3code/auth/internal/util"
    "github.com/d3code/auth/pkg/encrypt"
    "github.com/d3code/zlog"
    "github.com/google/uuid"
)

func GetSecrets() ([]model.Secret, error) {
    connection := config.DatabaseConnection()

    rows, selectError := connection.Query("SELECT BIN_TO_UUID(id) AS id, private_key FROM secret")
    defer util.CloseRows(rows)
    if selectError != nil {
        zlog.Log.Error(selectError)
        return nil, selectError
    }

    var secrets []model.Secret
    for rows.Next() {
        var secret model.Secret
        scanError := rows.Scan(&secret.Id, &secret.KeyPrivate)
        if scanError != nil {
            zlog.Log.Error(scanError)
            continue
        }
        secrets = append(secrets, secret)
    }

    return secrets, nil
}

func GetSecret(id string) (*model.Secret, error) {
    connection := config.DatabaseConnection()

    rows, selectError := connection.Query("SELECT BIN_TO_UUID(id) AS id, private_key FROM secret WHERE id = ?", id)
    defer util.CloseRows(rows)

    if selectError != nil {
        zlog.Log.Error(selectError)
        return nil, selectError
    }

    for rows.Next() {
        var secret model.Secret
        scanError := rows.Scan(&secret.Id, &secret.KeyPrivate)
        if scanError != nil {
            zlog.Log.Error(scanError)
            continue
        }
        return &secret, nil
    }

    return nil, nil
}

func CreateSecret() ([]model.Secret, error) {
    connection := config.DatabaseConnection()

    privateKey, _ := encrypt.RsaGenerate()
    privateKeyString := encrypt.RsaPrivateToString(privateKey)

    id := uuid.New().String()
    _, insertError := connection.Exec("INSERT INTO secret(id, private_key) VALUES (UUID_TO_BIN(?), ?)", id, privateKeyString)
    if insertError != nil {
        zlog.Log.Error(insertError)
        return nil, insertError
    }

    zlog.Log.Info("Created secret")
    return GetSecrets()
}

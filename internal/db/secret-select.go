package db

import (
    _ "embed"
    "github.com/d3code/auth/internal/config"
    "github.com/d3code/auth/internal/model"
    "github.com/d3code/auth/internal/util"
    "github.com/d3code/zlog"
)

//go:embed sql/secret/secret-select.sql
var SqlSecretSelect string

//go:embed sql/secret/secret-select-id.sql
var SqlSecretSelectId string

func GetSecrets() ([]model.Secret, error) {
    connection := config.DatabaseConnection()

    rows, selectError := connection.Query(SqlSecretSelect)
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

    rows, selectError := connection.Query(SqlSecretSelectId, id)
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

package db

import (
    "github.com/d3code/auth/internal/config"
    "github.com/d3code/auth/internal/model"
)

func GetUser(username string) (*model.User, error) {
    connection := config.DatabaseConnection()

    var retrievedAccount model.User
    result := connection.QueryRow("SELECT BIN_TO_UUID(id), username, password, scope, active, created FROM view_user WHERE username = ?", username)

    scanError := result.Scan(&retrievedAccount.Id, &retrievedAccount.Username, &retrievedAccount.Password, &retrievedAccount.Scope, &retrievedAccount.Active, &retrievedAccount.Created)
    if scanError != nil {
        return nil, scanError
    }

    return &retrievedAccount, nil
}

func GetUserById(id string) (*model.User, error) {
    connection := config.DatabaseConnection()

    var retrievedAccount model.User
    result := connection.QueryRow("SELECT BIN_TO_UUID(id), username, password, scope, active, created FROM view_user WHERE id = UUID_TO_BIN(?)", id)

    scanError := result.Scan(&retrievedAccount.Id, &retrievedAccount.Username, &retrievedAccount.Password, &retrievedAccount.Scope, &retrievedAccount.Active, &retrievedAccount.Created)
    if scanError != nil {
        return nil, scanError
    }

    return &retrievedAccount, nil
}

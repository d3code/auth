package db

import (
    "github.com/d3code/auth/internal/config"
    "github.com/d3code/auth/internal/model"
    "github.com/d3code/zlog"
    "golang.org/x/crypto/bcrypt"
)

func CreateUser(user model.User) (*model.User, error) {
    connection := config.DatabaseConnection()

    // Hash Password
    hashedPassword, bcryptError := bcrypt.GenerateFromPassword([]byte(user.Password), 12)
    if bcryptError != nil {
        zlog.Log.Error(bcryptError)
        return nil, bcryptError
    }

    var createdAccount model.User
    result := connection.QueryRow("CALL user_create(?, ?)", user.Username, hashedPassword)

    scanError := result.Scan(&createdAccount.Id, &createdAccount.Username, &createdAccount.Password, &createdAccount.Scope, &createdAccount.Created)
    if scanError != nil {
        zlog.Log.Error(scanError)
        return nil, scanError
    }

    return &createdAccount, nil
}

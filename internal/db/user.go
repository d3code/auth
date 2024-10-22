package db

import (
    "github.com/d3code/auth/internal/config"
    "github.com/d3code/auth/internal/model"
    "github.com/d3code/zlog"
    "github.com/google/uuid"
    "golang.org/x/crypto/bcrypt"
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

func CreateUser(user model.User) (*model.User, error) {
    connection := config.DatabaseConnection()

    // Generate UUID
    id, err := uuid.NewUUID()
    if err != nil {
        zlog.Log.Error(err)
        return nil, err
    }

    // Set ID
    user.Id = id.String()

    // Hash Password
    hashedPassword, bcryptError := bcrypt.GenerateFromPassword([]byte(user.Password), 12)
    if bcryptError != nil {
        zlog.Log.Error(bcryptError)
        return nil, bcryptError
    }

    // Insert User
    result, err := connection.Exec("INSERT INTO user (id, username, password, name_family, name_given) VALUES (UUID_TO_BIN(?), ?, ?, ?, ?)", user.Id, user.Username, hashedPassword, user.NameFamily, user.NameGiven)
    if err != nil {
        zlog.Log.Error(err)
        return nil, err
    }

    rowsAffected, err := result.RowsAffected()
    if err != nil {
        zlog.Log.Error(err)
        return nil, err
    }

    if rowsAffected == 0 {
        return nil, nil
    }

    return GetUserById(user.Id)
}

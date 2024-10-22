package db

import (
    "github.com/d3code/auth/internal/config"
    "github.com/d3code/auth/internal/model"
    "github.com/google/uuid"
    "golang.org/x/crypto/bcrypt"
)

func GetUser(username string) (*model.User, error) {
    connection := config.DatabaseConnection()

    var user model.User
    result := connection.QueryRow("SELECT BIN_TO_UUID(id), username, password, email, scope, active, created FROM view_user WHERE username = ?", username)

    scanError := result.Scan(&user.Id, &user.Username, &user.Password, &user.Email, &user.Scope, &user.Active, &user.Created)
    if scanError != nil {
        return nil, scanError
    }

    return &user, nil
}

func GetUserById(id string) (*model.User, error) {
    connection := config.DatabaseConnection()

    var user model.User
    result := connection.QueryRow("SELECT BIN_TO_UUID(id), username, password, email, scope, active, created FROM view_user WHERE id = UUID_TO_BIN(?)", id)

    scanError := result.Scan(&user.Id, &user.Username, &user.Password, &user.Email, &user.Scope, &user.Active, &user.Created)
    if scanError != nil {
        return nil, scanError
    }

    return &user, nil
}

func CreateUser(user model.User) (*model.User, error) {
    connection := config.DatabaseConnection()

    id, err := uuid.NewUUID()
    if err != nil {
        return nil, err
    }

    hashedPassword, bcryptError := bcrypt.GenerateFromPassword([]byte(user.Password), 12)
    if bcryptError != nil {
        return nil, bcryptError
    }

    user.Id = id.String()
    result, err := connection.Exec("INSERT INTO user (id, username, password, email, name_family, name_given) VALUES (UUID_TO_BIN(?), ?, ?, ?, ?)", user.Id, user.Username, hashedPassword, user.Email, user.NameFamily, user.NameGiven)
    if err != nil {
        return nil, err
    }

    rowsAffected, err := result.RowsAffected()
    if err != nil {
        return nil, err
    }

    if rowsAffected == 0 {
        return nil, nil
    }

    return GetUserById(user.Id)
}

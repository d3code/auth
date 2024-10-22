package config

import (
    "database/sql"
    "fmt"
    "github.com/d3code/pkg/cfg"
    "github.com/d3code/zlog"
    _ "github.com/go-sql-driver/mysql"
    "sync"
)

var db *sql.DB

func DatabaseConnection() *sql.DB {
    if db == nil {
        db = connect()
    }

    return db
}

func connect() *sql.DB {
    database := config()
    var (
        user           = database.User           // root
        password       = database.Password       // password
        connectionType = database.ConnectionType // tcp
        databaseName   = database.DatabaseName   // database
        port           = database.Port           // 3306
        host           = database.Host           // 127.0.0.1
        unixSocketPath = database.ConnectionName // cloudsql/{project}:{region}:{instance}
    )

    connection := buildConnection(connectionType, host, port, unixSocketPath)
    connectionString := fmt.Sprintf("%s:%s@%s/%s?parseTime=true", user, password, connection, databaseName)

    open, err := sql.Open("mysql", connectionString)
    if err != nil {
        zlog.Log.Fatal(err)
    }

    return open
}

func buildConnection(connectionType string, host string, port string, unixSocketPath string) string {
    if connectionType == "tcp" {
        return fmt.Sprintf("tcp(%s:%s)", host, port)
    } else if connectionType == "unix" {
        return fmt.Sprintf("unix(/%s)", unixSocketPath)
    }

    zlog.Log.Fatalf("Invalid connection_type [ %s ]", connectionType)
    return ""
}

var (
    configDatabase DatabaseConfig
    onceDatabase   sync.Once
)

type DatabaseConfig struct {
    ConnectionType string `yaml:"connection_type"`
    User           string `yaml:"user"`
    Password       string `yaml:"password"`
    Host           string `yaml:"host"`
    Port           string `yaml:"port"`
    ConnectionName string `yaml:"connection_name"`
    DatabaseName   string `yaml:"database_name"`
}

func config() DatabaseConfig {
    onceDatabase.Do(func() {
        cfg.LoadConfiguration("database", &configDatabase)
    })

    return configDatabase
}

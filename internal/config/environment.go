package config

import (
    "github.com/d3code/zlog"
    "sync"
)

var (
    configEnvironment EnvironmentConfig
    onceEnvironment   sync.Once
)

func Environment() EnvironmentConfig {
    onceEnvironment.Do(func() {
        err := loadConfiguration("", &configEnvironment)
        if err != nil {
            zlog.Log.Fatalf("Error loading configuration: %s", err.Error())
        }
    })
    return configEnvironment
}

type EnvironmentConfig struct {
    Database Database `yaml:"database"`
    Register Register `yaml:"register"`
    Mail     Mail     `yaml:"mail"`
    Token    Token    `yaml:"token"`
}

type Database struct {
    ConnectionType string `yaml:"connection_type"`
    User           string `yaml:"user"`
    Password       string `yaml:"password"`
    Host           string `yaml:"host"`
    Port           string `yaml:"port"`
    ConnectionName string `yaml:"connection_name"`
    DatabaseName   string `yaml:"database_name"`
}

type Token struct {
    Issuer            string `yaml:"issuer"`
    IssuerGrpc        string `yaml:"issuer_grpc"`
    Audience          string `yaml:"audience"`
    Expiration        int    `yaml:"expiration"`
    ExpirationRefresh int    `yaml:"expiration_refresh"`
}

type Mail struct {
    SendgridApiKey   string `yaml:"sendgrid_api_key"`
    TemplateLocation string `yaml:"template_location"`
}

type Register struct {
    SendEmail             bool     `yaml:"send_email"`
    ValidateUsernameEmail bool     `yaml:"validate_username_email"`
    Restrict              bool     `yaml:"restrict"`
    AllowedAccounts       []string `yaml:"allowed_accounts"`
}

package config

import (
    "gopkg.in/yaml.v3"
    "os"
    "path"
    "regexp"
    "strings"
)

func loadConfiguration(name string, config interface{}) error {
    environment := getEnvironmentOrDefault("ENVIRONMENT", "local")
    configLocation := getEnvironmentOrDefault("CONFIG_LOCATION", "config")

    var configFilename string
    if name != "" {
        configFilename = environment + "_" + name + ".yaml"
    } else {
        configFilename = environment + ".yaml"
    }

    configPath := path.Join(configLocation, configFilename)
    configFile, err := os.ReadFile(configPath)
    if err != nil {
        return err
    }

    configFile = environmentTemplate(configFile)
    err = yaml.Unmarshal(configFile, config)
    return err
}

func getEnvironmentOrDefault(key string, fallback string) string {
    if value, ok := os.LookupEnv(key); ok {
        return value
    }
    return fallback
}

func environmentTemplate(input []byte) []byte {
    config := string(input)

    re := regexp.MustCompile(`[{]{2}([^{|}]*)[}]{2}`)
    matches := re.FindAllStringSubmatch(config, -1)

    for _, match := range matches {
        text := strings.TrimSpace(match[1])
        if value, found := os.LookupEnv(text); found {
            config = strings.ReplaceAll(config, match[0], value)
        } else {
            config = strings.ReplaceAll(config, match[0], text)
        }
    }

    return []byte(config)
}

package config

import (
	"github.com/spf13/viper"
	"strings"
)

// Config - config storage
type Config struct {
	DB   Database
	HTTP HTTP
}

// Database - database config storage
type Database struct {
	Username string
	Password string
	Host     string
	Port     string
	Database string
	SSLMode  string
}

// HTTP - http server config storage
type HTTP struct {
	Host string
	Port string
}

// Get - read config and return as Config struct
func Get(configPath, envPrefix string) (Config, error) {
	err := read(configPath, envPrefix)
	if err != nil {
		return Config{}, err
	}

	conf := Config{
		DB: Database{
			Username: viper.GetString("db.username"),
			Password: viper.GetString("db.password"),
			Host:     viper.GetString("db.host"),
			Port:     viper.GetString("db.port"),
			Database: viper.GetString("db.database"),
			SSLMode:  viper.GetString("db.sslmode"),
		},
		HTTP: HTTP{
			Host: viper.GetString("http.host"),
			Port: viper.GetString("http.port"),
		},
	}

	return conf, nil
}

// read - reads config from file and environment
// configPath - filepath to config file
// envPrefix - application specific environment prefix
func read(configPath, envPrefix string) error {
	viper.SetConfigFile(configPath)
	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	viper.SetEnvPrefix(envPrefix)
	return viper.ReadInConfig()
}

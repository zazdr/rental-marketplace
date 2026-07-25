package config

import (
	"errors"
	"os"
)

type Config struct {
	Server struct {
		Host string
		Port string
	}
	Db struct {
		URL string
	}
}

func New() (*Config, error) {
	config := &Config{}

	config.Server.Host = os.Getenv("APP_SERVER_HOST")
	if config.Server.Host == "" {
		return nil, errors.New("[config] APP_SERVER_HOST is not set")
	}

	config.Server.Port = os.Getenv("APP_SERVER_PORT")
	if config.Server.Port == "" {
		return nil, errors.New("[config] APP_SERVER_PORT is not set")
	}

	config.Db.URL = os.Getenv("APP_DB_URL")
	if config.Db.URL == "" {
		return nil, errors.New("[config] APP_DB_URL is not set")
	}

	return config, nil
}

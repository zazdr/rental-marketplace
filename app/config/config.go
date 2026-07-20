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

	config.Server.Host = os.Getenv("SERVER_HOST")
	if config.Server.Host == "" {
		return nil, errors.New("[config] SERVER_HOST is not set")
	}

	config.Server.Port = os.Getenv("SERVER_PORT")
	if config.Server.Port == "" {
		return nil, errors.New("[config] SERVER_PORT is not set")
	}

	config.Db.URL = os.Getenv("DB_URL")
	if config.Db.URL == "" {
		return nil, errors.New("[config] DB_URL is not set")
	}

	return config, nil
}

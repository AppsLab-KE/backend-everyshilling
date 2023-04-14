package config

import (
	"fmt"
	"os"
)

type Postgres struct {
	Port     string
	Host     string
	User     string
	Password string
	Database string
	TimeZone string
}

type Server struct {
	Port string
}

type Config struct {
	Server   Server
	Postgres Postgres
}

func LoadConfig() (*Config, error) {
	var config Config
	var err error

	config.Server.Port = os.Getenv("SERVER_PORT")
	config.Postgres.Port = os.Getenv("POSTGRES_PORT")
	config.Postgres.Host = os.Getenv("POSTGRES_HOST")
	config.Postgres.User = os.Getenv("POSTGRES_USER")
	config.Postgres.Password = os.Getenv("POSTGRES_PASSWORD")
	config.Postgres.Database = os.Getenv("POSTGRES_DATABASE")
	config.Postgres.TimeZone = os.Getenv("POSTGRES_TIMEZONE")

	if config.Server.Port == "" {
		err = fmt.Errorf("SERVER_PORT environment variable not set")
	}
	if config.Postgres.Port == "" {
		err = fmt.Errorf("POSTGRES_PORT environment variable not set")
	}
	if config.Postgres.Host == "" {
		err = fmt.Errorf("POSTGRES_HOST environment variable not set")
	}
	if config.Postgres.User == "" {
		err = fmt.Errorf("POSTGRES_USER environment variable not set")
	}
	if config.Postgres.Password == "" {
		err = fmt.Errorf("POSTGRES_PASSWORD environment variable not set")
	}
	if config.Postgres.Database == "" {
		err = fmt.Errorf("POSTGRES_DATABASE environment variable not set")
	}
	if config.Postgres.TimeZone == "" {
		err = fmt.Errorf("POSTGRES_TIMEZONE environment variable not set")
	}

	return &config, err
}

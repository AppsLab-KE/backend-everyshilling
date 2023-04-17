package config

import (
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
	var ok bool
	var missingConfigs []string

	config.Server.Port = os.Getenv("SERVER_PORT")
	config.Postgres.Port = os.Getenv("POSTGRES_PORT")
	config.Postgres.Host = os.Getenv("POSTGRES_HOST")
	config.Postgres.User = os.Getenv("POSTGRES_USER")
	config.Postgres.Password = os.Getenv("POSTGRES_PASSWORD")
	config.Postgres.Database = os.Getenv("POSTGRES_DATABASE")
	config.Postgres.TimeZone = os.Getenv("POSTGRES_TIMEZONE")

	if config.Server.Port, ok = os.LookupEnv("SERVER_PORT"); !ok {
		missingConfigs = append(missingConfigs, "SERVER_PORT")
	}

	if config.Postgres.Port, ok = os.LookupEnv("POSTGRES_PORT"); !ok {
		missingConfigs = append(missingConfigs, "POSTGRES_PORT")
	}

	if config.Postgres.Port, ok = os.LookupEnv("POSTGRES_HOST"); !ok {
		missingConfigs = append(missingConfigs, "POSTGRES_HOST")
	}

	if config.Postgres.Port, ok = os.LookupEnv("POSTGRES_USER"); !ok {
		missingConfigs = append(missingConfigs, "POSTGRES_USER")
	}

	if config.Postgres.Port, ok = os.LookupEnv("POSTGRES_PASSWORD"); !ok {
		missingConfigs = append(missingConfigs, "POSTGRES_PASSWORD")
	}

	if config.Postgres.Port, ok = os.LookupEnv("POSTGRES_DATABASE"); !ok {
		missingConfigs = append(missingConfigs, "POSTGRES_DATABASE")
	}

	if config.Postgres.Port, ok = os.LookupEnv("POSTGRES_TIMEZONE"); !ok {
		missingConfigs = append(missingConfigs, "POSTGRES_TIMEZONE")
	}

	return &config, err
}

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

	if config.Server.Port, ok = os.LookupEnv("SERVER_PORT"); !ok {
		missingConfigs = append(missingConfigs, "SERVER_PORT")
	}

	if config.Postgres.Port, ok = os.LookupEnv("POSTGRES_PORT"); !ok {
		missingConfigs = append(missingConfigs, "POSTGRES_PORT")
	}

	if config.Postgres.Host, ok = os.LookupEnv("POSTGRES_HOST"); !ok {
		missingConfigs = append(missingConfigs, "POSTGRES_HOST")
	}

	if config.Postgres.User, ok = os.LookupEnv("POSTGRES_USER"); !ok {
		missingConfigs = append(missingConfigs, "POSTGRES_USER")
	}

	if config.Postgres.Password, ok = os.LookupEnv("POSTGRES_PASSWORD"); !ok {
		missingConfigs = append(missingConfigs, "POSTGRES_PASSWORD")
	}

	if config.Postgres.Database, ok = os.LookupEnv("POSTGRES_DATABASE"); !ok {
		missingConfigs = append(missingConfigs, "POSTGRES_DATABASE")
	}

	if config.Postgres.TimeZone, ok = os.LookupEnv("POSTGRES_TIMEZONE"); !ok {
		missingConfigs = append(missingConfigs, "POSTGRES_TIMEZONE")
	}

	return &config, err
}

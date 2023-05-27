package config

import (
	"fmt"
	"os"
	"strconv"
)

type DB struct {
	Host string
	Port int
}

type Server struct {
	Port string
}
type Config struct {
	DB     DB
	Server Server
}

func LoadFromEnv() (*Config, error) {
	var config Config

	if dbHost, ok := os.LookupEnv("DB_HOST"); ok {
		config.DB.Host = dbHost
	} else {
		return nil, fmt.Errorf("missing DB_HOST environment variable")
	}

	if dbPortStr, ok := os.LookupEnv("DB_PORT"); ok {
		dbPort, err := strconv.Atoi(dbPortStr)
		if err != nil {
			return nil, fmt.Errorf("failed to parse DB_PORT: %w", err)
		}
		config.DB.Port = dbPort
	} else {
		return nil, fmt.Errorf("missing DB_PORT environment variable")
	}

	if serverPort, ok := os.LookupEnv("PORT"); ok {
		config.Server.Port = serverPort
	} else {
		return nil, fmt.Errorf("missing PORT environmental variable")
	}

	return &config, nil
}

package config

import (
	"fmt"
	"os"
	"strconv"
)

type ExchangeApi struct {
	Name   string
	APIKey string
}

type DB struct {
	Host string
	Port int
}

type Config struct {
	API ExchangeApi
	DB  DB
}

func LoadFromEnv() (*Config, error) {
	var config Config

	if apiName, ok := os.LookupEnv("EXCHANGE_API_NAME"); ok {
		config.API.Name = apiName
	} else {
		return nil, fmt.Errorf("missing EXCHANGE_API_NAME environment variable")
	}

	if apiKey, ok := os.LookupEnv("EXCHANGE_API_KEY"); ok {
		config.API.APIKey = apiKey
	} else {
		return nil, fmt.Errorf("missing EXCHANGE_API_KEY environment variable")
	}

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

	return &config, nil
}

package config

import (
	"fmt"
	"os"
)

type Jwt struct {
	Secret        string
	ExpiryMinutes int
}

type DatabaseService struct {
	Port string
	Host string
}

type OtpService struct {
	Port string
	Host string
}

type Redis struct {
	Port     string
	Host     string
	User     string
	Password string
}

type Rabbit struct {
	Port     string
	Host     string
	User     string
	Password string
}

type Config struct {
	Rabbit   Rabbit
	Jwt      Jwt
	Database DatabaseService
	OTP      OtpService
	Redis    Redis
}

func LoadConfig() (*Config, error) {
	rabbitPort, ok := os.LookupEnv("RABBIT_PORT")
	if !ok {
		return nil, fmt.Errorf("missing required environment variable RABBIT_PORT")
	}

	rabbitHost, ok := os.LookupEnv("RABBIT_HOST")
	if !ok {
		return nil, fmt.Errorf("missing required environment variable RABBIT_HOST")
	}

	rabbitUser, ok := os.LookupEnv("RABBIT_USER")
	if !ok {
		return nil, fmt.Errorf("missing required environment variable RABBIT_USER")
	}

	rabbitPassword, ok := os.LookupEnv("RABBIT_PASSWORD")
	if !ok {
		return nil, fmt.Errorf("missing required environment variable RABBIT_PASSWORD")
	}
	jwtSecret, ok := os.LookupEnv("JWT_SECRET")
	if !ok {
		return nil, fmt.Errorf("missing required environment variable JWT_SECRET")
	}

	dbPort, ok := os.LookupEnv("DB_PORT")
	if !ok {
		return nil, fmt.Errorf("missing required environment variable DB_PORT")
	}

	dbHost, ok := os.LookupEnv("DB_HOST")
	if !ok {
		return nil, fmt.Errorf("missing required environment variable DB_HOST")
	}

	otpPort, ok := os.LookupEnv("OTP_PORT")
	if !ok {
		return nil, fmt.Errorf("missing required environment variable OTP_PORT")
	}

	otpHost, ok := os.LookupEnv("OTP_HOST")
	if !ok {
		return nil, fmt.Errorf("missing required environment variable OTP_HOST")
	}

	redisHost, ok := os.LookupEnv("REDIS_HOST")
	if !ok {
		return nil, fmt.Errorf("missing required environment variable REDIS_HOST")
	}

	redisPort, ok := os.LookupEnv("REDIS_PORT")
	if !ok {
		return nil, fmt.Errorf("missing required environment variable OTP_HOST")
	}

	redisUser, ok := os.LookupEnv("REDIS_USER")
	if !ok {
		return nil, fmt.Errorf("missing required environment variable REDIS_USER")
	}

	redisPassword, ok := os.LookupEnv("REDIS_PASSWORD")
	if !ok {
		return nil, fmt.Errorf("missing required environment variable REDIS_PASSWORD")
	}

	cfg := &Config{
		Rabbit: Rabbit{
			Port:     rabbitPort,
			Host:     rabbitHost,
			User:     rabbitUser,
			Password: rabbitPassword,
		},
		Jwt: Jwt{
			Secret:        jwtSecret,
			ExpiryMinutes: 60,
		},
		Database: DatabaseService{
			Port: dbPort,
			Host: dbHost,
		},
		OTP: OtpService{
			Port: otpPort,
			Host: otpHost,
		},
		Redis: Redis{
			Port:     redisPort,
			Host:     redisHost,
			User:     redisUser,
			Password: redisPassword,
		},
	}

	return cfg, nil
}

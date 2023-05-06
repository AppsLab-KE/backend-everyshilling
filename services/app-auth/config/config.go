package config

import (
	"fmt"
	"os"
	"strconv"
)

type Jwt struct {
	Secret            string
	ExpiryMinutes     int
	RefreshExpiryDays int
}

type Consul struct {
	Port string
	Host string
}

func (c *Consul) Address() string {
	return fmt.Sprintf("%s:%s", c.Host, c.Port)
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

type Config struct {
	Jwt      Jwt
	Database DatabaseService
	OTP      OtpService
	Redis    Redis
	Consul   Consul
}

func LoadConfig() (*Config, error) {
	jwtSecret, ok := os.LookupEnv("JWT_SECRET")
	if !ok {
		return nil, fmt.Errorf("missing required environment variable JWT_SECRET")
	}

	jwtExpiry, ok := os.LookupEnv("JWT_EXPIRY")
	if !ok {
		return nil, fmt.Errorf("missing required environment variable JWT_EXPIRY")
	}

	// convert string to int
	jwtExpiryInt, err := strconv.Atoi(jwtExpiry)
	if err != nil {
		return nil, fmt.Errorf("invalid JWT_EXPIRY")
	}

	jwtRefreshExpiry, ok := os.LookupEnv("JWT_REFRESH_EXPIRY")
	if !ok {
		return nil, fmt.Errorf("missing required environment variable JWT_REFRESH_EXPIRY")
	}

	// convert string to int
	jwtRefreshExpiryInt, err := strconv.Atoi(jwtRefreshExpiry)
	if err != nil {
		return nil, fmt.Errorf("invalid JWT_REFRESH_EXPIRY")
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

	consulPort, ok := os.LookupEnv("CONSUL_PORT")
	if !ok {
		return nil, fmt.Errorf("missing required environment variable CONSUL_PORT")
	}

	consulHost, ok := os.LookupEnv("CONSUL_HOST")
	if !ok {
		return nil, fmt.Errorf("missing required environment variable CONSUL_HOST")
	}

	cfg := &Config{
		Consul: Consul{
			Port: consulPort,
			Host: consulHost,
		},
		Jwt: Jwt{
			Secret:            jwtSecret,
			ExpiryMinutes:     jwtExpiryInt,
			RefreshExpiryDays: jwtRefreshExpiryInt,
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

package config

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

type Cache struct {
	Port string
	Host string
}

type Rabbit struct {
	Port     string
	Host     string
	User     string
	Password string
}

type Config struct {
	Rabbit   Rabbit
	Cache    Cache
	Jwt      Jwt
	Database DatabaseService
	OTP      OtpService
}

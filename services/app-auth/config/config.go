package config

type Jwt struct {
	Secret        string
	ExpiryMinutes int
}

type Database struct {
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
	Database Database
}

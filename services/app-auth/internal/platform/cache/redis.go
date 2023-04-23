package cache

import (
	"github.com/AppsLab-KE/backend-everyshilling/services/app-authentication/config"
	"github.com/redis/go-redis/v9"
)

func NewClient(cfg config.Redis) *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr:     cfg.Host + ":" + cfg.Port,
		Password: "",
		DB:       0,
	})

	return rdb
}

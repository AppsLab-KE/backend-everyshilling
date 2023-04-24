package cache

import (
	"context"
	"github.com/AppsLab-KE/backend-everyshilling/services/app-authentication/config"
	"github.com/redis/go-redis/v9"
	log "github.com/sirupsen/logrus"
)

func NewClient(cfg config.Redis) (*redis.Client, error) {
	rdb := redis.NewClient(&redis.Options{
		Addr:     cfg.Host + ":" + cfg.Port,
		Password: "",
		DB:       0,
	})

	if status := rdb.Ping(context.Background()); status.Err() != nil {
		log.Panicf("error connecting to redis: %v", status.Err())
		return nil, status.Err()
	} else {
		log.Infof("connected to redis and successfuly pinged")
	}

	return rdb, nil
}

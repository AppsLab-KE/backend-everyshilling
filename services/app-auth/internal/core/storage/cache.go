package storage

import (
	"context"
	"github.com/AppsLab-KE/backend-everyshilling/services/app-authentication/config"
	"github.com/AppsLab-KE/backend-everyshilling/services/app-authentication/internal/core/adapters"
	"github.com/AppsLab-KE/backend-everyshilling/services/app-authentication/internal/platform/cache"
	"github.com/redis/go-redis/v9"
)

type Cache struct {
	client *redis.Client
}

func (c Cache) SavePhoneFromLoginOTP(ctx context.Context, trackerUUID, phone string) error {
	err := c.client.Set(ctx, trackerUUID, phone, 0).Err()
	return err
}

func (c Cache) GetPhoneFromLoginOTP(ctx context.Context, trackerUUID string) (string, error) {
	val, err := c.client.Get(ctx, trackerUUID).Result()
	return val, err
}

func (c Cache) SavePhoneFromResetOTP(ctx context.Context, trackerUUID, phone string) error {
	err := c.client.Set(ctx, trackerUUID, phone, 0).Err()
	return err
}

func (c Cache) GetPhoneFromResetOTP(ctx context.Context, trackerUUID string) (string, error) {
	val, err := c.client.Get(ctx, trackerUUID).Result()
	return val, err
}

func (c Cache) SavePhoneFromVerificationOTP(ctx context.Context, trackerUUID, phone string) error {
	err := c.client.Set(ctx, trackerUUID, phone, 0).Err()
	return err
}

func (c Cache) GetPhoneFromVerificationOTP(ctx context.Context, trackerUUID string) (string, error) {
	val, err := c.client.Get(ctx, trackerUUID).Result()
	return val, err
}

func NewCacheStorage(redisCfg config.Redis) (adapters.CacheStorage, error) {
	client := cache.NewClient(redisCfg)
	return &Cache{
		client: client,
	}, nil
}

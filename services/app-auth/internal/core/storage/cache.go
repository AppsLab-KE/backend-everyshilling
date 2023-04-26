package storage

import (
	"context"
	"github.com/AppsLab-KE/backend-everyshilling/services/app-authentication/config"
	"github.com/AppsLab-KE/backend-everyshilling/services/app-authentication/internal/core/adapters"
	"github.com/AppsLab-KE/backend-everyshilling/services/app-authentication/internal/platform/cache"
	"github.com/redis/go-redis/v9"
	"time"
)

const (
	Expiration = time.Minute * 10
)

type Cache struct {
	client *redis.Client
}

func (c Cache) BlacklistToken(ctx context.Context, userUUID string) error {
	err := c.client.Set(ctx, userUUID, "true", Expiration*10000).Err()
	return err
}

func (c Cache) IsTokenBlacklisted(ctx context.Context, userUUID string) (bool, error) {
	val, err := c.client.Get(ctx, userUUID).Result()
	if err != nil {
		if err == redis.Nil {
			return false, nil
		}
		return false, err
	}
	if val == "true" {
		return true, nil
	}
	return false, nil
}

func (c Cache) UnBlacklistToken(ctx context.Context, userUUID string) error {
	err := c.client.Del(ctx, userUUID).Err()
	return err
}

func (c Cache) InvalidateLoginTracker(ctx context.Context, trackerUUID string) error {
	err := c.client.Del(ctx, trackerUUID).Err()
	return err
}

func (c Cache) InvalidateResetTracker(ctx context.Context, trackerUID string) error {
	err := c.client.Del(ctx, trackerUID).Err()
	return err
}

func (c Cache) InvalidateVerificationTracker(ctx context.Context, trackerUUID string) error {
	err := c.client.Del(ctx, trackerUUID).Err()
	return err
}

func (c Cache) SavePhoneFromLoginOTP(ctx context.Context, trackerUUID, phone string) error {
	err := c.client.Set(ctx, trackerUUID, phone, Expiration).Err()
	return err
}

func (c Cache) GetPhoneFromLoginOTP(ctx context.Context, trackerUUID string) (string, error) {
	val, err := c.client.Get(ctx, trackerUUID).Result()
	return val, err
}

func (c Cache) SavePhoneFromResetOTP(ctx context.Context, trackerUUID, phone string) error {
	err := c.client.Set(ctx, trackerUUID, phone, Expiration).Err()
	return err
}

func (c Cache) GetPhoneFromResetOTP(ctx context.Context, trackerUUID string) (string, error) {
	val, err := c.client.Get(ctx, trackerUUID).Result()
	return val, err
}

func (c Cache) SavePhoneFromVerificationOTP(ctx context.Context, trackerUUID, phone string) error {
	err := c.client.Set(ctx, trackerUUID, phone, Expiration).Err()
	return err
}

func (c Cache) GetPhoneFromVerificationOTP(ctx context.Context, trackerUUID string) (string, error) {
	val, err := c.client.Get(ctx, trackerUUID).Result()
	return val, err
}

func NewCacheStorage(redisCfg config.Redis) (adapters.CacheStorage, error) {
	client, err := cache.NewClient(redisCfg)
	if err != nil {
		return nil, err
	}
	return &Cache{
		client: client,
	}, nil
}

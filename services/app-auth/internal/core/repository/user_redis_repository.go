package repository

import (
	"context"
	"github.com/AppsLab-KE/backend-everyshilling/services/app-authentication/internal/core/entity"
)

type UserRedisRepo interface {
	Set(ctx context.Context, user entity.User) error
	Get(ctx context.Context, uuid string) (*entity.User, error)
}

package repository

import (
	"context"
	"github.com/AppsLab-KE/backend-everyshilling/services/app-authentication/internal/core/entity"
)

type UserRepository interface {
	Create(ctx context.Context, user entity.User) error
	Update(ctx context.Context, user entity.User) error
	FindByID(ctx context.Context, uuid string) (*entity.User, error)
	FindByEmail(ctx context.Context, emailAddress string) (*entity.User, error)
}

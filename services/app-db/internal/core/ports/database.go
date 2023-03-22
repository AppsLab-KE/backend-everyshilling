package ports

import (
	"appslab.co.ke/everyshilling/app-db/internal/core/models"
	"context"
)

type Storage interface {
	SaveUser(ctx context.Context, user models.User) error
}

type UserService interface {
	CreateUser(ctx context.Context)
}

type UserRepo interface {
	SaveUser(ctx context.Context) error
}

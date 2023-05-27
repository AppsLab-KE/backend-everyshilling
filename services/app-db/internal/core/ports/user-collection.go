package ports

import (
	"context"
	"github.com/AppsLab-KE/backend-everyshilling/services/app-db/internal/core/models"
)

type UserStorage interface {
	CreateUser(context.Context, *models.User) (*models.User, error)
	UpdateUser(context.Context, *models.User) (*models.User, error)
	GetPagedUsers(context.Context, *models.Paging) ([]models.User, error)
	GetUserByField(context.Context, *map[string]interface{}, models.Paging) ([]models.User, error)
}

type UserCache interface {
}

type UserService interface {
	CreateUser(ctx context.Context)
}

type UserRepo interface {
	CreateUser(context.Context, *models.User) (*models.User, error)
	UpdateUser(context.Context, *models.User) (*models.User, error)
	GetPagedUsers(context.Context, *models.Paging) ([]models.User, error)
	GetUserByField(context.Context, *map[string]interface{}, models.Paging) ([]models.User, error)
}

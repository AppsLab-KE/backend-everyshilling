package ports

import (
	"context"
	"github.com/AppsLab-KE/backend-everyshilling/services/app-db/internal/core/models"
	"github.com/AppsLab-KE/backend-everyshilling/services/app-db/internal/dto"
)

type UserStorage interface {
	CreateUser(context.Context, *models.User) (*models.User, error)
	UpdateUser(context.Context, *models.User) (*models.User, error)
	GetPagedUsers(context.Context, *dto.Paging) ([]models.User, error)
	GetUserByField(context.Context, *map[string]interface{}, dto.Paging) ([]models.User, error)
}

type UserCache interface {
}

type UserService interface {
	CreateUser(ctx context.Context)
}

type UserRepo interface {
	CreateUser(context.Context, *models.User) (*models.User, error)
	UpdateUser(context.Context, *models.User) (*models.User, error)
	GetPagedUsers(context.Context, *dto.Paging) ([]models.User, error)
	GetUserByField(context.Context, *map[string]interface{}, dto.Paging) ([]models.User, error)
}

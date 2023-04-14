package repository

import (
	"context"
	"github.com/AppsLab-KE/backend-everyshilling/services/app-db/internal/core/models"
	"github.com/AppsLab-KE/backend-everyshilling/services/app-db/internal/core/ports"
	"github.com/AppsLab-KE/backend-everyshilling/services/app-db/internal/dto"
)

type userRepo struct {
	db    ports.UserStorage
	cache ports.UserCache
}

func (u userRepo) CreateUser(ctx context.Context, user *models.User) (*models.User, error) {
	return u.db.CreateUser(ctx, user)
}

func (u userRepo) UpdateUser(ctx context.Context, user *models.User) (*models.User, error) {
	return u.db.UpdateUser(ctx, user)
}

func (u userRepo) GetPagedUsers(ctx context.Context, paging *dto.Paging) ([]models.User, error) {
	return u.db.GetPagedUsers(ctx, paging)
}

func (u userRepo) GetUserByField(ctx context.Context, m *map[string]interface{}, paging dto.Paging) ([]models.User, error) {
	return u.db.GetUserByField(ctx, m, paging)
}

func NewUserRepo(db ports.UserStorage, cache ports.UserCache) ports.UserRepo {
	return &userRepo{
		db:    db,
		cache: cache,
	}
}

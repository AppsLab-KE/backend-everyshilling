package storage

import (
	"context"
	"fmt"
	"github.com/AppsLab-KE/backend-everyshilling/services/app-db/internal/core/models"
	"github.com/AppsLab-KE/backend-everyshilling/services/app-db/internal/core/ports"
	"github.com/AppsLab-KE/backend-everyshilling/services/app-db/internal/dto"
	"gorm.io/gorm"
)

type userStorage struct {
	client *gorm.DB
}

func (s userStorage) CreateUser(ctx context.Context, user *models.User) (*models.User, error) {
	// Check if user already exists
	var existingUser models.User
	if err := s.client.Where("email = ? OR phone = ?", user.Email, user.Phone).First(&existingUser).Error; err == nil {
		return nil, fmt.Errorf("user with email %s or phone %s already exists", user.Email, user.Phone)
	}

	// Create user
	if err := s.client.Create(user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func (s userStorage) UpdateUser(ctx context.Context, user *models.User) (*models.User, error) {
	// Check if user exists
	var existingUser models.User
	if err := s.client.Where("id = ?", user.ID).First(&existingUser).Error; err != nil {
		return nil, fmt.Errorf("user with ID %d not found", user.ID)
	}

	// Update user
	if err := s.client.Model(&existingUser).Updates(user).Error; err != nil {
		return nil, err
	}
	return user, nil
}
func (s userStorage) GetPagedUsers(ctx context.Context, paging *dto.Paging) ([]models.User, error) {
	var users []models.User
	if err := s.client.Offset(paging.Offset).Limit(paging.Limit).Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

func (s userStorage) GetUserByField(ctx context.Context, m *map[string]interface{}, paging dto.Paging) ([]models.User, error) {
	var users []models.User
	query := s.client.Offset(paging.Offset).Limit(paging.Limit)
	for k, v := range *m {
		query = query.Where(k+" = ?", v)
	}
	if err := query.Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

func NewUserStorage(client *gorm.DB) ports.UserStorage {
	return &userStorage{
		client: client,
	}
}

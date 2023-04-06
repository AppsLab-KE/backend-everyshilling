package repository

import (
	"github.com/AppsLab-KE/backend-everyshilling/services/app-authentication/internal/core/adapters"
	"github.com/AppsLab-KE/backend-everyshilling/services/app-authentication/internal/core/entity"
	"github.com/AppsLab-KE/backend-everyshilling/services/app-authentication/internal/core/storage"
	"github.com/AppsLab-KE/backend-everyshilling/services/app-authentication/internal/dto"
)

type authRepo struct {
	client storage.Db
}

func (a authRepo) CreateOtpCode(data entity.Otp) error {
	//TODO implement me
	panic("implement me")
}

func (a authRepo) CreateUser(user dto.RegisterRequest) (*entity.User, error) {
	//TODO implement me
	panic("implement me")
}

func (a authRepo) GetUserByPhone(phone string) (*entity.User, error) {
	//TODO implement me
	panic("implement me")
}

func (a authRepo) GetUserByEmail(email string) (*entity.User, error) {
	//TODO implement me
	panic("implement me")
}

func (a authRepo) UpdateUser(user entity.User) (*entity.User, error) {
	//TODO implement me
	panic("implement me")
}

func NewDefaultAuthAppDbRepo(db storage.Db) adapters.AuthRepo {
	return &authRepo{
		client: db,
	}
}

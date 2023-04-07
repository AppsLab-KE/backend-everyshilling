package repository

import (
	"context"
	"github.com/AppsLab-KE/backend-everyshilling/services/app-authentication/internal/core/adapters"
	"github.com/AppsLab-KE/backend-everyshilling/services/app-authentication/internal/core/entity"
	"github.com/AppsLab-KE/backend-everyshilling/services/app-authentication/internal/core/storage"
	"github.com/AppsLab-KE/backend-everyshilling/services/app-authentication/internal/dto"
	pb "github.com/AppsLab-KE/be-go-gen-grpc/db"
)

type authRepo struct {
	client storage.Db
}

func (a authRepo) CreateOtpCode(ctx context.Context, data entity.Otp) error {
	//TODO implement me
	panic("implement me")
}

func (a authRepo) CreateUser(ctx context.Context, registerRequest dto.RegisterRequest) (*entity.User, error) {
	userReq := pb.CreateUserReq
	userRes, err := a.client.CreateUser(ctx, userReq)
	if err != nil {
		return nil, err
	}

	_ = userRes
	user := entity.User{}
	return &user, nil
}

func (a authRepo) GetUserByPhone(ctx context.Context, phone string) (*entity.User, error) {
	keyValueReq := &pb.KeyValueRequest{}
	userRes, err := a.client.GetUserByField(ctx, keyValueReq)
	if err != nil {
		return nil, err
	}

	_ = userRes
	user := &entity.User{}

	return user, err
}

func (a authRepo) GetUserByEmail(ctx context.Context, email string) (*entity.User, error) {
	keyValueReq := &pb.KeyValueRequest{}
	userRes, err := a.client.GetUserByField(ctx, keyValueReq)
	if err != nil {
		return nil, err
	}
	_ = userRes
	user := &entity.User{}

	return user, err
}

func (a authRepo) UpdateUser(ctx context.Context, user entity.User) (*entity.User, error) {
	//TODO implement me
	panic("implement me")
}

func NewDefaultAuthAppDbRepo(db storage.Db) adapters.AuthRepo {
	return &authRepo{
		client: db,
	}
}

package grpc

import (
	"context"
	"github.com/AppsLab-KE/backend-everyshilling/services/app-authentication/internal/core/entity"
	"github.com/AppsLab-KE/backend-everyshilling/services/app-authentication/internal/core/repository"
	"github.com/AppsLab-KE/backend-everyshilling/services/app-authentication/internal/platform/database/grpc"
)

type userGrpcRepo struct {
	client *grpc.Client
}

func (u userGrpcRepo) Create(ctx context.Context, user entity.User) error {
	//TODO implement me
	panic("implement me")
}

func (u userGrpcRepo) Update(ctx context.Context, user entity.User) error {
	//TODO implement me
	panic("implement me")
}

func (u userGrpcRepo) FindByID(ctx context.Context, uuid string) (*entity.User, error) {
	//TODO implement me
	panic("implement me")
}

func (u userGrpcRepo) FindByEmail(ctx context.Context, emailAddress string) (*entity.User, error) {
	//TODO implement me
	panic("implement me")
}

func NewUserGRPCRepo(client *grpc.Client) repository.UserGRPCRepo {
	return &userGrpcRepo{
		client: client,
	}
}

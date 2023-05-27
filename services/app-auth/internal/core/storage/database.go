package storage

import (
	"context"
	"errors"
	"github.com/AppsLab-KE/backend-everyshilling/services/app-authentication/config"
	"github.com/AppsLab-KE/backend-everyshilling/services/app-authentication/internal/core/adapters"
	"github.com/AppsLab-KE/backend-everyshilling/services/app-authentication/internal/platform/apps"
	"github.com/AppsLab-KE/be-go-gen-grpc/db"
	"google.golang.org/grpc"
)

var (
	ErrDbDown = errors.New("database service down")
)

type dbStorage struct {
	dbClient db.DbServiceClient
}

func (d dbStorage) UpdateUser(ctx context.Context, req *db.UpdateUserReq, opts ...grpc.CallOption) (*db.UpdateUserRes, error) {
	if d.dbClient == nil {
		return nil, ErrDbDown
	}
	return d.dbClient.UpdateUser(ctx, req)
}

func (d dbStorage) HealthCheck(ctx context.Context, req *db.DefaultRequest, opts ...grpc.CallOption) (*db.HealthResponse, error) {
	if d.dbClient == nil {
		return nil, ErrDbDown
	}
	return d.dbClient.HealthCheck(ctx, req, opts...)
}

func (d dbStorage) CreateUser(ctx context.Context, req *db.CreateUserReq, opts ...grpc.CallOption) (*db.CreateUserRes, error) {
	if d.dbClient == nil {
		return nil, ErrDbDown
	}
	return d.dbClient.CreateUser(ctx, req, opts...)
}

func (d dbStorage) GetUserByField(ctx context.Context, req *db.GetByfieldReq, opts ...grpc.CallOption) (*db.GetByfieldRes, error) {
	if d.dbClient == nil {
		return nil, ErrDbDown
	}
	return d.dbClient.GetUserByField(ctx, req, opts...)
}

func (d dbStorage) GetPagedUsers(ctx context.Context, req *db.GetPagedUsersReq, opts ...grpc.CallOption) (*db.GetPagedUsersRes, error) {
	if d.dbClient == nil {
		return nil, ErrDbDown
	}
	return d.dbClient.GetPagedUsers(ctx, req, opts...)
}

func NewDbStorage(serviceCfg config.DatabaseService) (adapters.DBStorage, error) {
	client, err := apps.NewDBServiceClient(serviceCfg)
	if err != nil {
		return nil, err
	}
	return &dbStorage{
		dbClient: client,
	}, nil
}

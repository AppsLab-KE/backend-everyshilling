package storage

import (
	"context"
	"errors"
	"github.com/AppsLab-KE/backend-everyshilling/services/app-ratescron/config"
	"github.com/AppsLab-KE/backend-everyshilling/services/app-ratescron/internal/core/ports"
	"github.com/AppsLab-KE/backend-everyshilling/services/app-ratescron/internal/platform"
	"github.com/AppsLab-KE/be-go-gen-grpc/db"
	"google.golang.org/grpc"
)

var (
	ErrDbDown = errors.New("database service down")
)

type dbStorage struct {
	dbClient db.DbServiceClient
}

func (d dbStorage) CreateConversionRate(ctx context.Context, in *db.CreateConversionRateRequest, opts ...grpc.CallOption) (*db.CreateConversionRateResponse, error) {
	return d.dbClient.CreateConversionRate(ctx, in, opts...)
}

func (d dbStorage) ReadConversionRate(ctx context.Context, in *db.ReadConversionRateRequest, opts ...grpc.CallOption) (*db.ReadConversionRateResponse, error) {
	return d.dbClient.ReadConversionRate(ctx, in, opts...)
}

func (d dbStorage) UpdateConversionRate(ctx context.Context, in *db.UpdateConversionRateRequest, opts ...grpc.CallOption) (*db.UpdateConversionRateResponse, error) {
	return d.dbClient.UpdateConversionRate(ctx, in, opts...)
}

func (d dbStorage) DeleteConversionRate(ctx context.Context, in *db.DeleteConversionRateRequest, opts ...grpc.CallOption) (*db.DeleteConversionRateResponse, error) {
	return d.dbClient.DeleteConversionRate(ctx, in, opts...)
}

// Not implemented
func (d dbStorage) UpdateUser(ctx context.Context, req *db.UpdateUserReq, opts ...grpc.CallOption) (*db.UpdateUserRes, error) {
	return nil, errors.New("not implemented")
}

func (d dbStorage) HealthCheck(ctx context.Context, req *db.DefaultRequest, opts ...grpc.CallOption) (*db.HealthResponse, error) {
	return nil, errors.New("not implemented")

}

func (d dbStorage) CreateUser(ctx context.Context, req *db.CreateUserReq, opts ...grpc.CallOption) (*db.CreateUserRes, error) {
	return nil, errors.New("not implemented")

}

func (d dbStorage) GetUserByField(ctx context.Context, req *db.GetByfieldReq, opts ...grpc.CallOption) (*db.GetByfieldRes, error) {
	return nil, errors.New("not implemented")

}

func (d dbStorage) GetPagedUsers(ctx context.Context, req *db.GetPagedUsersReq, opts ...grpc.CallOption) (*db.GetPagedUsersRes, error) {
	return nil, errors.New("not implemented")

}

func NewDbStorage(serviceCfg config.DB) (ports.DBStorage, error) {
	client, err := platform.NewDBServiceClient(serviceCfg)
	if err != nil {
		return nil, err
	}
	return &dbStorage{
		dbClient: client,
	}, nil
}

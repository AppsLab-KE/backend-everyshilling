package ports

import (
	"context"
	"github.com/AppsLab-KE/be-go-gen-grpc/db"
	"google.golang.org/grpc"
)

type DBStorage interface {
	HealthCheck(ctx context.Context, in *db.DefaultRequest, opts ...grpc.CallOption) (*db.HealthResponse, error)
	// USERS
	CreateUser(ctx context.Context, in *db.CreateUserReq, opts ...grpc.CallOption) (*db.CreateUserRes, error)
	UpdateUser(ctx context.Context, in *db.UpdateUserReq, opts ...grpc.CallOption) (*db.UpdateUserRes, error)
	GetPagedUsers(ctx context.Context, in *db.GetPagedUsersReq, opts ...grpc.CallOption) (*db.GetPagedUsersRes, error)
	GetUserByField(ctx context.Context, in *db.GetByfieldReq, opts ...grpc.CallOption) (*db.GetByfieldRes, error)
	CreateConversionRate(ctx context.Context, in *db.CreateConversionRateRequest, opts ...grpc.CallOption) (*db.CreateConversionRateResponse, error)
	ReadConversionRate(ctx context.Context, in *db.ReadConversionRateRequest, opts ...grpc.CallOption) (*db.ReadConversionRateResponse, error)
	UpdateConversionRate(ctx context.Context, in *db.UpdateConversionRateRequest, opts ...grpc.CallOption) (*db.UpdateConversionRateResponse, error)
	DeleteConversionRate(ctx context.Context, in *db.DeleteConversionRateRequest, opts ...grpc.CallOption) (*db.DeleteConversionRateResponse, error)
}

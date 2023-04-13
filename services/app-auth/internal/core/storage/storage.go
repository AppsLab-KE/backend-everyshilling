package storage

import (
	"context"
	dbproto "github.com/AppsLab-KE/be-go-gen-grpc/db"
)

type Db interface {
	HealthCheck(context.Context, *dbproto.DefaultRequest) (*dbproto.HealthResponse, error)
	CreateUser(context.Context, *dbproto.CreateUserReq) (*dbproto.CreateUserRes, error)
	GetUserByField(context.Context, *dbproto.GetByfieldReq) (*dbproto.GetByfieldRes, error)
	GetPagedUsers(context.Context, *dbproto.GetPagedUsersReq) (*dbproto.GetPagedUsersRes, error)
}

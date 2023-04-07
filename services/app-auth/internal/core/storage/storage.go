package storage

import (
	"context"
	pb "github.com/AppsLab-KE/be-go-gen-grpc/db"
)

type Db interface {
	HealthCheck(context.Context, *pb.DefaultRequest) (*pb.HealthResponse, error)
	CreateUser(context.Context, *pb.CreateUserReq) (*pb.User, error)
	GetUserByField(context.Context, *pb.KeyValueRequest) (*pb.User, error)
}

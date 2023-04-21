package adapters

import (
	"context"
	dbproto "github.com/AppsLab-KE/be-go-gen-grpc/db"
	otpProto "github.com/AppsLab-KE/be-go-gen-grpc/otp"
	"google.golang.org/grpc"
)

type DBStorage interface {
	HealthCheck(ctx context.Context, req *dbproto.DefaultRequest, opts ...grpc.CallOption) (*dbproto.HealthResponse, error)
	CreateUser(ctx context.Context, req *dbproto.CreateUserReq, opts ...grpc.CallOption) (*dbproto.CreateUserRes, error)
	UpdateUser(ctx context.Context, req *dbproto.UpdateUserReq, opts ...grpc.CallOption) (*dbproto.UpdateUserRes, error)
	GetUserByField(ctx context.Context, req *dbproto.GetByfieldReq, opts ...grpc.CallOption) (*dbproto.GetByfieldRes, error)
	GetPagedUsers(ctx context.Context, req *dbproto.GetPagedUsersReq, opts ...grpc.CallOption) (*dbproto.GetPagedUsersRes, error)
}

type OTPStorage interface {
	HealthCheck(ctx context.Context, in *otpProto.DefaultRequest, opts ...grpc.CallOption) (*otpProto.HealthResponse, error)
	CreateAndSendOtp(ctx context.Context, in *otpProto.CreateAndSendOtpReq, opts ...grpc.CallOption) (*otpProto.CreateAndSendOtpRes, error)
	ResendOTP(ctx context.Context, in *otpProto.ResendOTPRed, opts ...grpc.CallOption) (*otpProto.ResendOTPRes, error)
	VerifyOtp(ctx context.Context, in *otpProto.VerifyOTPReq, opts ...grpc.CallOption) (*otpProto.VerifyOTPRes, error)
}

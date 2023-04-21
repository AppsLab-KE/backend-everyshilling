package storage

import (
	"context"
	"errors"
	"github.com/AppsLab-KE/backend-everyshilling/services/app-authentication/config"
	"github.com/AppsLab-KE/backend-everyshilling/services/app-authentication/internal/core/adapters"
	"github.com/AppsLab-KE/backend-everyshilling/services/app-authentication/internal/platform/apps"
	"github.com/AppsLab-KE/be-go-gen-grpc/otp"
	"google.golang.org/grpc"
)

var (
	ErrOTPDown = errors.New("OTP service down")
)

type otpStorage struct {
	client otp.OtpServiceClient
}

func (o otpStorage) ResendOTP(ctx context.Context, in *otp.ResendOTPRed, opts ...grpc.CallOption) (*otp.ResendOTPRes, error) {
	if o.client == nil {
		return nil, ErrOTPDown
	}
	return o.client.ResendOTP(ctx, in, opts...)
}

func (o otpStorage) HealthCheck(ctx context.Context, in *otp.DefaultRequest, opts ...grpc.CallOption) (*otp.HealthResponse, error) {
	if o.client == nil {
		return nil, ErrOTPDown
	}
	return o.client.HealthCheck(ctx, in, opts...)
}

func (o otpStorage) CreateAndSendOtp(ctx context.Context, in *otp.CreateAndSendOtpReq, opts ...grpc.CallOption) (*otp.CreateAndSendOtpRes, error) {
	if o.client == nil {
		return nil, ErrOTPDown
	}
	return o.client.CreateAndSendOtp(ctx, in, opts...)
}

func (o otpStorage) VerifyOtp(ctx context.Context, in *otp.VerifyOTPReq, opts ...grpc.CallOption) (*otp.VerifyOTPRes, error) {
	if o.client == nil {
		return nil, ErrOTPDown
	}
	return o.client.VerifyOtp(ctx, in, opts...)
}

func NewOtpStorage(serviceCfg config.OtpService) (adapters.OTPStorage, error) {
	client, err := apps.NewOTPServiceClient(serviceCfg)
	if err != nil {
		// return nil, err
	}

	return &otpStorage{
		client: client,
	}, nil
}

package adapters

import (
	"context"
	"github.com/AppsLab-KE/backend-everyshilling/services/app-authentication/internal/core/entity"
	"github.com/AppsLab-KE/backend-everyshilling/services/app-authentication/internal/dto"
)

type AuthRepo interface {
	CreateOtpCode(ctx context.Context, data dto.OtpGenReq) (*dto.OtpGenRes, error)
	VerifyOtpCode(ctx context.Context, data dto.OtpVerificationReq) (*dto.OtpVerificationRes, error)
	ResendOtpCode(ctx context.Context, data dto.ResendOTPReq) (*dto.ResendOTPRes, error)
	CreateUser(ctx context.Context, registerRequest dto.RegisterReq) (*entity.User, error)
	GetUserByPhone(ctx context.Context, phone string) (*entity.User, error)
	GetUserByEmail(ctx context.Context, email string) (*entity.User, error)
	UpdateUser(ctx context.Context, user entity.User) (*entity.User, error)

	SavePhoneFromLoginOTP(ctx context.Context, trackerUUID, phone string) error
	InvalidateLoginTracker(ctx context.Context, trackerUUID string) error
	GetPhoneFromLoginOTP(ctx context.Context, trackerUUID string) (string, error)

	SavePhoneFromResetOTP(ctx context.Context, trackerUUID, phone string) error
	InvalidateResetTracker(ctx context.Context, trackerUID string) error
	GetPhoneFromResetOTP(ctx context.Context, trackerUUID string) (string, error)

	SavePhoneFromVerificationOTP(ctx context.Context, trackerUUID, phone string) error
	InvalidateVerificationTracker(ctx context.Context, trackerUUID string) error
	GetPhoneFromVerificationOTP(ctx context.Context, trackerUUID string) (string, error)

	InvalidateSession() error
}

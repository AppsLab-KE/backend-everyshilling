package adapters

import (
	"context"
	"github.com/AppsLab-KE/backend-everyshilling/services/app-authentication/internal/core/entity"
	"github.com/AppsLab-KE/backend-everyshilling/services/app-authentication/internal/dto"
)

type AuthRepo interface {
	CreateOtpCode(ctx context.Context, data dto.OtpGenReq) (*dto.OtpGenRes, error)
	VerifyOtpCode(ctx context.Context, data dto.OtpVerificationReq) (*dto.OtpVerificationRes, error)
	ResendOtpCode(ctx context.Context, data dto.ResendLoginOTPReq) (*dto.ResendLoginOTPRes, error)
	CreateUser(ctx context.Context, registerRequest dto.RegisterRequest) (*entity.User, error)
	GetUserByPhone(ctx context.Context, phone string) (*entity.User, error)
	GetUserByEmail(ctx context.Context, email string) (*entity.User, error)
	UpdateUser(ctx context.Context, user entity.User) (*entity.User, error)
}

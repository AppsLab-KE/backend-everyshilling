package adapters

import (
	"context"
	"github.com/AppsLab-KE/backend-everyshilling/services/app-authentication/internal/core/entity"
	"github.com/AppsLab-KE/backend-everyshilling/services/app-authentication/internal/dto"
)

type AuthService interface {
	RequestOtp(request dto.OtpGenReq) (dto.DefaultRes[*dto.OtpGenRes], error)
	CreateUser(registerRequest dto.RegisterRequest) (*dto.UserRegistrationRes, error)
}

type SessionService interface {
	Invalidate() dto.DefaultRes[interface{}]
}

type AuthRepo interface {
	CreateOtpCode(ctx context.Context, data entity.Otp) error
	CreateUser(ctx context.Context, registerRequest dto.RegisterRequest) (*entity.User, error)
	GetUserByPhone(ctx context.Context, phone string) (*entity.User, error)
	GetUserByEmail(ctx context.Context, email string) (*entity.User, error)
	UpdateUser(ctx context.Context, user entity.User) (*entity.User, error)
}

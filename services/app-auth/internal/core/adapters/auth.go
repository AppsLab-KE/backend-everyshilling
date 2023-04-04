package adapters

import (
	"github.com/AppsLab-KE/backend-everyshilling/services/app-authentication/internal/core/entity"
	"github.com/AppsLab-KE/backend-everyshilling/services/app-authentication/internal/dto"
)

type AuthService interface {
	RequestOtp(request dto.OtpReq) dto.DefaultRes[any]
	CreateUser(registerRequest dto.RegisterRequest) (*dto.UserRegistrationRes, error)
}

type SessionService interface {
	Invalidate() dto.DefaultRes[interface{}]
}

type AuthRepo interface {
	CreateOtpCode(data entity.Otp) error
	CreateUser(registerRequest dto.RegisterRequest) (*entity.User, error)
	GetUserByPhone(phone string) (*entity.User, error)
	GetUserByEmail(email string) (*entity.User, error)
	UpdateUser(user entity.User) (*entity.User, error)
}

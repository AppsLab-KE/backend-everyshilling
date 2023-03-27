package adapters

import (
	"github.com/AppsLab-KE/backend-everyshilling/services/app-authentication/internal/core/entity"
	"github.com/AppsLab-KE/backend-everyshilling/services/app-authentication/internal/dto"
)

type AuthService interface {
	RequestOtp(request dto.OtpReq) dto.DefaultRes
}

type SessionService interface {
	Invalidate() dto.DefaultRes
}

type AuthRepo interface {
	CreateOtpCode(data entity.Otp) error
}

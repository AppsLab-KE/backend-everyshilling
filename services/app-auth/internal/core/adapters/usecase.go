package adapters

import (
	"context"
	"github.com/AppsLab-KE/backend-everyshilling/services/app-authentication/internal/dto"
)

type AuthUseCase interface {
	RegisterUser(ctx context.Context, user dto.RegisterRequest) (*dto.UserRegistrationRes, error)
	SendResetOtp(ctx context.Context, user dto.OtpGenReq) (*dto.OtpGenRes, error)
	LoginUser(ctx context.Context, user dto.LoginInitReq) (*dto.LoginInitRes, error)
	VerifyLoginOTP(ctx context.Context, req dto.OtpVerificationReq) (*dto.OtpVerificationRes, error)
	VerifyResetOTP(ctx context.Context, uuid string, body dto.OtpVerificationReq) (*dto.OtpVerificationRes, error)
	ChangePassword(ctx context.Context, uuid string, body dto.RequestResetCredentials) (*dto.ResetRes, error)
}

package adapters

import (
	"context"
	"github.com/AppsLab-KE/backend-everyshilling/services/app-authentication/internal/dto"
)

type AuthUseCase interface {
	RegisterUser(ctx context.Context, user dto.RegisterReq) (*dto.UserRegistrationRes, error)
	SendResetOtp(ctx context.Context, user dto.OtpGenReq) (*dto.OtpGenRes, error)
	LoginUser(ctx context.Context, user dto.LoginInitReq) (*dto.LoginInitRes, error)
	VerifyLoginOTP(ctx context.Context, req dto.OtpVerificationReq) (*dto.LoginRes, error)
	VerifyResetOTP(ctx context.Context, uuid string, body dto.OtpVerificationReq) (*dto.OtpVerificationRes, error)
	ChangePassword(ctx context.Context, uuid string, body dto.RequestResetCredentials) (*dto.ResetRes, error)

	VerifyPhoneOTP(verificationRequest dto.OtpVerificationReq) (*dto.OtpVerificationRes, error)
	SendVerifyPhoneOTP(request dto.OtpGenReq) (*dto.OtpGenRes, error)

	ResendVerifyPhoneOTP(request dto.ResendOTPReq) (*dto.ResendOTPRes, error)
	ResendLoginOTP(request dto.ResendOTPReq) (*dto.ResendOTPRes, error)
	ResendResetOTP(request dto.ResendOTPReq) (*dto.ResendOTPRes, error)

	RefreshToken(request dto.RefreshTokenReq) (*dto.RefreshTokenRes, error)
	// Logout logs out the user by invalidating the user's session.
	Logout(uuid string) error

	VerifyAccessToken(token string) (string, error)
}

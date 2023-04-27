package usecase

import (
	"context"
	"github.com/AppsLab-KE/backend-everyshilling/services/app-authentication/internal/core/adapters"
	"github.com/AppsLab-KE/backend-everyshilling/services/app-authentication/internal/core/entity"
	"github.com/AppsLab-KE/backend-everyshilling/services/app-authentication/internal/dto"
	"github.com/AppsLab-KE/backend-everyshilling/services/app-authentication/pkg/validation"
)

type AuthUseCase struct {
	authService    adapters.AuthService
	sessionService adapters.SessionService
}

func (a *AuthUseCase) VerifyAccessToken(token string) (string, error) {
	return a.authService.VerifyAccessToken(token)
}

func (a *AuthUseCase) RefreshToken(request dto.RefreshTokenReq) (*dto.RefreshTokenRes, error) {
	res, err := a.authService.RefreshToken(request)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (a *AuthUseCase) Logout(uuid string) error {
	err := a.authService.Logout(uuid)
	if err != nil {
		return err
	}
	return nil
}

func (a *AuthUseCase) LoginUser(ctx context.Context, req dto.LoginInitReq) (*dto.LoginInitRes, error) {
	// validate phone
	if !validation.ValidatePhone(req.PhoneNumber) {
		return nil, entity.NewValidationError("invalid phone number")
	}
	res, err := a.authService.SendLoginOtp(req)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (a *AuthUseCase) SendResetOtp(ctx context.Context, req dto.OtpGenReq) (*dto.OtpGenRes, error) {
	res, err := a.authService.SendResetOTP(req)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (a *AuthUseCase) VerifyLoginOTP(ctx context.Context, req dto.OtpVerificationReq) (*dto.LoginRes, error) {
	res, err := a.authService.VerifyLoginOtp(req)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (a *AuthUseCase) VerifyResetOTP(ctx context.Context, uuid string, body dto.OtpVerificationReq) (*dto.OtpVerificationRes, error) {
	res, err := a.authService.VerifyResetOTP(body)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (a *AuthUseCase) ChangePassword(ctx context.Context, uuid string, body dto.RequestResetCredentials) (*dto.ResetRes, error) {
	resetReq := dto.ResetReq{
		ConfirmPassword: body.ConfirmPassword,
		Password:        body.Password,
		TrackerUUID:     uuid,
	}
	res, err := a.authService.ChangePassword(resetReq)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// RegisterUser Implements authservice to register a new user
func (a *AuthUseCase) RegisterUser(ctx context.Context, user dto.RegisterReq) (*dto.UserRegistrationRes, error) {
	// validate password
	if !validation.ValidatePassword(user.Password) {
		return nil, entity.NewValidationError("password should be at least 8 letters " +
			"and must include combination of small letters, uppercase letters, numbers and symbols.")
	}

	// validate phone
	if !validation.ValidatePhone(user.PhoneNumber) {
		return nil, entity.NewValidationError("phone should be in the format +2547XXXXXXXX")
	}
	res, err := a.authService.CreateUser(user)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (a *AuthUseCase) VerifyPhoneOTP(verificationRequest dto.OtpVerificationReq) (*dto.OtpVerificationRes, error) {
	res, err := a.authService.VerifyPhoneOTP(verificationRequest)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (a *AuthUseCase) SendVerifyPhoneOTP(request dto.AccountVerificationOTPGenReq) (*dto.OtpGenRes, error) {
	// validate phone
	if !validation.ValidatePhone(request.Phone) {
		return nil, entity.NewValidationError("phone number should be in the format +2547XXXXXXXX")
	}
	res, err := a.authService.SendVerifyPhoneOTP(request)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (a *AuthUseCase) ResendVerifyPhoneOTP(request dto.ResendOTPReq) (*dto.ResendOTPRes, error) {
	res, err := a.authService.ResendVerifyPhoneOTP(request)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (a *AuthUseCase) ResendLoginOTP(request dto.ResendOTPReq) (*dto.ResendOTPRes, error) {
	res, err := a.authService.ResendLoginOTP(request)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (a *AuthUseCase) ResendResetOTP(request dto.ResendOTPReq) (*dto.ResendOTPRes, error) {
	res, err := a.authService.ResendResetOTP(request)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func NewAuthUsecase(as adapters.AuthService, ss adapters.SessionService) adapters.AuthUseCase {
	return &AuthUseCase{
		authService:    as,
		sessionService: ss,
	}
}

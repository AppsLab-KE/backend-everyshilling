package usecase

import (
	"context"
	"github.com/AppsLab-KE/backend-everyshilling/services/app-authentication/internal/core/adapters"
	"github.com/AppsLab-KE/backend-everyshilling/services/app-authentication/internal/dto"
)

type AuthUseCase struct {
	authService    adapters.AuthService
	sessionService adapters.SessionService
}

func (a *AuthUseCase) LoginUser(ctx context.Context, req dto.LoginInitReq) (*dto.LoginInitRes, error) {
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

func (a *AuthUseCase) VerifyLoginOTP(ctx context.Context, req dto.OtpVerificationReq) (*dto.OtpVerificationRes, error) {
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
func (a *AuthUseCase) RegisterUser(ctx context.Context, user dto.RegisterRequest) (*dto.UserRegistrationRes, error) {
	// TODO: Validate struct
	res, err := a.authService.CreateUser(user)
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

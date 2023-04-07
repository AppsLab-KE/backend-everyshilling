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

// RegisterUser Implements authservice to register a new user
func (a *AuthUseCase) RegisterUser(ctx context.Context, user dto.RegisterRequest) (*dto.UserRegistrationRes, error) {
	// TODO: Validate struct
	res, err := a.authService.CreateUser(user)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (a *AuthUseCase) ResetPassword(ctx context.Context, user dto.RequestResetCredentials) (*dto.ResetRes, error) {
	// Send OTP

	return &dto.ResetRes{}, nil
}

func (a *AuthUseCase) LoginUser(ctx context.Context, user dto.RequestLogin) (*dto.UserLoginRes, error) {
	//add specified login credentials
	//apply the verified
	//implement jwt

	return &dto.UserLoginRes{}, nil
}

func (a *AuthUseCase) VerifyLoginOTP(ctx context.Context, req dto.OtpVerificationReq) (*dto.OtpVerificationRes, error) {
	//add specified login credentials
	//apply the verified
	//implement jwt

	return &dto.OtpVerificationRes{}, nil
}

func (a *AuthUseCase) VerifyResetOTP(ctx context.Context, uuid string, body dto.VerifyResetOTPJSONRequestBody) (interface{}, interface{}) {
	return &dto.VerifyResetOTPJSONRequestBody{}, nil
}

func (a *AuthUseCase) ChangePassword(ctx context.Context, uuid string, body dto.ChangePasswordJSONRequestBody) (interface{}, interface{}) {
	return dto.ChangePasswordJSONRequestBody{}, nil
}

func NewAuthUsecase(as adapters.AuthService, ss adapters.SessionService) *AuthUseCase {
	return &AuthUseCase{
		authService:    as,
		sessionService: ss,
	}
}

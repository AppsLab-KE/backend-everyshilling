package usecase

import (
	"context"

	"github.com/AppsLab-KE/backend-everyshilling/services/app-authentication/internal/core/adapters"
	"github.com/AppsLab-KE/backend-everyshilling/services/app-authentication/internal/core/entity"
	"github.com/AppsLab-KE/backend-everyshilling/services/app-authentication/internal/dto"
)

type AuthUseCase struct {
	authService    adapters.AuthService
	sessionService adapters.SessionService
}

func (a *AuthUseCase) RegisterUser(ctx context.Context, user *dto.RegisterRequest) (*entity.User, error) {
	// check if user exists
	// if exists, return error

	return &entity.User{}, nil
}

func (a *AuthUseCase) ResetPassword(ctx context.Context, user *dto.RequestResetCredentials) (*entity.User, error) {
	//check if the passwords match
	//if they match ,return an error
	//if RequestResetCredentials.Password == nil || RequestResetCredentials.ConfirmPassword == nil || *RequestResetCredentials.Password != *RequestResetCredentials.ConfirmPassword {
	//	return errors.New("passwords do not match")

	return &entity.User{}, nil
}

func NewAuthUsecase(as adapters.AuthService, ss adapters.SessionService) *AuthUseCase {
	return &AuthUseCase{
		authService:    as,
		sessionService: ss,
	}
}

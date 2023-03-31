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

func NewAuthUsecase(as adapters.AuthService, ss adapters.SessionService) *AuthUseCase {
	return &AuthUseCase{
		authService:    as,
		sessionService: ss,
	}
}

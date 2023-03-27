package usecase

import "github.com/AppsLab-KE/backend-everyshilling/services/app-authentication/internal/core/adapters"

type AuthUseCase struct {
	authService    adapters.AuthService
	sessionService adapters.SessionService
}

func NewAuthUsecase(as adapters.AuthService, ss adapters.SessionService) *AuthUseCase {
	return &AuthUseCase{
		authService:    as,
		sessionService: ss,
	}
}

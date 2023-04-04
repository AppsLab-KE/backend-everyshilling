package dto

import "github.com/AppsLab-KE/backend-everyshilling/services/app-authentication/internal/core/entity"

type DefaultRes[T UserRegistrationRes | UserLoginRes | ResetRes] struct {
	Message string `json:"message"`
	Error   string `json:"errors"`
	Code    int    `json:"code"`
	Data    T      `json:"data"`
}

type UserRegistrationRes struct {
	entity.User
	Token string
}

type UserLoginRes struct {
	entity.User
	Token string
}

type ResetRes struct {
}

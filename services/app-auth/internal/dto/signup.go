package dto

import "github.com/AppsLab-KE/backend-everyshilling/services/app-authentication/internal/core/entity"

// RegisterReq defines model for RegisterRequest.
type RegisterReq struct {
	// Email valid email
	Email string `json:"email" binding:"required,email"`

	// Name name
	Name string `json:"name" binding:"required"`

	// Password strong password
	Password string `json:"password" binding:"required" binding:"required,min=8"`

	// PhoneNumber valid phone number
	PhoneNumber string `json:"phone_number" binding:"required,e164"`
}

type UserRegistrationRes struct {
	entity.User
}

type AccountVerificationRes struct {
	entity.User
	Token        string `json:"bearer_token"`
	RefreshToken string `json:"refresh_token"`
}

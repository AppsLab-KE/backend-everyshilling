package dto

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

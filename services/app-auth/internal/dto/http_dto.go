// Package dto provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen version v1.12.4 DO NOT EDIT.
package dto

// Email defines model for Email.
type Email struct {
	// Email valid email
	Email *string `json:"email,omitempty"`
}

// RegisterRequest defines model for RegisterRequest.
type RegisterRequest struct {
	// Email valid email
	Email string `json:"email"`

	// Name name
	Name string `json:"name"`

	// Password strong password
	Password string `json:"password"`

	// PhoneNumber valid phone number
	PhoneNumber string `json:"phone number"`
}

// RequestLogin defines model for RequestLogin.
type RequestLogin struct {
	// Phone User's phone number
	Phone string `json:"phone"`
}

// RequestOTP defines model for RequestOTP.
type RequestOTP struct {
	// OtpCode Generated OTP
	OtpCode *string `json:"otp_code,omitempty"`
}

// RequestResetCredentials defines model for RequestResetCredentials.
type RequestResetCredentials struct {
	ConfirmPassword *string `json:"confirm_password,omitempty"`
	Password        *string `json:"password,omitempty"`
}

// Response defines model for Response.
type Response struct {
	Code    *float32                `json:"code,omitempty"`
	Data    *map[string]interface{} `json:"data,omitempty"`
	Error   *string                 `json:"error,omitempty"`
	Message *string                 `json:"message,omitempty"`
}

// ResponseError defines model for ResponseError.
type ResponseError struct {
	Code    *string `json:"code,omitempty"`
	Error   *string `json:"error,omitempty"`
	Message *string `json:"message,omitempty"`
}

// ResponseSuccess defines model for ResponseSuccess.
type ResponseSuccess struct {
	Code    *string `json:"code,omitempty"`
	Error   *string `json:"error,omitempty"`
	Message *string `json:"message,omitempty"`
}

// RegisterJSONRequestBody defines body for Register for application/json ContentType.
type RegisterJSONRequestBody = RegisterRequest

// ResetJSONRequestBody defines body for Reset for application/json ContentType.
type ResetJSONRequestBody = Email

// ChangePasswordJSONRequestBody defines body for ChangePassword for application/json ContentType.
type ChangePasswordJSONRequestBody = RequestResetCredentials

// VerifyResetOTPJSONRequestBody defines body for VerifyResetOTP for application/json ContentType.
type VerifyResetOTPJSONRequestBody = RequestOTP

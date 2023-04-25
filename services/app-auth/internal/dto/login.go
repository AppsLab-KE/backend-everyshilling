package dto

import "github.com/AppsLab-KE/backend-everyshilling/services/app-authentication/internal/core/entity"

type LoginInitReq struct {
	PhoneNumber string `json:"phone_number" binding:"required"`
	Password    string `json:"password" binding:"required"`
}

type LoginInitRes struct {
	StatusCode   int32  `json:"-"`
	Message      string `json:"message,omitempty"`
	TrackingUuid string `json:"tracking_uuid,omitempty"`
}

type LoginOTPBody struct {
	OtpCode string `json:"otp_code" binding:"required"`
}

type ResendOTPReq struct {
	TrackingUID string `json:"tracking_uid"`
}

type ResendOTPRes struct {
	StatusCode   int32  `json:"-"`
	Message      string `json:"message,omitempty"`
	TrackingUuid string `json:"tracking_uuid,omitempty"`
}

type LoginRes struct {
	StatusCode int32        `json:"-"`
	Message    string       `json:"message,omitempty"`
	Token      string       `json:"token"`
	User       *entity.User `json:"user"`
}

type UserLoginRes struct {
	entity.User
	Token string
}

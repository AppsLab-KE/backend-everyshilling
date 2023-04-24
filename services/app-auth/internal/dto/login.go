package dto

import "github.com/AppsLab-KE/backend-everyshilling/services/app-authentication/internal/core/entity"

type LoginInitReq struct {
	Phone    string `json:"phone"`
	Password string `json:"password"`
}

type LoginInitRes struct {
	StatusCode   int32  `json:"-"`
	Message      string `json:"message,omitempty"`
	TrackingUuid string `json:"tracking_uuid,omitempty"`
}

type LoginOTPBody struct {
	OtpCode string `json:"otp_code"`
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

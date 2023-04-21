package dto

import "github.com/AppsLab-KE/backend-everyshilling/services/app-authentication/internal/core/entity"

type LoginInitReq struct {
	Phone    string
	Password string
}

type LoginInitRes struct {
	StatusCode   int32  `json:"status_code,omitempty"`
	Message      string `json:"message,omitempty"`
	TrackingUuid string `json:"tracking_uuid,omitempty"`
}

type LoginOTPBody struct {
	OtpCode string `json:"otp_code"`
}

type ResendLoginOTPReq struct {
	TrackingUID string
}

type ResendLoginOTPRes struct {
	StatusCode   int32  `json:"status_code,omitempty"`
	Message      string `json:"message,omitempty"`
	TrackingUuid string `json:"tracking_uuid,omitempty"`
}

type UserLoginRes struct {
	entity.User
	Token string
}

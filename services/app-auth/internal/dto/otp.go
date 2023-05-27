package dto

type OtpGenReq struct {
	Phone string `json:"phone"`
}

type AccountVerificationOTPGenReq struct {
	Phone    string `json:"phone"`
	UserUUID string `json:"user_uuid"`
}

type OtpGenRes struct {
	StatusCode   int32  `json:"-,omitempty"`
	Message      string `json:"-,omitempty"`
	TrackingUuid string `json:"tracking_uuid,omitempty"`
}

type OtpVerificationReq struct {
	TrackingUID string
	OtpCode     string `json:"otp_code"`
}

type OtpVerificationRes struct {
	StatusCode int32  `json:"-,omitempty"`
	Message    string `json:"-,omitempty"`
}

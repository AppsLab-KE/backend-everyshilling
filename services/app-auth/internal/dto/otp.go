package dto

type OtpGenReq struct {
	Phone string
}

type OtpGenRes struct {
	StatusCode   int32  `json:"status_code,omitempty"`
	Message      string `json:"message,omitempty"`
	TrackingUuid string `json:"tracking_uuid,omitempty"`
}

type OtpVerificationReq struct {
	TrackingUID string
	OtpCode     string
}

type OtpVerificationRes struct {
	StatusCode int32  `json:"status_code,omitempty"`
	Message    string `json:"message,omitempty"`
}

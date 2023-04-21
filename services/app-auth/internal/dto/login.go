package dto

type LoginInitReq struct {
	Phone    string
	Password string
}

type LoginInitRes struct {
	StatusCode   int32  `json:"status_code,omitempty"`
	Message      string `json:"message,omitempty"`
	TrackingUuid string `json:"tracking_uuid,omitempty"`
}

type ResendLoginOTPReq struct {
	TrackingUID string
}

type ResendLoginOTPRes struct {
	StatusCode   int32  `json:"status_code,omitempty"`
	Message      string `json:"message,omitempty"`
	TrackingUuid string `json:"tracking_uuid,omitempty"`
}

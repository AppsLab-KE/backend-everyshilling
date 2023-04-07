package dto

type OtpVerificationReq struct {
	TrackingUID string
	OtpCode     string
}

type OtpVerificationRes struct {
}

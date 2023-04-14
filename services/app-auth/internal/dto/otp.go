package dto

type OtpGenReq struct {
}

type OtpGenRes struct {
}

type OtpVerificationReq struct {
	TrackingUID string
	OtpCode     string
}

type OtpVerificationRes struct {
}

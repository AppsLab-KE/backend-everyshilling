package dto

type OtpReq struct {
	PhoneNumber string `json:"phone_number"`
}

type ValidateOtpReq struct {
	PhoneNumber string `json:"phone_number"`
	OtpCode     string `json:"otp_code"`
}

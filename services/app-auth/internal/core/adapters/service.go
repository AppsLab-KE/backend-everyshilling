package adapters

import "github.com/AppsLab-KE/backend-everyshilling/services/app-authentication/internal/dto"

type SessionService interface {
	Invalidate() dto.DefaultRes[interface{}]
}

type AuthService interface {

	// SendResetOTP sends an OTP to the user's registered phone number or email to initiate the password reset process.
	// It takes in a dto.OtpGenReq as input and returns a dto.OtpGenRes and an error as output.
	SendResetOTP(request dto.OtpGenReq) (*dto.OtpGenRes, error)

	// VerifyResetOTP verifies the OTP entered by the user to reset the password.
	// It takes in a dto.OtpVerificationReq as input and returns a dto.OtpVerificationRes and an error as output.
	VerifyResetOTP(request dto.OtpVerificationReq) (*dto.OtpVerificationRes, error)

	// ChangePassword updates the user's password after successful verification of the OTP sent to the user.
	// It takes in a dto.RequestResetCredentials as input and returns a dto.ResetRes and an error as output.
	ChangePassword(request dto.RequestResetCredentials) (*dto.ResetRes, error)

	// SendLoginOtp sends an OTP to the user's registered phone number or email to initiate the login process.
	// It takes in a dto.LoginInitReq as input and returns a dto.LoginInitRes and an error as output.
	SendLoginOtp(request dto.LoginInitReq) (*dto.LoginInitRes, error)

	// VerifyLoginOtp verifies the OTP entered by the user to complete the login process.
	// It takes in a dto.OtpVerificationReq as input and returns a dto.OtpVerificationRes and an error as output.
	VerifyLoginOtp(request dto.OtpVerificationReq) (*dto.OtpVerificationRes, error)

	// ResendLoginOTP sends a new OTP to the user's registered phone number or email in case the previous OTP expired or the user did not receive it.
	// It takes in a dto.ResendLoginOTPReq as input and returns a dto.ResendLoginOTPRes and an error as output.
	ResendLoginOTP(request dto.ResendLoginOTPReq) (*dto.ResendLoginOTPRes, error)

	// CreateUser creates a new user account with the details provided in the dto.RegisterRequest.
	// It returns a dto.UserRegistrationRes and an error as output.
	CreateUser(registerRequest dto.RegisterRequest) (*dto.UserRegistrationRes, error)
	// SendVerifyPhoneOTP sends an OTP to the user's registered phone number or email to initiate the phone number verification process.
	// It takes in a dto.OtpVerificationReq as input and returns a dto.OtpVerificationRes and an error as output.
	SendVerifyPhoneOTP(verificationRequest dto.OtpVerificationReq) (*dto.OtpVerificationRes, error)
	// VerifyPhoneOTP verifies the OTP entered by the user to complete the phone number verification process.
	// It takes in a dto.OtpVerificationReq as input and returns a dto.OtpVerificationRes and an error as output.
	VerifyPhoneOTP(verificationRequest dto.OtpVerificationReq) (*dto.OtpVerificationRes, error)

	// ResendVerifyPhoneOTP sends a new OTP to the user's registered phone number in case the previous OTP expired or the user did not receive it.
	// It takes in a dto.ResendLoginOTPReq as input and returns a dto.ResendLoginOTPRes and an error as output.
	ResendVerifyPhoneOTP(request dto.ResendLoginOTPReq) (*dto.ResendLoginOTPRes, error)
}

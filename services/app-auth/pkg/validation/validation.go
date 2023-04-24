package validation

import (
	"regexp"
	"strings"
)

var (
	emailRegexp   = regexp.MustCompile("^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\\.[a-zA-Z]{2,}$")
	phoneRegexp   = regexp.MustCompile("^(\\+?254|0)[7821][0-9]{8}$")
	otpRegex      = regexp.MustCompile("^\\d{6}$")
	passwordRegex = regexp.MustCompile("^.{8,}$")
)

func ValidateEmail(email string) bool {
	if len(email) == 0 {
		return false
	}

	email = strings.TrimSpace(email)

	if emailRegexp != nil {
		return emailRegexp.MatchString(email)
	}
	return false
}

func ValidateOTP(otp string) bool {
	if len(otp) != 6 {
		return false
	}

	otp = strings.TrimSpace(otp)
	if otpRegex != nil {
		return otpRegex.MatchString(otp)
	}
	return false
}

func ValidatePassword(password string) bool {
	if len(password) < 8 {
		return false
	}
	if passwordRegex != nil {
		return passwordRegex.MatchString(password)
	}
	return false
}

func ValidatePhone(phone string) bool {
	if len(phone) < 10 {
		return false
	}

	phone = strings.TrimSpace(phone)
	if phoneRegexp != nil {
		return phoneRegexp.MatchString(phone)
	}

	return false
}

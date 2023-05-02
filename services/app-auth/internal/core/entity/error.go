package entity

import "errors"

type ValidationError error

func NewValidationError(err string) ValidationError {
	return errors.New(err)
}

type OTPError error

func NewOTPError(err string) OTPError {
	return errors.New(err)
}

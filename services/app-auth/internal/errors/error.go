package errors

import "errors"

var (
	ErrUserExists      = errors.New("user exists")
	ErrUserCreation    = errors.New("io error while creating user. see server logs")
	ErrValidationError = errors.New("validation error")
)

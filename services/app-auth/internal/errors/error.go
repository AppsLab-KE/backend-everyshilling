package errors

const (
	userExists      = "user exists"
	userCreation    = "io error while creating user. see server logs"
	validationError = "validation error"
)

type ErrUserExists struct{}
type ErrUserCreation struct{}
type ErrRequestValidation struct{}

func (e *ErrUserExists) Error() string {
	return userExists
}

func (e *ErrUserCreation) Error() string {
	return userCreation
}

func (e *ErrRequestValidation) Error() string {
	return validationError
}

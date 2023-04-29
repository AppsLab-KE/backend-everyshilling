package entity

func NewValidationError(err string) *ValidationError {
	return &ValidationError{
		msg: err,
	}
}

type ValidationError struct {
	msg string
}

func (v *ValidationError) Error() string {
	return v.msg
}

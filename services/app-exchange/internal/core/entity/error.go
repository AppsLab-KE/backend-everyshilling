package entity

import "errors"

type ValidationError error

func NewValidationError(err string) ValidationError {
	return errors.New(err)
}

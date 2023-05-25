package repository

import (
	"errors"
	"github.com/AppsLab-KE/backend-everyshilling/services/app-exchange/internal/core/adapters"
)

var (
	ErrUserNotFound = errors.New("user not found")
)

type authRepo struct {
	currencyStorage adapters.CurrencyStorage
}

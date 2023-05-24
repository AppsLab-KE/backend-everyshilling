package ports

import (
	"github.com/AppsLab-KE/backend-everyshilling/services/app-ratescron/internal/core/model"
)

type ConversionRateRepository interface {
	// GetConversionRate returns the conversion rate for the given base currency
	GetConversionRate(baseCurrency string) ([]model.ConversionRate, error)
}

type ExchangeRateAPIStorage interface {
	// GetRates returns the exchange rates for the given base currency
	GetRates(baseCurrency string) ([]model.ConversionRate, error)
}

type FixerAPIStorage interface {
	// GetRates returns the exchange rates for the given base currency
	GetRates(baseCurrency string) ([]model.ConversionRate, error)
}

type RatesAPIRepository interface {
	// GetRates returns the exchange rates for the given base currency
	GetRates(baseCurrency string) ([]model.ConversionRate, error)
}

package repository

import (
	"github.com/AppsLab-KE/backend-everyshilling/services/app-ratescron/config"
	"github.com/AppsLab-KE/backend-everyshilling/services/app-ratescron/internal/core/model"
	"github.com/AppsLab-KE/backend-everyshilling/services/app-ratescron/internal/core/ports"
	"github.com/AppsLab-KE/backend-everyshilling/services/app-ratescron/internal/core/storage"
)

type store interface {
	GetRates(baseCurrency string) ([]model.ConversionRate, error)
}

type ratesApi struct {
	storage store
}

func (r ratesApi) GetRates(baseCurrency string) ([]model.ConversionRate, error) {
	return r.storage.GetRates(baseCurrency)
}

func NewRatesAPIRepository(exchangeRateAPI config.ExchangeApi) ports.RatesAPIRepository {
	var requestedStore store
	switch exchangeRateAPI.Name {
	case "fixer":
		requestedStore = storage.NewFixerAPIStorage(exchangeRateAPI.APIKey)
	case "exchange-rate":
		requestedStore = storage.NewExchangeRateStorage(exchangeRateAPI.APIKey)
	}
	return &ratesApi{
		storage: requestedStore,
	}
}

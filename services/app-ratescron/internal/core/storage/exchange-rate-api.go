package storage

import (
	"github.com/AppsLab-KE/backend-everyshilling/services/app-ratescron/internal/core/model"
	"github.com/AppsLab-KE/backend-everyshilling/services/app-ratescron/internal/core/ports"
	"github.com/AppsLab-KE/backend-everyshilling/services/app-ratescron/internal/platform"
)

type exchangeRateAPIStorage struct {
	api platform.ExchangeRateAPI
}

func (e exchangeRateAPIStorage) GetRates(baseCurrency string) ([]model.ConversionRate, error) {
	apiResponse, err := e.api.GetRates(baseCurrency)
	if err != nil {
		return nil, err
	}
	return apiResponse.ExtractedRates(), nil
}

func NewExchangeRateStorage(apiKey string) ports.ExchangeRateAPIStorage {
	return &exchangeRateAPIStorage{api: platform.NewExchangeRateAPI(apiKey)}
}

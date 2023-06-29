package storage

import (
	"github.com/AppsLab-KE/backend-everyshilling/services/app-ratescron/internal/core/ports"
	"github.com/AppsLab-KE/backend-everyshilling/services/app-ratescron/internal/dto"
	"github.com/AppsLab-KE/backend-everyshilling/services/app-ratescron/internal/platform"
)

type fixerAPIStorage struct {
	api platform.ExchangeRateAPI
}

func (e fixerAPIStorage) GetRates(baseCurrency string) (*dto.ExchangeRateApiResponse, error) {
	return e.api.GetRates(baseCurrency)
}

func NewFixerAPIStorage(apiKey string) ports.FixerAPIStorage {
	return &exchangeRateAPIStorage{api: platform.NewFixerAPI(apiKey)}
}

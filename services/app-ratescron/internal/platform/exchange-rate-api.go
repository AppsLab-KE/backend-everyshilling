package platform

import (
	"encoding/json"
	"fmt"
	"github.com/AppsLab-KE/backend-everyshilling/services/app-ratescron/internal/dto"
	"net/http"
)

const (
	ExchangeRateAPIURL = "https://v6.exchangerate-api.com/v6/%s/latest/%s"
)

// ExchangeRateAPI is the interface for the exchange rate API
type ExchangeRateAPI interface {
	GetRates(baseCurrency string) (*dto.ExchangeRateApiResponse, error)
}

// NewExchangeRateAPI returns a new instance of the exchange rate  ExchangeRateAPI interface
func NewExchangeRateAPI(apiKey string) ExchangeRateAPI {
	return &ExchangeRateAPIImpl{
		APIKey: apiKey,
	}
}

// ExchangeRateAPIImpl is the implementation of the ExchangeRateAPI interface
type ExchangeRateAPIImpl struct {
	APIKey string
}

// GetRates returns the exchange rates for the given base currency
func (e *ExchangeRateAPIImpl) GetRates(baseCurrency string) (*dto.ExchangeRateApiResponse, error) {
	url := fmt.Sprintf(ExchangeRateAPIURL, e.APIKey, baseCurrency)
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	var apiResponse dto.ExchangeRateApiResponse
	err = json.NewDecoder(resp.Body).Decode(&apiResponse)
	if err != nil {
		return nil, err
	}
	return &apiResponse, nil
}

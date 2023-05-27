package platform

import (
	"fmt"
	"github.com/AppsLab-KE/backend-everyshilling/services/app-ratescron/internal/dto"
	"github.com/goccy/go-json"
	"net/http"
)

const (
	FixerAPIURL = "http://data.fixer.io/api/latest?access_key=%s&base=%s"
)

// FixerAPI is the interface for fixer API
type FixerAPI interface {
	GetRates(baseCurrency string) (*dto.FixerAPIResponse, error)
}

// NewFixerAPI returns a new instance of the fixer API
func NewFixerAPI(apiKey string) ExchangeRateAPI {
	return &ExchangeRateAPIImpl{
		APIKey: apiKey,
	}
}

// FixerAPIImpl is the implementation of the fixer API
type FixerAPIImpl struct {
	APIKey string
}

// GetRates returns the exchange rates for the given base currency
func (e *FixerAPIImpl) GetRates(baseCurrency string) (*dto.FixerAPIResponse, error) {
	url := fmt.Sprintf(FixerAPIURL, e.APIKey, baseCurrency)
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	var apiResponse dto.FixerAPIResponse
	err = json.NewDecoder(resp.Body).Decode(&apiResponse)
	if err != nil {
		return nil, err
	}
	return &apiResponse, nil
}

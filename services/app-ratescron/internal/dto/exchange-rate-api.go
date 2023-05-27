package dto

import "github.com/AppsLab-KE/backend-everyshilling/services/app-ratescron/internal/core/model"

// ExchangeRateApiResponse is the response from the exchange rate API
// https://www.exchangerate-api.com/
type ExchangeRateApiResponse struct {
	Result             string             `json:"result"`
	Documentation      string             `json:"documentation"`
	TermsOfUse         string             `json:"terms_of_use"`
	TimeLastUpdateUnix int64              `json:"time_last_update_unix"`
	TimeLastUpdateUTC  string             `json:"time_last_update_utc"`
	TimeNextUpdateUnix int64              `json:"time_next_update_unix"`
	TimeNextUpdateUTC  string             `json:"time_next_update_utc"`
	BaseCode           string             `json:"base_code"`
	ConversionRates    map[string]float64 `json:"conversion_rates"`
}

// ExtractedRates returns the conversion rates from the API response
func (r *ExchangeRateApiResponse) ExtractedRates() []model.ConversionRate {
	rates := make([]model.ConversionRate, 0)
	for k, v := range r.ConversionRates {
		rates = append(rates, model.ConversionRate{
			FromCurrency:       r.BaseCode,
			ToCurrency:         k,
			DateUpdatedUnixUTC: r.TimeLastUpdateUnix,
			Rate:               v,
		})
	}
	return rates
}

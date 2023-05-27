package dto

import "github.com/AppsLab-KE/backend-everyshilling/services/app-ratescron/internal/core/model"

type FixerAPIResponse struct {
	Success   bool               `json:"success"`
	Timestamp int64              `json:"timestamp"`
	Base      string             `json:"base"`
	Date      string             `json:"date"`
	Rates     map[string]float64 `json:"rates"`
}

func (r *FixerAPIResponse) ExtractedRates() []model.ConversionRate {
	rates := make([]model.ConversionRate, 0)
	for k, v := range r.Rates {
		rates = append(rates, model.ConversionRate{
			FromCurrency:       r.Base,
			ToCurrency:         k,
			DateUpdatedUnixUTC: r.Timestamp,
			Rate:               v,
		})
	}
	return rates
}

package model

type ConversionRate struct {
	FromCurrency       string
	ToCurrency         string
	Rate               float64
	DateUpdatedUnixUTC int64
}

package models

type ConversionRate struct {
	BaseModel
	FromCurrency string  `json:"from_currency" gorm:"column:from_currency"`
	ToCurrency   string  `json:"to_currency" gorm:"column:to_currency"`
	Rate         float64 `json:"rate" gorm:"column:rate"`
	TimeStampUTC int64   `json:"time_stamp_utc" gorm:"column:time_stamp_utc"`
}

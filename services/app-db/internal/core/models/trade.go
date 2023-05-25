package models

type Trade struct {
	TradeId      string `json:"trade_id"`
	AccountId    string `json:"account_id"`
	TradeType    string `json:"trade_type"`
	TradeStatus  string `json:"trade_status"`
	FromCurrency string `json:"from_currency"`
	ToCurrency   string `json:"to_currency"`
	FromAmount   int64  `json:"from_amount"`
	FinalAmount  int64  `json:"final_amount"`
}

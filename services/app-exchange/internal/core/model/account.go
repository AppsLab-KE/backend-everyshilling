package model

type Account struct {
	AccountId       string `json:"account_id"`
	UserId          string `json:""`
	Balance         int64  `json:"balance"`
	BaseCurrency    string `json:"base_currency"`
	CreatedAt       int64  `json:"created_at"`
	ParentAccountId string `json:"parent_account_id"`
}

type Transaction struct {
	TransactionId          string `json:"transaction_id"`
	AccountId              string `json:"account_id"`
	Amount                 int64  `json:"amount"`
	TransactionType        string `json:"transaction_type"`
	TransactionStatus      string `json:"transaction_status"`
	TransactionCode        string `json:"transaction_code"`
	TransactionDescription string `json:"transaction_description"`
	CreatedAt              int64  `json:"created_at"`
}

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

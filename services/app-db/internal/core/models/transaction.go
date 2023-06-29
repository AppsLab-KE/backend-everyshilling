package models

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

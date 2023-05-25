package entity

type AccountOverviewResponse struct {
	Balance            *float32                `json:"balance"`
	Data               *AccountOverviewData    `json:"data"`
	PaymentDetails     *PaymentDetails         `json:"paymentDetails"`
	TransactionSummary *map[string]interface{} `json:"transactionSummary"`
}

type AccountOverviewData struct {
	Balance            *float32            `json:"balance"`
	PaymentDetails     *PaymentDetails     `json:"paymentDetails,omitempty"`
	TransactionSummary *TransactionSummary `json:"transactionSummary,omitempty"`
}

type PaymentDetails struct {
	AmountPaid       *float32 `json:"amountPaid,omitempty"`
	RemainingBalance *float32 `json:"remainingBalance,omitempty"`
}

type TransactionSummary struct {
	Date        *string `json:"date,omitempty"`
	Description *string `json:"description,omitempty"`
}

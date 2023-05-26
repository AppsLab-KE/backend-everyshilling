package models

type Account struct {
	AccountId       string `json:"account_id"`
	UserId          string `json:""`
	Balance         int64  `json:"balance"`
	BaseCurrency    string `json:"base_currency"`
	CreatedAt       int64  `json:"created_at"`
	ParentAccountId string `json:"parent_account_id"`
	AccountID       string `json:"account_id"`
}

func (a Account) Error() string {
	//TODO implement me
	panic("implement me")
}

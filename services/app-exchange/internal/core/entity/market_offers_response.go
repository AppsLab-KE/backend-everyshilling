package entity

type MarketplaceOffersResponse struct {
	Currencies *Currencies                     `json:"currencies,omitempty"`
	Data       []*MarketplaceOffersResponseObj `json:"data,omitempty"`
}

type Currencies struct {
	KSH *float32 `json:"KSH,omitempty"`
	USD *float32 `json:"USD,omitempty"`
}

type MarketplaceOffersResponseObj struct {
	Currency *string  `json:"currency,omitempty"`
	Rate     *float32 `json:"rate,omitempty"`
}

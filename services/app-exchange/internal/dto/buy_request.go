package dto

// BuyRequest defines the structure for a buy request.
type BuyReq struct {
	Amount   float32 `json:"amount"`
	Delivery string  `json:"delivery"`
	PayIn    string  `json:"payIn"`
	YouBuy   string  `json:"youBuy"`
}

// BuyReqData defines the structure for buy request data.
type BuyReqData struct {
	Code    float32     `json:"code"`
	Data    *BuyRequest `json:"data"`
	Error   string      `json:"error"`
	Message string      `json:"message"`
}

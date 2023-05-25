package dto

type TopUpReq struct {
	Amount        float32 `json:"amount"`
	PaymentMethod string  `json:"paymentMethod"`
}

// TopUpReqData  defines the structure for top-up request data.
type TopUpReqData struct {
	Code    float32       `json:"code"`
	Data    *TopUpRequest `json:"data"`
	Error   string        `json:"error"`
	Message string        `json:"message"`
}

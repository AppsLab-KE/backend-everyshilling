package dto

type DefaultRes struct {
	Message string      `json:"message"`
	Error   string      `json:"error"`
	Code    int         `json:"code"`
	Data    interface{} `json:"data"`
}

package dto

type GeneralResponse[T any] struct {
	Code    int    `json:"code"`
	Data    T      `json:"data"`
	Error   string `json:"errors"`
	Message string `json:"message"`
}

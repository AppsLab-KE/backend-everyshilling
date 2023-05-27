package dto

type DefaultRes[T any] struct {
	Message string `json:"message"`
	Error   string `json:"errors"`
	Code    int    `json:"status_code"`
	Data    T      `json:"data"`
}

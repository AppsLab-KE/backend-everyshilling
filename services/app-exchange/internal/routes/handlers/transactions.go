package handlers

import (
	"context"
	"errors"
	"github.com/AppsLab-KE/be-go-gen-grpc/exchange"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (h Handler) CreateTransaction(ctx context.Context, request *exchange.CreateTransactionRequest) (*exchange.CreateTransactionResponse, error) {
	return nil, errors.New("not supported")
}

func (h Handler) DeleteTransaction(ctx context.Context, request *exchange.DeleteTransactionRequest) (*emptypb.Empty, error) {
	return nil, errors.New("not supported")

}

func (h Handler) UpdateTransaction(ctx context.Context, request *exchange.UpdateTransactionRequest) (*emptypb.Empty, error) {
	return nil, errors.New("not supported")

}

func (h Handler) GetTransaction(ctx context.Context, request *exchange.GetTransactionRequest) (*exchange.Transaction, error) {
	return nil, errors.New("not supported")
}

func (h Handler) GetTransactionByAccount(ctx context.Context, request *exchange.GetTransactionByAccountRequest) (*exchange.GetTransactionByAccountResponse, error) {
	return nil, errors.New("not supported")
}

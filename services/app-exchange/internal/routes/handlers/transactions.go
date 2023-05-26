package handlers

import (
	"context"
	"github.com/AppsLab-KE/be-go-gen-grpc/exchange"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (h Handler) CreateTransaction(ctx context.Context, request *exchange.CreateTransactionRequest) (*exchange.CreateTransactionResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (h Handler) DeleteTransaction(ctx context.Context, request *exchange.DeleteTransactionRequest) (*emptypb.Empty, error) {
	//TODO implement me
	panic("implement me")
}

func (h Handler) UpdateTransaction(ctx context.Context, request *exchange.UpdateTransactionRequest) (*emptypb.Empty, error) {
	//TODO implement me
	panic("implement me")
}

func (h Handler) GetTransaction(ctx context.Context, request *exchange.GetTransactionRequest) (*exchange.Transaction, error) {
	//TODO implement me
	panic("implement me")
}

func (h Handler) GetTransactionByAccount(ctx context.Context, request *exchange.GetTransactionByAccountRequest) (*exchange.GetTransactionByAccountResponse, error) {
	//TODO implement me
	panic("implement me")
}

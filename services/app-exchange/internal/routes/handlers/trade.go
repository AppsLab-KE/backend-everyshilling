package handlers

import (
	"context"
	"github.com/AppsLab-KE/be-go-gen-grpc/exchange"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (h Handler) CreateTrade(ctx context.Context, request *exchange.CreateTradeRequest) (*exchange.CreateTradeResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (h Handler) DeleteTrade(ctx context.Context, request *exchange.DeleteTradeRequest) (*emptypb.Empty, error) {
	//TODO implement me
	panic("implement me")
}

func (h Handler) UpdateTrade(ctx context.Context, request *exchange.UpdateTradeRequest) (*emptypb.Empty, error) {
	//TODO implement me
	panic("implement me")
}

func (h Handler) GetTrade(ctx context.Context, request *exchange.GetTradeRequest) (*exchange.Trade, error) {
	//TODO implement me
	panic("implement me")
}

func (h Handler) GetTradeByAccount(ctx context.Context, request *exchange.GetTradeByAccountRequest) (*exchange.GetTradeByAccountResponse, error) {
	//TODO implement me
	panic("implement me")
}

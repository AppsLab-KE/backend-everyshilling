package handlers

import (
	"context"
	"github.com/AppsLab-KE/be-go-gen-grpc/db"
	"github.com/AppsLab-KE/be-go-gen-grpc/exchange"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (h Handler) CreateTrade(ctx context.Context, request *exchange.CreateTradeRequest) (*exchange.CreateTradeResponse, error) {
	if request == nil {
		return nil, ErrEmptyRequest
	}

	tradeRequest := &db.CreateTradeRequest{
		Trade: &db.Trade{
			AccountId:    request.Trade.AccountId,
			TradeType:    request.Trade.TradeType,
			TradeStatus:  request.Trade.TradeStatus,
			FromCurrency: request.Trade.FromCurrency,
			ToCurrency:   request.Trade.ToCurrency,
			FromAmount:   request.Trade.FromAmount,
		},
	}

	tradeResponse, err := h.tradeService.CreateTrade(ctx, tradeRequest)
	if err != nil {
		return nil, err
	}

	tradeResponseResult := &exchange.CreateTradeResponse{
		TradeId: tradeResponse.TradeId,
	}

	return tradeResponseResult, nil
}

func (h Handler) DeleteTrade(ctx context.Context, request *exchange.DeleteTradeRequest) (*emptypb.Empty, error) {
	if request == nil {
		return nil, ErrEmptyRequest
	}

	deleteRequest := &db.DeleteTradeRequest{
		TradeId: request.TradeId,
	}

	_, err := h.tradeService.DeleteTrade(ctx, deleteRequest)
	if err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}

func (h Handler) UpdateTrade(ctx context.Context, request *exchange.UpdateTradeRequest) (*emptypb.Empty, error) {
	if request == nil {
		return nil, ErrEmptyRequest
	}

	updateRequest := &db.UpdateTradeRequest{
		Trade: &db.Trade{
			TradeId:      request.Trade.TradeId,
			AccountId:    request.Trade.AccountId,
			TradeType:    request.Trade.TradeType,
			TradeStatus:  request.Trade.TradeStatus,
			FromCurrency: request.Trade.FromCurrency,
			ToCurrency:   request.Trade.ToCurrency,
			FromAmount:   request.Trade.FromAmount,
		},
	}

	_, err := h.tradeService.UpdateTrade(ctx, updateRequest)
	if err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}

func (h Handler) GetTrade(ctx context.Context, request *exchange.GetTradeRequest) (*exchange.Trade, error) {
	if request == nil {
		return nil, ErrEmptyRequest
	}

	getRequest := &db.GetTradeRequest{
		TradeId: request.TradeId,
	}

	tradeResponse, err := h.tradeService.GetTrade(ctx, getRequest)
	if err != nil {
		return nil, err
	}

	tradeResponseResult := &exchange.Trade{
		TradeId:      tradeResponse.TradeId,
		AccountId:    tradeResponse.AccountId,
		TradeType:    tradeResponse.TradeType,
		TradeStatus:  tradeResponse.TradeStatus,
		FromCurrency: tradeResponse.FromCurrency,
		ToCurrency:   tradeResponse.ToCurrency,
		FromAmount:   tradeResponse.FromAmount,
	}

	return tradeResponseResult, nil
}

func (h Handler) GetTradeByAccount(ctx context.Context, request *exchange.GetTradeByAccountRequest) (*exchange.GetTradeByAccountResponse, error) {
	if request == nil {
		return nil, ErrEmptyRequest
	}

	getRequest := &db.GetTradeByAccountRequest{
		AccountId: request.AccountId,
	}

	tradeResponse, err := h.tradeService.GetTradeByAccount(ctx, getRequest)
	if err != nil {
		return nil, err
	}

	tradeResponseResult := &exchange.GetTradeByAccountResponse{
		Trades: []*exchange.Trade{},
	}

	for _, trade := range tradeResponse.Trades {
		tradeResponseResult.Trades = append(tradeResponseResult.Trades, &exchange.Trade{
			TradeId:      trade.TradeId,
			AccountId:    trade.AccountId,
			TradeType:    trade.TradeType,
			TradeStatus:  trade.TradeStatus,
			FromCurrency: trade.FromCurrency,
			ToCurrency:   trade.ToCurrency,
			FromAmount:   trade.FromAmount,
		})
	}

	return tradeResponseResult, nil
}

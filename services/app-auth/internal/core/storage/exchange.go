package storage

import (
	"context"
	"github.com/AppsLab-KE/backend-everyshilling/services/app-authentication/internal/core/adapters"
	"github.com/AppsLab-KE/be-go-gen-grpc/exchange"
	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/grpc"
)

type exchangeStorageImpl struct {
	exchangeClient exchange.ExchangeServiceClient
}

func (e exchangeStorageImpl) CreateConversionRate(ctx context.Context, in *exchange.CreateConversionRateRequest, opts ...grpc.CallOption) (*exchange.CreateConversionRateResponse, error) {
	response, err := e.exchangeClient.CreateConversionRate(ctx, in, opts...)
	if err != nil {
		return nil, err
	}
	return response, nil
}

func (e exchangeStorageImpl) ReadConversionRate(ctx context.Context, in *exchange.ReadConversionRateRequest, opts ...grpc.CallOption) (*exchange.ReadConversionRateResponse, error) {
	response, err := e.exchangeClient.ReadConversionRate(ctx, in, opts...)
	if err != nil {
		return nil, err
	}
	return response, nil
}

func (e exchangeStorageImpl) UpdateConversionRate(ctx context.Context, in *exchange.UpdateConversionRateRequest, opts ...grpc.CallOption) (*exchange.UpdateConversionRateResponse, error) {
	response, err := e.exchangeClient.UpdateConversionRate(ctx, in, opts...)
	if err != nil {
		return nil, err
	}
	return response, nil
}

func (e exchangeStorageImpl) DeleteConversionRate(ctx context.Context, in *exchange.DeleteConversionRateRequest, opts ...grpc.CallOption) (*exchange.DeleteConversionRateResponse, error) {
	response, err := e.exchangeClient.DeleteConversionRate(ctx, in, opts...)
	if err != nil {
		return nil, err
	}
	return response, nil
}

func (e exchangeStorageImpl) CreateAccount(ctx context.Context, in *exchange.CreateAccountRequest, opts ...grpc.CallOption) (*exchange.CreateAccountResponse, error) {
	response, err := e.exchangeClient.CreateAccount(ctx, in, opts...)
	if err != nil {
		return nil, err
	}
	return response, nil
}

func (e exchangeStorageImpl) DeleteAccount(ctx context.Context, in *exchange.DeleteAccountRequest, opts ...grpc.CallOption) (*exchange.DeleteAccountResponse, error) {
	response, err := e.exchangeClient.DeleteAccount(ctx, in, opts...)
	if err != nil {
		return nil, err
	}
	return response, nil
}

func (e exchangeStorageImpl) UpdateAccount(ctx context.Context, in *exchange.UpdateAccountRequest, opts ...grpc.CallOption) (*exchange.UpdateAccountResponse, error) {
	response, err := e.exchangeClient.UpdateAccount(ctx, in, opts...)
	if err != nil {
		return nil, err
	}
	return response, nil
}

func (e exchangeStorageImpl) SearchAccount(ctx context.Context, in *exchange.SearchAccountRequest, opts ...grpc.CallOption) (*exchange.SearchAccountResponse, error) {
	response, err := e.exchangeClient.SearchAccount(ctx, in, opts...)
	if err != nil {
		return nil, err
	}
	return response, nil
}

func (e exchangeStorageImpl) CreateTransaction(ctx context.Context, in *exchange.CreateTransactionRequest, opts ...grpc.CallOption) (*exchange.CreateTransactionResponse, error) {
	response, err := e.exchangeClient.CreateTransaction(ctx, in, opts...)
	if err != nil {
		return nil, err
	}
	return response, nil
}

func (e exchangeStorageImpl) DeleteTransaction(ctx context.Context, in *exchange.DeleteTransactionRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	response, err := e.exchangeClient.DeleteTransaction(ctx, in, opts...)
	if err != nil {
		return nil, err
	}
	return response, nil
}

func (e exchangeStorageImpl) UpdateTransaction(ctx context.Context, in *exchange.UpdateTransactionRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	response, err := e.exchangeClient.UpdateTransaction(ctx, in, opts...)
	if err != nil {
		return nil, err
	}
	return response, nil
}

func (e exchangeStorageImpl) GetTransaction(ctx context.Context, in *exchange.GetTransactionRequest, opts ...grpc.CallOption) (*exchange.Transaction, error) {
	response, err := e.exchangeClient.GetTransaction(ctx, in, opts...)
	if err != nil {
		return nil, err
	}
	return response, nil
}

func (e exchangeStorageImpl) GetTransactionByAccount(ctx context.Context, in *exchange.GetTransactionByAccountRequest, opts ...grpc.CallOption) (*exchange.GetTransactionByAccountResponse, error) {
	response, err := e.exchangeClient.GetTransactionByAccount(ctx, in, opts...)
	if err != nil {
		return nil, err
	}
	return response, nil
}

func (e exchangeStorageImpl) CreateTrade(ctx context.Context, in *exchange.CreateTradeRequest, opts ...grpc.CallOption) (*exchange.CreateTradeResponse, error) {
	response, err := e.exchangeClient.CreateTrade(ctx, in, opts...)
	if err != nil {
		return nil, err
	}
	return response, nil
}

func (e exchangeStorageImpl) DeleteTrade(ctx context.Context, in *exchange.DeleteTradeRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	response, err := e.exchangeClient.DeleteTrade(ctx, in, opts...)
	if err != nil {
		return nil, err
	}
	return response, nil
}

func (e exchangeStorageImpl) UpdateTrade(ctx context.Context, in *exchange.UpdateTradeRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	response, err := e.exchangeClient.UpdateTrade(ctx, in, opts...)
	if err != nil {
		return nil, err
	}
	return response, nil
}

func (e exchangeStorageImpl) GetTrade(ctx context.Context, in *exchange.GetTradeRequest, opts ...grpc.CallOption) (*exchange.Trade, error) {
	response, err := e.exchangeClient.GetTrade(ctx, in, opts...)
	if err != nil {
		return nil, err
	}
	return response, nil
}

func (e exchangeStorageImpl) GetTradeByAccount(ctx context.Context, in *exchange.GetTradeByAccountRequest, opts ...grpc.CallOption) (*exchange.GetTradeByAccountResponse, error) {
	response, err := e.exchangeClient.GetTradeByAccount(ctx, in, opts...)
	if err != nil {
		return nil, err
	}
	return response, nil
}

func NewExchangeStorageImpl(exchangeClient exchange.ExchangeServiceClient) adapters.ExchangeStorage {
	return &exchangeStorageImpl{
		exchangeClient: exchangeClient,
	}
}

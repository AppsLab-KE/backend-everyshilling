package repository

import (
	"context"
	"github.com/AppsLab-KE/backend-everyshilling/services/app-authentication/internal/core/adapters"
	"github.com/AppsLab-KE/be-go-gen-grpc/exchange"
	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/grpc"
)

type ExchangeRepositoryImplementation struct {
	exchangeStorage adapters.ExchangeStorage
}

func (e ExchangeRepositoryImplementation) CreateConversionRate(ctx context.Context, in *exchange.CreateConversionRateRequest, opts ...grpc.CallOption) (*exchange.CreateConversionRateResponse, error) {
	// Call the corresponding method in exchangeStorage to create a conversion rate
	res, err := e.exchangeStorage.CreateConversionRate(ctx, in)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (e ExchangeRepositoryImplementation) ReadConversionRate(ctx context.Context, in *exchange.ReadConversionRateRequest, opts ...grpc.CallOption) (*exchange.ReadConversionRateResponse, error) {
	res, err := e.exchangeStorage.ReadConversionRate(ctx, in)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (e ExchangeRepositoryImplementation) UpdateConversionRate(ctx context.Context, in *exchange.UpdateConversionRateRequest, opts ...grpc.CallOption) (*exchange.UpdateConversionRateResponse, error) {
	res, err := e.exchangeStorage.UpdateConversionRate(ctx, in)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (e ExchangeRepositoryImplementation) DeleteConversionRate(ctx context.Context, in *exchange.DeleteConversionRateRequest, opts ...grpc.CallOption) (*exchange.DeleteConversionRateResponse, error) {
	res, err := e.exchangeStorage.DeleteConversionRate(ctx, in)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (e ExchangeRepositoryImplementation) CreateAccount(ctx context.Context, in *exchange.CreateAccountRequest, opts ...grpc.CallOption) (*exchange.CreateAccountResponse, error) {
	res, err := e.exchangeStorage.CreateAccount(ctx, in)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (e ExchangeRepositoryImplementation) DeleteAccount(ctx context.Context, in *exchange.DeleteAccountRequest, opts ...grpc.CallOption) (*exchange.DeleteAccountResponse, error) {
	res, err := e.exchangeStorage.DeleteAccount(ctx, in)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (e ExchangeRepositoryImplementation) UpdateAccount(ctx context.Context, in *exchange.UpdateAccountRequest, opts ...grpc.CallOption) (*exchange.UpdateAccountResponse, error) {
	res, err := e.exchangeStorage.UpdateAccount(ctx, in)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (e ExchangeRepositoryImplementation) SearchAccount(ctx context.Context, in *exchange.SearchAccountRequest, opts ...grpc.CallOption) (*exchange.SearchAccountResponse, error) {
	res, err := e.exchangeStorage.SearchAccount(ctx, in)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (e ExchangeRepositoryImplementation) CreateTransaction(ctx context.Context, in *exchange.CreateTransactionRequest, opts ...grpc.CallOption) (*exchange.CreateTransactionResponse, error) {
	res, err := e.exchangeStorage.CreateTransaction(ctx, in)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (e ExchangeRepositoryImplementation) DeleteTransaction(ctx context.Context, in *exchange.DeleteTransactionRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	res, err := e.exchangeStorage.DeleteTransaction(ctx, in)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (e ExchangeRepositoryImplementation) UpdateTransaction(ctx context.Context, in *exchange.UpdateTransactionRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	res, err := e.exchangeStorage.UpdateTransaction(ctx, in)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (e ExchangeRepositoryImplementation) GetTransaction(ctx context.Context, in *exchange.GetTransactionRequest, opts ...grpc.CallOption) (*exchange.Transaction, error) {
	res, err := e.exchangeStorage.GetTransaction(ctx, in)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (e ExchangeRepositoryImplementation) GetTransactionByAccount(ctx context.Context, in *exchange.GetTransactionByAccountRequest, opts ...grpc.CallOption) (*exchange.GetTransactionByAccountResponse, error) {
	res, err := e.exchangeStorage.GetTransactionByAccount(ctx, in)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (e ExchangeRepositoryImplementation) CreateTrade(ctx context.Context, in *exchange.CreateTradeRequest, opts ...grpc.CallOption) (*exchange.CreateTradeResponse, error) {
	res, err := e.exchangeStorage.CreateTrade(ctx, in)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (e ExchangeRepositoryImplementation) DeleteTrade(ctx context.Context, in *exchange.DeleteTradeRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	res, err := e.exchangeStorage.DeleteTrade(ctx, in)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (e ExchangeRepositoryImplementation) UpdateTrade(ctx context.Context, in *exchange.UpdateTradeRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	res, err := e.exchangeStorage.UpdateTrade(ctx, in)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (e ExchangeRepositoryImplementation) GetTrade(ctx context.Context, in *exchange.GetTradeRequest, opts ...grpc.CallOption) (*exchange.Trade, error) {
	res, err := e.exchangeStorage.GetTrade(ctx, in)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (e ExchangeRepositoryImplementation) GetTradeByAccount(ctx context.Context, in *exchange.GetTradeByAccountRequest, opts ...grpc.CallOption) (*exchange.GetTradeByAccountResponse, error) {
	res, err := e.exchangeStorage.GetTradeByAccount(ctx, in)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func NewExchangeRepositoryImplementation(exchangeStorage adapters.ExchangeStorage) adapters.ExchangeStorage {
	return &ExchangeRepositoryImplementation{
		exchangeStorage: exchangeStorage,
	}
}

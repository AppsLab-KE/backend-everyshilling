package usecase

import (
	"context"
	"github.com/AppsLab-KE/backend-everyshilling/services/app-authentication/internal/core/adapters"
	"github.com/AppsLab-KE/be-go-gen-grpc/exchange"
	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/grpc"
)

type ExchangeUseCase struct {
	exchangeService adapters.ExchangeService
}

func (e ExchangeUseCase) CreateConversionRate(ctx context.Context, in *exchange.CreateConversionRateRequest, opts ...grpc.CallOption) (*exchange.CreateConversionRateResponse, error) {
	// Call the exchange service method to create the conversion rate
	response, err := e.exchangeService.CreateConversionRate(ctx, in, opts...)
	if err != nil {
		return nil, err
	}
	return response, nil
}

func (e ExchangeUseCase) ReadConversionRate(ctx context.Context, in *exchange.ReadConversionRateRequest, opts ...grpc.CallOption) (*exchange.ReadConversionRateResponse, error) {
	// Call the exchange service method to read the conversion rate
	response, err := e.exchangeService.ReadConversionRate(ctx, in, opts...)
	if err != nil {
		return nil, err
	}
	return response, nil

}

func (e ExchangeUseCase) UpdateConversionRate(ctx context.Context, in *exchange.UpdateConversionRateRequest, opts ...grpc.CallOption) (*exchange.UpdateConversionRateResponse, error) {
	response, err := e.exchangeService.UpdateConversionRate(ctx, in, opts...)
	if err != nil {
		return nil, err
	}
	return response, nil
}

func (e ExchangeUseCase) DeleteConversionRate(ctx context.Context, in *exchange.DeleteConversionRateRequest, opts ...grpc.CallOption) (*exchange.DeleteConversionRateResponse, error) {
	response, err := e.exchangeService.DeleteConversionRate(ctx, in, opts...)
	if err != nil {
		return nil, err
	}
	return response, nil

}

func (e ExchangeUseCase) CreateAccount(ctx context.Context, in *exchange.CreateAccountRequest, opts ...grpc.CallOption) (*exchange.CreateAccountResponse, error) {
	response, err := e.exchangeService.CreateAccount(ctx, in, opts...)
	if err != nil {
		return nil, err
	}
	return response, nil
}

func (e ExchangeUseCase) DeleteAccount(ctx context.Context, in *exchange.DeleteAccountRequest, opts ...grpc.CallOption) (*exchange.DeleteAccountResponse, error) {
	response, err := e.exchangeService.DeleteAccount(ctx, in, opts...)
	if err != nil {
		return nil, err
	}
	return response, nil
}

func (e ExchangeUseCase) UpdateAccount(ctx context.Context, in *exchange.UpdateAccountRequest, opts ...grpc.CallOption) (*exchange.UpdateAccountResponse, error) {
	response, err := e.exchangeService.UpdateAccount(ctx, in, opts...)
	if err != nil {
		return nil, err
	}
	return response, nil
}

func (e ExchangeUseCase) SearchAccount(ctx context.Context, in *exchange.SearchAccountRequest, opts ...grpc.CallOption) (*exchange.SearchAccountResponse, error) {
	response, err := e.exchangeService.SearchAccount(ctx, in, opts...)
	if err != nil {
		return nil, err
	}
	return response, nil
}

func (e ExchangeUseCase) CreateTransaction(ctx context.Context, in *exchange.CreateTransactionRequest, opts ...grpc.CallOption) (*exchange.CreateTransactionResponse, error) {
	response, err := e.exchangeService.CreateTransaction(ctx, in, opts...)
	if err != nil {
		return nil, err
	}
	return response, nil
}

func (e ExchangeUseCase) DeleteTransaction(ctx context.Context, in *exchange.DeleteTransactionRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	response, err := e.exchangeService.DeleteTransaction(ctx, in, opts...)
	if err != nil {
		return nil, err
	}
	return response, nil
}

func (e ExchangeUseCase) UpdateTransaction(ctx context.Context, in *exchange.UpdateTransactionRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	response, err := e.exchangeService.UpdateTransaction(ctx, in, opts...)
	if err != nil {
		return nil, err
	}
	return response, nil
}

func (e ExchangeUseCase) GetTransaction(ctx context.Context, in *exchange.GetTransactionRequest, opts ...grpc.CallOption) (*exchange.Transaction, error) {
	response, err := e.exchangeService.GetTransaction(ctx, in, opts...)
	if err != nil {
		return nil, err
	}
	return response, nil
}

func (e ExchangeUseCase) GetTransactionByAccount(ctx context.Context, in *exchange.GetTransactionByAccountRequest, opts ...grpc.CallOption) (*exchange.GetTransactionByAccountResponse, error) {
	response, err := e.exchangeService.GetTransactionByAccount(ctx, in, opts...)
	if err != nil {
		return nil, err
	}
	return response, nil
}

func (e ExchangeUseCase) CreateTrade(ctx context.Context, in *exchange.CreateTradeRequest, opts ...grpc.CallOption) (*exchange.CreateTradeResponse, error) {
	response, err := e.exchangeService.CreateTrade(ctx, in, opts...)
	if err != nil {
		return nil, err
	}
	return response, nil
}

func (e ExchangeUseCase) DeleteTrade(ctx context.Context, in *exchange.DeleteTradeRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	response, err := e.exchangeService.DeleteTrade(ctx, in, opts...)
	if err != nil {
		return nil, err
	}
	return response, nil
}

func (e ExchangeUseCase) UpdateTrade(ctx context.Context, in *exchange.UpdateTradeRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	response, err := e.exchangeService.UpdateTrade(ctx, in, opts...)
	if err != nil {
		return nil, err
	}
	return response, nil
}

func (e ExchangeUseCase) GetTrade(ctx context.Context, in *exchange.GetTradeRequest, opts ...grpc.CallOption) (*exchange.Trade, error) {
	response, err := e.exchangeService.GetTrade(ctx, in, opts...)
	if err != nil {
		return nil, err
	}
	return response, nil
}

func (e ExchangeUseCase) GetTradeByAccount(ctx context.Context, in *exchange.GetTradeByAccountRequest, opts ...grpc.CallOption) (*exchange.GetTradeByAccountResponse, error) {
	response, err := e.exchangeService.GetTradeByAccount(ctx, in, opts...)
	if err != nil {
		return nil, err
	}
	return response, nil
}

func NewExchangeUseCase(exchangeService adapters.ExchangeService) adapters.ExchangeService {
	return &ExchangeUseCase{
		exchangeService: exchangeService,
	}
}

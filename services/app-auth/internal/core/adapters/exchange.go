package adapters

import (
	"context"
	"github.com/AppsLab-KE/be-go-gen-grpc/exchange"
	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/grpc"
)

type ExchangeStorage interface {
	// RATES
	CreateConversionRate(ctx context.Context, in *exchange.CreateConversionRateRequest, opts ...grpc.CallOption) (*exchange.CreateConversionRateResponse, error)
	ReadConversionRate(ctx context.Context, in *exchange.ReadConversionRateRequest, opts ...grpc.CallOption) (*exchange.ReadConversionRateResponse, error)
	UpdateConversionRate(ctx context.Context, in *exchange.UpdateConversionRateRequest, opts ...grpc.CallOption) (*exchange.UpdateConversionRateResponse, error)
	DeleteConversionRate(ctx context.Context, in *exchange.DeleteConversionRateRequest, opts ...grpc.CallOption) (*exchange.DeleteConversionRateResponse, error)
	// ACCOUNT
	CreateAccount(ctx context.Context, in *exchange.CreateAccountRequest, opts ...grpc.CallOption) (*exchange.CreateAccountResponse, error)
	DeleteAccount(ctx context.Context, in *exchange.DeleteAccountRequest, opts ...grpc.CallOption) (*exchange.DeleteAccountResponse, error)
	UpdateAccount(ctx context.Context, in *exchange.UpdateAccountRequest, opts ...grpc.CallOption) (*exchange.UpdateAccountResponse, error)
	SearchAccount(ctx context.Context, in *exchange.SearchAccountRequest, opts ...grpc.CallOption) (*exchange.SearchAccountResponse, error)
	// TRANSACTION
	CreateTransaction(ctx context.Context, in *exchange.CreateTransactionRequest, opts ...grpc.CallOption) (*exchange.CreateTransactionResponse, error)
	DeleteTransaction(ctx context.Context, in *exchange.DeleteTransactionRequest, opts ...grpc.CallOption) (*empty.Empty, error)
	UpdateTransaction(ctx context.Context, in *exchange.UpdateTransactionRequest, opts ...grpc.CallOption) (*empty.Empty, error)
	GetTransaction(ctx context.Context, in *exchange.GetTransactionRequest, opts ...grpc.CallOption) (*exchange.Transaction, error)
	GetTransactionByAccount(ctx context.Context, in *exchange.GetTransactionByAccountRequest, opts ...grpc.CallOption) (*exchange.GetTransactionByAccountResponse, error)
	// TRADING
	CreateTrade(ctx context.Context, in *exchange.CreateTradeRequest, opts ...grpc.CallOption) (*exchange.CreateTradeResponse, error)
	DeleteTrade(ctx context.Context, in *exchange.DeleteTradeRequest, opts ...grpc.CallOption) (*empty.Empty, error)
	UpdateTrade(ctx context.Context, in *exchange.UpdateTradeRequest, opts ...grpc.CallOption) (*empty.Empty, error)
	GetTrade(ctx context.Context, in *exchange.GetTradeRequest, opts ...grpc.CallOption) (*exchange.Trade, error)
	GetTradeByAccount(ctx context.Context, in *exchange.GetTradeByAccountRequest, opts ...grpc.CallOption) (*exchange.GetTradeByAccountResponse, error)
}

type ExchangeRepository interface {

	// RATES
	CreateConversionRate(ctx context.Context, in *exchange.CreateConversionRateRequest, opts ...grpc.CallOption) (*exchange.CreateConversionRateResponse, error)
	ReadConversionRate(ctx context.Context, in *exchange.ReadConversionRateRequest, opts ...grpc.CallOption) (*exchange.ReadConversionRateResponse, error)
	UpdateConversionRate(ctx context.Context, in *exchange.UpdateConversionRateRequest, opts ...grpc.CallOption) (*exchange.UpdateConversionRateResponse, error)
	DeleteConversionRate(ctx context.Context, in *exchange.DeleteConversionRateRequest, opts ...grpc.CallOption) (*exchange.DeleteConversionRateResponse, error)
	// ACCOUNT
	CreateAccount(ctx context.Context, in *exchange.CreateAccountRequest, opts ...grpc.CallOption) (*exchange.CreateAccountResponse, error)
	DeleteAccount(ctx context.Context, in *exchange.DeleteAccountRequest, opts ...grpc.CallOption) (*exchange.DeleteAccountResponse, error)
	UpdateAccount(ctx context.Context, in *exchange.UpdateAccountRequest, opts ...grpc.CallOption) (*exchange.UpdateAccountResponse, error)
	SearchAccount(ctx context.Context, in *exchange.SearchAccountRequest, opts ...grpc.CallOption) (*exchange.SearchAccountResponse, error)
	// TRANSACTION
	CreateTransaction(ctx context.Context, in *exchange.CreateTransactionRequest, opts ...grpc.CallOption) (*exchange.CreateTransactionResponse, error)
	DeleteTransaction(ctx context.Context, in *exchange.DeleteTransactionRequest, opts ...grpc.CallOption) (*empty.Empty, error)
	UpdateTransaction(ctx context.Context, in *exchange.UpdateTransactionRequest, opts ...grpc.CallOption) (*empty.Empty, error)
	GetTransaction(ctx context.Context, in *exchange.GetTransactionRequest, opts ...grpc.CallOption) (*exchange.Transaction, error)
	GetTransactionByAccount(ctx context.Context, in *exchange.GetTransactionByAccountRequest, opts ...grpc.CallOption) (*exchange.GetTransactionByAccountResponse, error)
	// TRADING
	CreateTrade(ctx context.Context, in *exchange.CreateTradeRequest, opts ...grpc.CallOption) (*exchange.CreateTradeResponse, error)
	DeleteTrade(ctx context.Context, in *exchange.DeleteTradeRequest, opts ...grpc.CallOption) (*empty.Empty, error)
	UpdateTrade(ctx context.Context, in *exchange.UpdateTradeRequest, opts ...grpc.CallOption) (*empty.Empty, error)
	GetTrade(ctx context.Context, in *exchange.GetTradeRequest, opts ...grpc.CallOption) (*exchange.Trade, error)
	GetTradeByAccount(ctx context.Context, in *exchange.GetTradeByAccountRequest, opts ...grpc.CallOption) (*exchange.GetTradeByAccountResponse, error)
}

type ExchangeService interface {

	// RATES
	CreateConversionRate(ctx context.Context, in *exchange.CreateConversionRateRequest, opts ...grpc.CallOption) (*exchange.CreateConversionRateResponse, error)
	ReadConversionRate(ctx context.Context, in *exchange.ReadConversionRateRequest, opts ...grpc.CallOption) (*exchange.ReadConversionRateResponse, error)
	UpdateConversionRate(ctx context.Context, in *exchange.UpdateConversionRateRequest, opts ...grpc.CallOption) (*exchange.UpdateConversionRateResponse, error)
	DeleteConversionRate(ctx context.Context, in *exchange.DeleteConversionRateRequest, opts ...grpc.CallOption) (*exchange.DeleteConversionRateResponse, error)
	// ACCOUNT
	CreateAccount(ctx context.Context, in *exchange.CreateAccountRequest, opts ...grpc.CallOption) (*exchange.CreateAccountResponse, error)
	DeleteAccount(ctx context.Context, in *exchange.DeleteAccountRequest, opts ...grpc.CallOption) (*exchange.DeleteAccountResponse, error)
	UpdateAccount(ctx context.Context, in *exchange.UpdateAccountRequest, opts ...grpc.CallOption) (*exchange.UpdateAccountResponse, error)
	SearchAccount(ctx context.Context, in *exchange.SearchAccountRequest, opts ...grpc.CallOption) (*exchange.SearchAccountResponse, error)
	// TRANSACTION
	CreateTransaction(ctx context.Context, in *exchange.CreateTransactionRequest, opts ...grpc.CallOption) (*exchange.CreateTransactionResponse, error)
	DeleteTransaction(ctx context.Context, in *exchange.DeleteTransactionRequest, opts ...grpc.CallOption) (*empty.Empty, error)
	UpdateTransaction(ctx context.Context, in *exchange.UpdateTransactionRequest, opts ...grpc.CallOption) (*empty.Empty, error)
	GetTransaction(ctx context.Context, in *exchange.GetTransactionRequest, opts ...grpc.CallOption) (*exchange.Transaction, error)
	GetTransactionByAccount(ctx context.Context, in *exchange.GetTransactionByAccountRequest, opts ...grpc.CallOption) (*exchange.GetTransactionByAccountResponse, error)
	// TRADING
	CreateTrade(ctx context.Context, in *exchange.CreateTradeRequest, opts ...grpc.CallOption) (*exchange.CreateTradeResponse, error)
	DeleteTrade(ctx context.Context, in *exchange.DeleteTradeRequest, opts ...grpc.CallOption) (*empty.Empty, error)
	UpdateTrade(ctx context.Context, in *exchange.UpdateTradeRequest, opts ...grpc.CallOption) (*empty.Empty, error)
	GetTrade(ctx context.Context, in *exchange.GetTradeRequest, opts ...grpc.CallOption) (*exchange.Trade, error)
	GetTradeByAccount(ctx context.Context, in *exchange.GetTradeByAccountRequest, opts ...grpc.CallOption) (*exchange.GetTradeByAccountResponse, error)
}

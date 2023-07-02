package ports

import (
	"context"
	"github.com/AppsLab-KE/be-go-gen-grpc/db"
	"google.golang.org/protobuf/types/known/emptypb"
)

type DBStorage interface {
	HealthCheck(ctx context.Context, in *db.DefaultRequest) (*db.HealthResponse, error)
	// RATES
	ReadConversionRate(ctx context.Context, in *db.ReadConversionRateRequest) (*db.ReadConversionRateResponse, error)
	// ACCOUNT
	CreateAccount(ctx context.Context, in *db.CreateAccountRequest) (*db.CreateAccountResponse, error)
	DeleteAccount(ctx context.Context, in *db.DeleteAccountRequest) (*db.DeleteAccountResponse, error)
	UpdateAccount(ctx context.Context, in *db.UpdateAccountRequest) (*db.UpdateAccountResponse, error)
	SearchAccount(ctx context.Context, in *db.SearchAccountRequest) (*db.SearchAccountResponse, error)
	// TRANSACTION
	CreateTransaction(ctx context.Context, in *db.CreateTransactionRequest) (*db.CreateTransactionResponse, error)
	DeleteTransaction(ctx context.Context, in *db.DeleteTransactionRequest) (emptypb.Empty, error)
	UpdateTransaction(ctx context.Context, in *db.UpdateTransactionRequest) (emptypb.Empty, error)
	GetTransaction(ctx context.Context, in *db.GetTransactionRequest) (*db.Transaction, error)
	GetTransactionByAccount(ctx context.Context, in *db.GetTransactionByAccountRequest) (*db.GetTransactionByAccountResponse, error)
	// TRADING
	CreateTrade(ctx context.Context, in *db.CreateTradeRequest) (*db.CreateTradeResponse, error)
	DeleteTrade(ctx context.Context, in *db.DeleteTradeRequest) (emptypb.Empty, error)
	UpdateTrade(ctx context.Context, in *db.UpdateTradeRequest) (emptypb.Empty, error)
	GetTrade(ctx context.Context, in *db.GetTradeRequest) (*db.Trade, error)
	GetTradeByAccount(ctx context.Context, in *db.GetTradeByAccountRequest) (*db.GetTradeByAccountResponse, error)
}

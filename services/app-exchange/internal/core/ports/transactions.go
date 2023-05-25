package ports

import (
	"context"
	"github.com/AppsLab-KE/be-go-gen-grpc/db"
	"google.golang.org/protobuf/types/known/emptypb"
)

type TransactionRepository interface {
	CreateTransaction(ctx context.Context, in *db.CreateTransactionRequest) (*db.CreateTransactionResponse, error)
	DeleteTransaction(ctx context.Context, in *db.DeleteTransactionRequest) (emptypb.Empty, error)
	UpdateTransaction(ctx context.Context, in *db.UpdateTransactionRequest) (emptypb.Empty, error)
	GetTransaction(ctx context.Context, in *db.GetTransactionRequest) (*db.Transaction, error)
	GetTransactionByAccount(ctx context.Context, in *db.GetTransactionByAccountRequest) (*db.GetTransactionByAccountResponse, error)
}

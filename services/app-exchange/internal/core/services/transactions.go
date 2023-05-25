package services

import (
	"context"
	"errors"
	"github.com/AppsLab-KE/backend-everyshilling/services/app-exchange/internal/core/ports"
	"github.com/AppsLab-KE/be-go-gen-grpc/db"
	"google.golang.org/protobuf/types/known/emptypb"
)

var (
	InvalidTransactionTypeErr = errors.New("invalid transaction type")
	NonExistingAccountErr     = errors.New("account does not exist")
	NonExistingTransactionErr = errors.New("transaction does not exist")
	InsufficientFundErr       = errors.New("insufficient funds")
)

const (
	TransactionDebit  = "debit"
	TransactionCredit = "credit"
)

type transactionService struct {
	transactionRepo ports.TransactionRepository
	accountsRepo    ports.AccountRepository
}

func (t transactionService) CreateTransaction(ctx context.Context, in *db.CreateTransactionRequest) (*db.CreateTransactionResponse, error) {
	// validate transaction type
	if in.Transaction.TransactionType != TransactionDebit && in.Transaction.TransactionType != TransactionCredit {
		return nil, InvalidTransactionTypeErr
	}

	// validate account exists
	account, err := t.accountsRepo.SearchAccount(ctx, &db.SearchAccountRequest{AccountId: in.Transaction.AccountId})
	if err != nil {
		return nil, NonExistingAccountErr
	}

	// get the account matching the account id
	var accountLookedUp *db.Account
	for _, acc := range account.Accounts {
		if acc.AccountId == in.Transaction.AccountId {
			accountLookedUp = acc
		}
	}

	if accountLookedUp == nil {
		return nil, NonExistingAccountErr
	}

	// check balance
	if in.Transaction.TransactionType == TransactionDebit && accountLookedUp.Balance < in.Transaction.Amount {
		return nil, InsufficientFundErr
	}

	res, err := t.transactionRepo.CreateTransaction(ctx, in)
	if err != nil {
		return nil, err
	}

	// update account balance
	if in.Transaction.TransactionType == TransactionDebit {
		accountLookedUp.Balance -= in.Transaction.Amount
	} else {
		accountLookedUp.Balance += in.Transaction.Amount
	}

	_, err = t.accountsRepo.UpdateAccount(ctx, &db.UpdateAccountRequest{Account: accountLookedUp})
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (t transactionService) DeleteTransaction(ctx context.Context, in *db.DeleteTransactionRequest) (emptypb.Empty, error) {
	// check if transaction exists
	_, err := t.transactionRepo.GetTransaction(ctx, &db.GetTransactionRequest{TransactionId: in.TransactionId})
	if err != nil {
		return emptypb.Empty{}, NonExistingTransactionErr
	}

	// delete transaction
	_, err = t.transactionRepo.DeleteTransaction(ctx, in)
	if err != nil {
		return emptypb.Empty{}, err
	}

	return emptypb.Empty{}, nil
}

func (t transactionService) UpdateTransaction(ctx context.Context, in *db.UpdateTransactionRequest) (emptypb.Empty, error) {
	_, err := t.transactionRepo.GetTransaction(ctx, &db.GetTransactionRequest{TransactionId: in.Transaction.TransactionId})
	if err != nil {
		return emptypb.Empty{}, NonExistingTransactionErr
	}

	_, err = t.transactionRepo.UpdateTransaction(ctx, in)
	if err != nil {
		return emptypb.Empty{}, err
	}

	return emptypb.Empty{}, nil
}

func (t transactionService) GetTransaction(ctx context.Context, in *db.GetTransactionRequest) (*db.Transaction, error) {
	transactions, err := t.transactionRepo.GetTransaction(ctx, in)
	if err != nil {
		return nil, err
	}

	return transactions, nil
}

func (t transactionService) GetTransactionByAccount(ctx context.Context, in *db.GetTransactionByAccountRequest) (*db.GetTransactionByAccountResponse, error) {
	transactions, err := t.transactionRepo.GetTransactionByAccount(ctx, in)
	if err != nil {
		return nil, err
	}

	return transactions, nil
}

func NewTransactionService(repository ports.TransactionRepository) ports.TransactionService {
	return &transactionService{
		transactionRepo: repository,
	}
}

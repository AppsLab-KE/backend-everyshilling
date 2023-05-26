package services

import (
	"context"
	"errors"
	"github.com/AppsLab-KE/backend-everyshilling/services/app-exchange/internal/core/ports"
	"github.com/AppsLab-KE/be-go-gen-grpc/db"
)

var (
	AccountExistsErr         = errors.New("account already exists")
	ErrParentAccountRequired = errors.New("parent account is required")
)

type accountsService struct {
	accountRepository ports.AccountRepository
}

func (a accountsService) CreateAccount(ctx context.Context, in *db.CreateAccountRequest) (*db.CreateAccountResponse, error) {
	// list all the account that the user has
	accounts, err := a.accountRepository.SearchAccount(ctx, &db.SearchAccountRequest{UserId: in.Account.UserId})
	if err != nil {
		return nil, err
	}

	// check if the account already exists
	for _, account := range accounts.Accounts {
		if account.BaseCurrency == in.Account.BaseCurrency {
			return nil, AccountExistsErr
		}
	}

	// if an account exists, make sure parent account is specified
	if len(accounts.Accounts) > 0 && in.Account.ParentAccountId == "" {
		return nil, ErrParentAccountRequired
	}

	// check if the parent account exists
	_, err = a.accountRepository.SearchAccount(ctx, &db.SearchAccountRequest{AccountId: in.Account.ParentAccountId})
	if err != nil {
		return nil, NonExistingAccountErr
	}

	// go ahead and create the account
	return a.accountRepository.CreateAccount(ctx, in)

}

func (a accountsService) DeleteAccount(ctx context.Context, in *db.DeleteAccountRequest) (*db.DeleteAccountResponse, error) {
	// just delete the account
	return a.accountRepository.DeleteAccount(ctx, in)
}

func (a accountsService) UpdateAccount(ctx context.Context, in *db.UpdateAccountRequest) (*db.UpdateAccountResponse, error) {
	// just update the account
	return a.accountRepository.UpdateAccount(ctx, in)
}

func (a accountsService) SearchAccount(ctx context.Context, in *db.SearchAccountRequest) (*db.SearchAccountResponse, error) {
	// just search the account
	return a.accountRepository.SearchAccount(ctx, in)
}

func NewAccountsService(accountRepository ports.AccountRepository) ports.AccountsService {
	return &accountsService{accountRepository: accountRepository}
}

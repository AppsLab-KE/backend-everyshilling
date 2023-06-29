package handlers

import (
	"context"
	"github.com/AppsLab-KE/be-go-gen-grpc/db"
	"github.com/AppsLab-KE/be-go-gen-grpc/exchange"
	"time"
)

func (h Handler) CreateAccount(ctx context.Context, request *exchange.CreateAccountRequest) (*exchange.CreateAccountResponse, error) {
	accountsRequest := &db.CreateAccountRequest{
		Account: &db.Account{
			UserId:          request.Account.UserId,
			BaseCurrency:    request.Account.BaseCurrency,
			CreatedAt:       time.Now().UnixNano(),
			ParentAccountId: request.Account.ParentAccountId,
		},
	}
	res, err := h.accountsService.CreateAccount(ctx, accountsRequest)
	if err != nil {
		return nil, err
	}

	accountsResult := &exchange.CreateAccountResponse{
		Account: &exchange.Account{
			AccountId:       res.Account.AccountId,
			UserId:          res.Account.UserId,
			Balance:         res.Account.Balance,
			BaseCurrency:    res.Account.BaseCurrency,
			CreatedAt:       res.Account.CreatedAt,
			ParentAccountId: res.Account.ParentAccountId,
		},
	}

	return accountsResult, nil
}

func (h Handler) DeleteAccount(ctx context.Context, request *exchange.DeleteAccountRequest) (*exchange.DeleteAccountResponse, error) {
	deleteRequest := &db.DeleteAccountRequest{
		AccountId: request.AccountId,
	}

	_, err := h.accountsService.DeleteAccount(ctx, deleteRequest)
	if err != nil {
		return nil, err
	}

	deleteResult := &exchange.DeleteAccountResponse{}

	return deleteResult, nil

}

func (h Handler) UpdateAccount(ctx context.Context, request *exchange.UpdateAccountRequest) (*exchange.UpdateAccountResponse, error) {
	updateRequest := &db.UpdateAccountRequest{
		Account: &db.Account{
			AccountId:       request.Account.AccountId,
			UserId:          request.Account.UserId,
			Balance:         request.Account.Balance,
			BaseCurrency:    request.Account.BaseCurrency,
			CreatedAt:       request.Account.CreatedAt,
			ParentAccountId: request.Account.ParentAccountId,
		},
	}

	_, err := h.accountsService.UpdateAccount(ctx, updateRequest)
	if err != nil {
		return nil, err
	}

	updateResult := &exchange.UpdateAccountResponse{}

	return updateResult, nil
}

func (h Handler) SearchAccount(ctx context.Context, request *exchange.SearchAccountRequest) (*exchange.SearchAccountResponse, error) {
	searchRequest := &db.SearchAccountRequest{
		AccountId: request.AccountId,
	}

	res, err := h.accountsService.SearchAccount(ctx, searchRequest)
	if err != nil {
		return nil, err
	}

	searchResult := &exchange.SearchAccountResponse{
		Accounts: []*exchange.Account{},
	}

	for _, account := range res.Accounts {
		searchResult.Accounts = append(searchResult.Accounts, &exchange.Account{
			AccountId:       account.AccountId,
			UserId:          account.UserId,
			Balance:         account.Balance,
			BaseCurrency:    account.BaseCurrency,
			CreatedAt:       account.CreatedAt,
			ParentAccountId: account.ParentAccountId,
		})
	}

	return searchResult, nil
}

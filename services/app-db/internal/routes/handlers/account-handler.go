package handlers

import (
	"context"
	"fmt"
	"github.com/AppsLab-KE/backend-everyshilling/services/app-db/internal/core/models"
	"github.com/AppsLab-KE/be-go-gen-grpc/db"
)

func (h *Handler) CreateAccount(ctx context.Context, req *db.CreateAccountRequest) (*db.CreateAccountResponse, error) {
	if req == nil {
		return nil, ErrEmptyRequest
	}

	var accounts []*models.Account
	for _, account := range req.Account {
		accounts = append(accounts, &models.Account{
			AccountId:       account,
			UserId:          account.UserId,
			Balance:         account.Balance,
			BaseCurrency:    account.BaseCurrency,
			CreatedAt:       account.CreatedAt,
			ParentAccountId: account.ParentAccountId,
		})
	}

	createdAccounts, err := h.accountsRepo.CreateAccounts(ctx, accounts)
	if err != nil {
		return nil, err
	}

	resp := &db.CreateAccountResponse{}

	for _, account := range createdAccounts {
		resp.Account = append(resp.Account, &db.Account{
			AccountId:       account.AccountId,
			UserId:          account.UserId,
			Balance:         account.Balance,
			BaseCurrency:    account.BaseCurrency,
			CreatedAt:       account.CreatedAt,
			ParentAccountId: account.ParentAccountId,
		})
	}

	return resp, nil
}

func (h *Handler) DeleteAccount(ctx context.Context, req *db.DeleteAccountRequest) (*db.DeleteAccountResponse, error) {

	if req == nil {
		return nil, ErrEmptyRequest
	}

	// Extract the account ID from the request
	accountID := req.AccountId

	// Delete the account by account ID
	err := h.accountsRepo.DeleteAccountByID(ctx, accountID)
	if err != nil {
		return nil, err
	}

	// Create the response
	resp := &db.DeleteAccountResponse{
		Message: fmt.Sprintf("Account with ID %s deleted successfully", accountID),
	}

	return resp, nil
}

func (h *Handler) UpdateAccount(ctx context.Context, req *db.UpdateAccountRequest) (*db.UpdateAccountResponse, error) {

}

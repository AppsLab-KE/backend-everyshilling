package usecase

import (
	"context"

	"github.com/AppsLab-KE/backend-everyshilling/services/app-exchange/internal/core/ports"
)

type AccountUseCase struct {
	accountService ports.AccountService
}

func NewAccountUseCase(accountService ports.AccountService) ports.AccountUseCase {
	return &AccountUseCase{
		accountService: accountService,
	}
}

func (uc *AccountUseCase) GetAccountOverview(ctx context.Context) (interface{}, error) {
	// Call the account service to retrieve the account overview
	accountData, err := uc.accountService.GetAccountOverview(ctx)
	if err != nil {
		return nil, err
	}

	// Implement the logic to process the account data as needed

	return accountData, nil
}

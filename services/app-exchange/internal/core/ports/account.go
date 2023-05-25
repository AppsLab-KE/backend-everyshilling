package ports

import "context"

type AccountService interface {
	GetAccountOverview(ctx context.Context) (interface{}, error)
	// Add more methods related to account functionality
}

type AccountRepository interface {
	GetAccountOverview(ctx context.Context) (interface{}, error)
	// Add more methods related to account repository functionality
}

type AccountUseCase interface {
	GetAccountOverview(ctx context.Context) (interface{}, error)
	// Add more methods related to account use case functionality
}

type CurrencyService interface {
	GetAccountOverview(ctx context.Context) (interface{}, error)
	// Add more methods related to account functionality
}

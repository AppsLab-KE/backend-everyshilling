package service

import (
	"context"
	"github.com/AppsLab-KE/backend-everyshilling/services/app-exchange/internal/core/adapters"
	"github.com/AppsLab-KE/backend-everyshilling/services/app-exchange/internal/dto"
)

type CurrencyService struct {
	repo adapters.CurrencyRepo
}

// errors implementation
var ()

func (as *CurrencyService) GetAccountOverview(request dto.AccountOverviewResponse) (*dto.AccountOverviewResponseObj, error) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

}

func (as *CurrencyService) PostBuyCurrency(request dto.BuyRequestData) (*dto.BuyReqData, error) {
	ctx, cancel := context.WithCancel(context.Background())

	defer cancel()
	//implementations

}

func (as *CurrencyService) GetMarketplaceOffers(request dto.MarketplaceOffersResponse) (*dto.MarketplaceOffersResponseObj, error) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	//implementations

}

func (as *CurrencyService) PostTopUpAccount(request dto.TopUpRequest) (*dto.TopUpRequestData, error) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	//implementations

}

package adapters

import "github.com/AppsLab-KE/backend-everyshilling/services/app-exchange/internal/dto"

type SessionService interface {
	Invalidate() dto.GeneralResponse[interface{}]
}

type CurrencyService interface {
	GetAccountOverview(request dto.AccountOverviewResponse) (*dto.AccountOverviewResponseObj, error)
	PostBuyCurrency(request dto.BuyRequestData) (*dto.BuyReqData, error)
	GetMarketplaceOffers(request dto.MarketplaceOffersResponse) (*dto.MarketplaceOffersResponseObj, error)
	PostTopUpAccount(request dto.TopUpRequest) (*dto.TopUpRequestData, error)
}

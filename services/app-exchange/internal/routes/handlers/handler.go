package handlers

import (
	"github.com/AppsLab-KE/backend-everyshilling/services/app-exchange/internal/core/ports"
	"github.com/AppsLab-KE/be-go-gen-grpc/exchange"
	log "github.com/sirupsen/logrus"
)

type Handler struct {
	exchange.UnimplementedExchangeServiceServer
	exchangeService    ports.ExchangeService
	accountsService    ports.AccountsService
	tradeService       ports.TradeService
	transactionService ports.TransactionService
}

func (h Handler) mustEmbedUnimplementedExchangeServiceServer() {
	//TODO implement me
	log.Error("not implemented")
}

func NewHandler(exchangeService ports.ExchangeService, accountsService ports.AccountsService, tradeService ports.TradeService, transactionService ports.TransactionService) *Handler {
	return &Handler{
		exchangeService:    exchangeService,
		accountsService:    accountsService,
		tradeService:       tradeService,
		transactionService: transactionService,
	}
}

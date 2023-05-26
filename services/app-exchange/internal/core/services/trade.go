package services

import (
	"context"
	"errors"
	"github.com/AppsLab-KE/backend-everyshilling/services/app-exchange/internal/core/ports"
	"github.com/AppsLab-KE/be-go-gen-grpc/db"
	"google.golang.org/protobuf/types/known/emptypb"
	"time"
)

var (
	ErrSameCurrency        = errors.New("cannot trade same currency")
	ErrNonMatchingCurrency = errors.New("currency does not match account currency")
	ErrDestinationAccount  = errors.New("destination account does not have the currency")
)

type tradeService struct {
	tradeRepository       ports.TradeRepository
	accountRepository     ports.AccountRepository
	rateRepository        ports.ExchangeRepository
	transactionRepository ports.TransactionRepository
}

func (t tradeService) CreateTrade(ctx context.Context, in *db.CreateTradeRequest) (*db.CreateTradeResponse, error) {
	// check that currencies are not the same
	if in.Trade.FromCurrency == in.Trade.ToCurrency {
		return nil, ErrSameCurrency
	}

	// check that the account has enough balance
	account, err := t.accountRepository.SearchAccount(ctx, &db.SearchAccountRequest{AccountId: in.Trade.AccountId})
	if err != nil {
		return nil, NonExistingAccountErr
	}

	// fetch the account that is being traded with
	var tradeAccount *db.Account
	var destinationAccount *db.Account

	for _, acc := range account.Accounts {
		if acc.AccountId == in.Trade.AccountId {
			tradeAccount = acc
		}
	}

	if tradeAccount == nil {
		return nil, NonExistingAccountErr
	}

	// check if currency matched the account currency
	if tradeAccount.BaseCurrency != in.Trade.FromCurrency {
		return nil, ErrNonMatchingCurrency
	}

	if tradeAccount.Balance < in.Trade.FromAmount {
		return nil, InsufficientFundErr
	}

	// get a destination account that matches the to currency
	for _, acc := range account.Accounts {
		if acc.BaseCurrency == in.Trade.ToCurrency {
			destinationAccount = acc
		}
	}

	if destinationAccount == nil {
		return nil, ErrDestinationAccount
	}

	// check the rate
	rates, err := t.rateRepository.ReadConversionRate(ctx, &db.ReadConversionRateRequest{})
	if err != nil {
		return nil, err
	}

	// get the base currency
	var baseCurrency = in.Trade.FromCurrency

	var exchangeRate float64
	for _, rate := range rates.ConversionRate {
		if rate.FromCurrency == baseCurrency && rate.ToCurrency == in.Trade.FromCurrency {
			exchangeRate = rate.Rate
		}
	}

	// if base currency is not USD then convert money using USD as the base currency
	if baseCurrency != "USD" {
		for _, rate := range rates.ConversionRate {
			if rate.FromCurrency == "USD" && rate.ToCurrency == in.Trade.FromCurrency {
				exchangeRate = rate.Rate
			}
		}
	}

	// calculate the amount to be deducted from the account
	amountToBeDeducted := float64(in.Trade.FromAmount) * exchangeRate

	// deduct the amount from the account
	tradeAccount.Balance -= int64(in.Trade.FromAmount)

	// update the account
	_, err = t.accountRepository.UpdateAccount(ctx, &db.UpdateAccountRequest{Account: tradeAccount})
	if err != nil {
		return nil, err
	}

	// add the amount to the destination account
	destinationAccount.Balance += int64(amountToBeDeducted)
	// update the account
	_, err = t.accountRepository.UpdateAccount(ctx, &db.UpdateAccountRequest{Account: destinationAccount})
	if err != nil {
		return nil, err
	}

	// create debit and credit transactions
	// credit transaction
	creditTransaction := &db.Transaction{
		AccountId:              in.Trade.AccountId,
		Amount:                 in.Trade.FromAmount,
		TransactionType:        "credit",
		TransactionStatus:      "",
		TransactionCode:        "",
		TransactionDescription: "",
		CreatedAt:              time.Now().UnixNano(),
	}

	_, err = t.transactionRepository.CreateTransaction(ctx, &db.CreateTransactionRequest{Transaction: creditTransaction})
	if err != nil {
		return nil, err
	}

	// debitTransaction
	debitTransaction := &db.Transaction{
		AccountId:              destinationAccount.AccountId,
		Amount:                 int64(amountToBeDeducted),
		TransactionType:        "debit",
		TransactionStatus:      "",
		TransactionCode:        "",
		TransactionDescription: "",
		CreatedAt:              time.Now().UnixNano(),
	}

	_, err = t.transactionRepository.CreateTransaction(ctx, &db.CreateTransactionRequest{Transaction: debitTransaction})
	if err != nil {
		return nil, err
	}

	// create the trade
	trade, err := t.tradeRepository.CreateTrade(ctx, in)
	if err != nil {
		return nil, err
	}

	// return the trade
	return trade, nil

}

func (t tradeService) DeleteTrade(ctx context.Context, in *db.DeleteTradeRequest) (emptypb.Empty, error) {
	// just delete the trade without any checks
	_, err := t.tradeRepository.DeleteTrade(ctx, in)
	if err != nil {
		return emptypb.Empty{}, err
	}
	return emptypb.Empty{}, nil
}

func (t tradeService) UpdateTrade(ctx context.Context, in *db.UpdateTradeRequest) (emptypb.Empty, error) {
	// just update the trade without any checks
	_, err := t.tradeRepository.UpdateTrade(ctx, in)
	if err != nil {
		return emptypb.Empty{}, err
	}
	return emptypb.Empty{}, nil
}

func (t tradeService) GetTrade(ctx context.Context, in *db.GetTradeRequest) (*db.Trade, error) {
	// just get the trade without any checks
	trade, err := t.tradeRepository.GetTrade(ctx, in)
	if err != nil {
		return nil, err
	}
	return trade, nil
}

func (t tradeService) GetTradeByAccount(ctx context.Context, in *db.GetTradeByAccountRequest) (*db.GetTradeByAccountResponse, error) {
	// just get the trade without any checks
	trade, err := t.tradeRepository.GetTradeByAccount(ctx, in)
	if err != nil {
		return nil, err
	}
	return trade, nil
}

func NewTradeService(tradeRepository ports.TradeRepository) ports.TradeService {
	return &tradeService{
		tradeRepository: tradeRepository,
	}
}

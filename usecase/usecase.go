package usecase

import "clean-go/entities"

type UseCase interface {
	GetTransactions() (resp []entities.CoinIOGetTransactionsResponse, err error)
	GetExchanges() (resp []entities.CoinIOGetExchangesResponse, err error)
}

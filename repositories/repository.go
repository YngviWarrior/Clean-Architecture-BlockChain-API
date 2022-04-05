package repository

import (
	"clean-go/entities"
)

type Repository struct {
	TransactionRepo
	ExchangeRepo
}

type TransactionRepo interface {
	GetTransactions() ([]entities.Transaction, error)
	CreateTransaction() (entities.Transaction, error)
}

type ExchangeRepo interface {
	GetExchanges() ([]entities.Exchange, error)
	CreateExchanges([]entities.CoinIOGetExchangesResponse) error
}

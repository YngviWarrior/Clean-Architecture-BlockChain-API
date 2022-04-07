package repository

import (
	"clean-go/apis"
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
	GetExchangesByID(id string) (obj entities.Exchange, err error)
	CreateExchanges([]apis.GetExchangeResponse) error
}

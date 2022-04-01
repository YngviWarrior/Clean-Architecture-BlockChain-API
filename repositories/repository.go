package repository

import (
	"clean-go/entities"
)

type Repository interface {
	TransactionRepo
}

type TransactionRepo interface {
	GetTransactions() (entities.Transaction, error)
	CreateTransaction() (entities.Transaction, error)
}

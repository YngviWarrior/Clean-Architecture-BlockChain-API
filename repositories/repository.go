package repository

import "clean-go/entities"

type Repository interface {
	transactionRepository
}

type transactionRepository interface {
	GetTransactions() (entities.Transaction, error)
	CreateTransaction() (entities.Transaction, error)
}

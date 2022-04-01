package usecase

import (
	"clean-go/entities"
	repository "clean-go/repositories"
)

func GetTransactions() (t entities.Transaction, err error) {
	tx := repository.Transaction{}
	t, err = tx.GetTransactions()

	return t, err
}

package usecase

import (
	"clean-go/entities"
	repository "clean-go/repositories"
)

func GetTransactions() (resp entities.Transaction, err error) {
	tx := repository.Transaction{}
	resp, err = tx.GetTransactions()

	return resp, err
}

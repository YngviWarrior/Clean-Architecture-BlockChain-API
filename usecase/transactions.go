package usecase

import (
	"clean-go/entities"
	repository "clean-go/repositories"
	"fmt"
)

func GetTransactions() (t entities.Transaction, err error) {
	transactionRepo, err := repository.Repository.GetTransactions
	fmt.Println(transactionRepo, err)

	return t, err
}

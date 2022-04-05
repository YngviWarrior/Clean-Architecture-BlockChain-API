package usecase

import (
	"clean-go/entities"
	repository "clean-go/repositories"
)

func GetTransactions() (resp []entities.Transaction, err error) {
	var repoInterface repository.Repository
	resp, err = repoInterface.GetTransactions()

	return resp, err
}

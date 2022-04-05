package main

import (
	"clean-go/usecase"
)

func main() {
	// var repoInterface repository.Repository
	// res, err := repoInterface.GetExchanges()
	// fmt.Println(res, err)

	usecase.GetExchanges()
	// tx, err := usecase.GetTransactions()
}

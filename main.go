package main

import (
	"clean-go/usecase"
	"fmt"
)

func main() {
	exc, _ := usecase.GetExchanges()
	// tx, err := usecase.GetTransactions()

	for _, e := range exc {
		fmt.Printf("ID: %v -- Name: %v \n", e.ExchangeID, e.Name)
	}
}

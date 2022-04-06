package controllers

import (
	"clean-go/usecase"
	"fmt"
	"net/http"
)

type TransactionController struct{}

func (p TransactionController) GetTransaction(w http.ResponseWriter, r *http.Request) {
	_, err := usecase.GetTransactions()

	if err != nil {
		fmt.Println("Erro")
	}

	// return resp
}

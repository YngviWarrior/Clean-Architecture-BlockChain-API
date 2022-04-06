package main

import (
	"clean-go/controllers"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	// var repoInterface repository.Repository
	// res, err := repoInterface.GetExchanges()
	// fmt.Println(res, err)

	// usecase.GetExchanges()
	// tx, err := usecase.GetTransactions()

	var ExchangeController controllers.ExchangeController
	var TransactionController controllers.TransactionController

	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", controllers.HomeLink).Methods("GET")
	router.HandleFunc("/exchanges", ExchangeController.GetExchange).Methods("GET")
	router.HandleFunc("/transactions", TransactionController.GetTransaction).Methods("GET")
	log.Fatal(http.ListenAndServe(":8080", router))
}

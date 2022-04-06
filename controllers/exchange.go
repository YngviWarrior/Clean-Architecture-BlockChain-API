package controllers

import (
	"clean-go/usecase"
	"encoding/json"
	"fmt"
	"net/http"
)

type ExchangeController struct{}

func (p ExchangeController) GetExchange(w http.ResponseWriter, r *http.Request) {
	exchange, err := usecase.GetExchanges()

	if err != nil {
		fmt.Println(err)
	}

	err = json.NewEncoder(w).Encode(exchange)

	if err != nil {
		fmt.Println(err)
	}

}

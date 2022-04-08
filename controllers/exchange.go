package controllers

import (
	"clean-go/usecase"
	"encoding/json"
	"fmt"
	"net/http"
)

type ExchangeController struct{}

func (p ExchangeController) GetExchange(w http.ResponseWriter, r *http.Request) {
	var response JsonResponse
	exchange, err := usecase.GetExchanges()

	if err != nil {
		response.Success = false
		fmt.Println(err)
	}

	response.Success = true
	response.Data = exchange

	err = json.NewEncoder(w).Encode(response)

	if err != nil {
		fmt.Println(err)
	}

}

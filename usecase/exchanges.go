package usecase

import (
	"clean-go/apis"
	"clean-go/entities"
	repository "clean-go/repositories"
)

func GetExchanges() (resp []entities.CoinIOGetExchangesResponse, err error) {
	var apiInterface apis.Api
	exchanges, err := apiInterface.GetExchanges()

	for _, exchange := range exchanges {
		var obj entities.CoinIOGetExchangesResponse
		obj.ExchangeID = exchange.ExchangeID
		obj.Name = exchange.Name
		obj.Website = exchange.Website

		resp = append(resp, obj)

	}

	var repoInterface repository.Repository
	err = repoInterface.CreateExchanges(resp)

	if err != nil {
		return nil, err
	}

	return resp, err
}

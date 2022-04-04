package usecase

import (
	"clean-go/apis"
	"clean-go/entities"
)

func GetExchanges() (resp []entities.CoinIOGetExchangesResponse, err error) {
	apiExchanges := apis.Exchange{}
	exchanges, err := apiExchanges.GetExchanges()

	for _, exchange := range exchanges {
		var obj entities.CoinIOGetExchangesResponse
		obj.ExchangeID = exchange.ExchangeID
		obj.Name = exchange.Name
		resp = append(resp, obj)
	}

	// dbExchanges := repository.Exchange{}
	// err = dbExchanges.CreateExchanges(resp)

	if err != nil {
		return
	}

	return resp, err
}

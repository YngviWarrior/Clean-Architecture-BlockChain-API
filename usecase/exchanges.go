package usecase

import (
	"clean-go/apis"
	"clean-go/entities"
	repository "clean-go/repositories"
	"fmt"
)

func checkExists(ids []entities.Exchange, exchanges []apis.GetExchangeResponse) bool {
	for j := 0; j < len(exchanges); j++ {
		for i := 0; i < len(ids); i++ {
			if ids[i].ID == exchanges[j].ExchangeID {
				return true
			}
		}
	}

	return false
}

func GetExchanges() (resp []entities.CoinIOGetExchangesResponse, err error) {
	var apiInterface apis.Api

	var repoInterface repository.Repository
	var ids []string

	dbExchanges, err := repoInterface.GetExchanges()

	if err != nil {
		fmt.Println(err)
		return
	}

	for _, exchange := range dbExchanges {
		ids = append(ids, exchange.ID)
	}

	exchanges, err := apiInterface.GetExchanges()

	if err != nil {
		fmt.Println(err)
		return
	}

	var list []apis.GetExchangeResponse

	for _, exchange := range exchanges {
		var obj entities.CoinIOGetExchangesResponse
		obj.ExchangeID = exchange.ExchangeID
		obj.Name = exchange.Name
		obj.Website = exchange.Website
		resp = append(resp, obj)
	}

	for _, exchange := range resp {
		has := checkExists(dbExchanges, exchanges)

		if has == true {
			continue
		}

		var obj apis.GetExchangeResponse
		obj.ExchangeID = exchange.ExchangeID
		obj.Name = exchange.Name
		obj.Website = exchange.Website
		list = append(list, obj)
	}

	if len(list) > 0 {
		err = repoInterface.CreateExchanges(list)

		if err != nil {
			return nil, err
		}
	}

	return resp, err
}

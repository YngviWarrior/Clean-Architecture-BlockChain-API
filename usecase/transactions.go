package usecase

import (
	"clean-go/apis"
	"clean-go/entities"
	repository "clean-go/repositories"
)

func GetTransactions() (t entities.Transaction, err error) {
	tx := repository.Transaction{}
	t, err = tx.GetTransactions()

	return t, err
}

type getExchangesResponse struct {
	ExchangeID         string `json:"exchange_id"`
	Website            string `json:"website"`
	Name               string `json:"name"`
	DataStart          string `json:"data_start"`
	DataEnd            string `json:"data_end"`
	DataQuoteStart     string `json:"data_quote_start"`
	DataQuoteEnd       string `json:"data_quote_end"`
	DataOrderBookStart string `json:"data_orderbook_start"`
	DataOrderBookEnd   string `json:"data_orderbook_end"`
	DataTradeStart     string `json:"data_trade_start"`
	DataTradeEnd       string `json:"data_trade_end"`
}

func GetExchanges() (resp []getExchangesResponse, err error) {
	apiExchanges := apis.Transaction{}
	exchanges, err := apiExchanges.FindExchanges()

	for _, exchange := range exchanges {
		var obj getExchangesResponse
		obj.ExchangeID = exchange.ExchangeID
		obj.Name = exchange.Name
		resp = append(resp, obj)
	}

	dbExchanges := repository.Transaction{}
	err = dbExchanges.CreateExchanges(resp)

	if err != nil {
		return
	}

	return resp, err
}

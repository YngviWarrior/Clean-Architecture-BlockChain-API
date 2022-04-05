package apis

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type GetTransactionResponse struct {
	ExchangeID string `json:"exchange_id"`
}

type GetExchangeResponse struct {
	ExchangeID         string  `json:"exchange_id"`
	Website            string  `json:"website"`
	Name               string  `json:"name"`
	DataStart          string  `json:"data_start"`
	DataEnd            string  `json:"data_end"`
	DataQuoteStart     string  `json:"data_quote_start"`
	DataQuoteEnd       string  `json:"data_quote_end"`
	DataOrderBookStart string  `json:"data_orderbook_start"`
	DataOrderBookEnd   string  `json:"data_orderbook_end"`
	DataTradeStart     string  `json:"data_trade_start"`
	DataTradeEnd       string  `json:"data_trade_end"`
	DataSymbolsCount   int64   `json:"data_symbols_count"`
	Volume1HrsUsd      float64 `json:"volume_1hrs_usd"`
	Volume1DayUsd      float64 `json:"volume_1day_usd"`
	Volume1MthUsd      float64 `json:"volume_1mth_usd"`
}

func (Api) FindTransaction() (response []GetTransactionResponse, err error) {
	return
}

func (Api) GetExchanges() (response []GetExchangeResponse, err error) {
	client := &http.Client{}
	req, err := http.NewRequest("GET", "http://rest-sandbox.coinapi.io/v1/exchanges", nil)

	if err != nil {
		fmt.Print(err.Error())
	}

	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	// req.Header.Add("X-CoinAPI-Key", "540A13CE-0504-41A7-9D95-0494DC440E37")

	q := req.URL.Query()
	q.Add("apikey", "540A13CE-0504-41A7-9D95-0494DC440E37")
	req.URL.RawQuery = q.Encode()

	resp, err := client.Do(req)

	bodyBytes, err := ioutil.ReadAll(resp.Body)

	json.Unmarshal(bodyBytes, &response)
	defer resp.Body.Close()

	if err != nil {
		fmt.Print(err.Error())
	}

	return response, err
}

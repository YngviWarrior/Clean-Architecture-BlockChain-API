package entities

type Exchange struct {
	ExchangeID string `json:"exchange_id"`
	Website    string `json:"website"`
	Name       string `json:"name"`
}

type CoinIOGetExchangesResponse struct {
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

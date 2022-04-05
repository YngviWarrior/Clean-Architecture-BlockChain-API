package apis

import "clean-go/entities"

type Api struct {
	CoinIO
}

type CoinIO interface {
	GetTransaction() ([]entities.CoinIOGetTransactionsResponse, error)
	GetExchanges() ([]entities.CoinIOGetExchangesResponse, error)
}

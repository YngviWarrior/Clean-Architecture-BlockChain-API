package apis

import "clean-go/entities"

type Apis interface {
	CoinIO
}

type CoinIO interface {
	GetTransaction() ([]entities.CoinIOGetTransactionsResponse, error)
	GetExchanges() ([]entities.CoinIOGetExchangesResponse, error)
}

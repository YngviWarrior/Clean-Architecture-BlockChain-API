package apis

type Apis interface {
	CoinIO
}

type CoinIO interface {
	FindTransaction() (interface{}, error)
	FindExchanges() (interface{}, error)
}

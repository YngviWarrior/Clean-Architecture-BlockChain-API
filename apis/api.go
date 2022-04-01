package apis

type Apis interface {
	CoinIO
}

type CoinIO interface {
	findTransaction()
}

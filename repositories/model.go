package repository

type Transaction struct {
	Transaction int64  `json:"transaction"`
	Txid        string `json:"txid"`
	Nonce       string `json:"nonce"`
}

type Exchange struct {
	ExchangeID int64  `json:"transaction"`
	Name       string `json:"txid"`
}

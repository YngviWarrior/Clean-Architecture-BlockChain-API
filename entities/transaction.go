package entities

type Transaction struct {
	Transaction int64  `json:"transaction"`
	Txid        string `json:"txid"`
	Nonce       string `json:"nonce"`
}

func NewTransaction() *Transaction {
	t := &Transaction{}
	return t
}

package repositories

import (
	"fmt"
)

type transaction struct {
	txid  string
	nonce string
}

func (r transaction) GetTransactions() {
	db, err := conn()
	fmt.Println(db, err)
}

func (r transaction) CreateTransaction() {

}

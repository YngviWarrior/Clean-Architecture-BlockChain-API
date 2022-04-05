package repository

import (
	"clean-go/entities"
	"database/sql"
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func exec(db *sql.DB, sql string) sql.Result {
	result, err := db.Exec(sql)
	if err != nil {
		panic(err)
	}

	return result
}

func conn() (db *sql.DB, err error) {
	db, err = sql.Open("mysql", "igor:123456@/go")

	if err != nil {
		panic(err)
	}

	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)

	return
}

/* Transactions */
func (Transaction) GetTransactions() (list []entities.Transaction, err error) {
	db, err := conn()

	if err != nil {
		return
	}

	defer db.Close()

	res, err := db.Query("SELECT * FROM transaction")

	defer res.Close()

	if err != nil {
		return
	}

	for res.Next() {
		obj := entities.Transaction{}
		err = res.Scan(&obj.Transaction, &obj.Txid, &obj.Nonce)

		if err != nil {
			return
		}

		list = append(list, obj)
	}

	return
}

func (Transaction) CreateTransaction() (obj entities.Transaction, err error) {
	return
}

func (Exchange) GetExchanges() (list []entities.Exchange, err error) {
	db, err := conn()

	if err != nil {
		return
	}

	defer db.Close()

	res, err := db.Query("SELECT * FROM transaction")

	defer res.Close()

	if err != nil {
		return
	}

	for res.Next() {
		obj := entities.Exchange{}
		err = res.Scan(&obj.ID, &obj.Name, &obj.Website)

		if err != nil {
			return
		}

		list = append(list, obj)
	}

	return
}

func (Exchange) CreateExchanges(exchanges []entities.CoinIOGetExchangesResponse) (err error) {
	db, err := conn()

	if err != nil {
		return
	}

	sql := "INSERT INTO exchange VALUES "

	for i, exchange := range exchanges {
		res, _ := db.Query("SELECT * FROM exchange WHERE exchange = ?", exchange.ExchangeID)

		if res.Next() {
			fmt.Printf("Already Exists %v", exchange.Name)
		} else {
			if i == (len(exchanges) - 1) {
				sql += fmt.Sprintf("('%v', '%v', '%v')", exchange.ExchangeID, exchange.Name, exchange.Website)
			} else {
				sql += fmt.Sprintf("('%v', '%v', '%v'), ", exchange.ExchangeID, exchange.Name, exchange.Website)
			}
		}

		defer res.Close()
	}

	res, err := db.Query(sql)

	if err != nil {
		return
	}

	defer res.Close()

	return
}

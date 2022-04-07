package repository

import (
	"clean-go/apis"
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
func (Repository) GetTransactions() (list []entities.Transaction, err error) {
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

func (Repository) CreateTransaction() (obj entities.Transaction, err error) {
	return
}

func (Repository) GetExchanges() (list []entities.Exchange, err error) {
	db, err := conn()

	if err != nil {
		return
	}

	defer db.Close()

	res, err := db.Query("SELECT * FROM exchange")

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

func (Repository) GetExchangeByID(id string) (obj entities.Exchange, err error) {
	db, err := conn()

	if err != nil {
		return
	}

	defer db.Close()

	res, err := db.Query("SELECT * FROM exchange WHERE exchange = ?", id)

	defer res.Close()

	if err != nil {
		return
	}

	for res.Next() {
		err = res.Scan(&obj.ID, &obj.Name, &obj.Website)

		if err != nil {
			return
		}
	}

	return
}

func (Repository) GetExchangesByID(ids []string) (list []entities.Exchange, err error) {
	db, err := conn()

	if err != nil {
		return
	}

	defer db.Close()

	sql := "SELECT * FROM exchange WHERE exchange IN ("

	for i, id := range ids {
		res, erro := db.Query("SELECT * FROM exchange WHERE exchange = ?", id)

		defer res.Close()

		if erro != nil {
			fmt.Println(erro)
		}

		if res.Next() {

		} else {
			if i == (len(id) - 1) {
				sql += fmt.Sprintf("'%v', ", id)
			} else {
				sql += fmt.Sprintf("'%v')", id)
			}
		}
	}

	res, err := db.Query(sql)

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

func (Repository) CreateExchanges(exchanges []apis.GetExchangeResponse) (err error) {
	db, err := conn()

	if err != nil {
		return
	}

	sql := "INSERT INTO exchange VALUES "
	sqlValues := ""

	for i, exchange := range exchanges {
		if i == (len(exchanges) - 1) {
			sqlValues += fmt.Sprintf("('%v', '%v', '%v')", exchange.ExchangeID, exchange.Name, exchange.Website)
		} else {
			sqlValues += fmt.Sprintf("('%v', '%v', '%v'), ", exchange.ExchangeID, exchange.Name, exchange.Website)
		}
	}

	if sqlValues != "" {
		sql = fmt.Sprintf("%s %s", sql, sqlValues)
		res, erro := db.Query(sql)

		defer res.Close()

		if erro != nil {
			fmt.Println(erro)
		}
	}

	return
}

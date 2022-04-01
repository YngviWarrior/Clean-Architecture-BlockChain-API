package repository

import (
	"clean-go/entities"
	"database/sql"
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
func (Transaction) GetTransactions() (obj entities.Transaction, err error) {
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
		err = res.Scan(&obj.Transaction, &obj.Txid, &obj.Nonce)

		if err != nil {
			return
		}
	}

	return
}

func CreateTransaction() (obj entities.Transaction, err error) {
	return
}

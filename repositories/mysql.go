package repositories

import (
	"database/sql"
)

func exec(db *sql.DB, sql string) sql.Result {
	result, err := db.Exec(sql)
	if err != nil {
		panic(err)
	}

	return result
}

func conn() (db *sql.DB, err error) {
	db, err = sql.Open("mysql", "igor:123456@go")
	if err != nil {
		panic(err)
	}

	defer db.Close()

	return
}

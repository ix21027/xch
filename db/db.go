package db

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

func Init() {
	var err error

	if DB, err = sql.Open("sqlite3", "xch.db"); err != nil {
		fmt.Println(err)
		return
	}

	createSubscribers()
}

func CloseConn() {
	DB.Close()
}
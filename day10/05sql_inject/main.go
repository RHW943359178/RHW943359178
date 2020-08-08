package main

import (
	"database/sql"
	"fmt"
	_ "src/github.com/go-sql-driver/mysql"
	"src/github.com/jmoiron/sqlx"
)

var db *sqlx.DB

func initDB() (err error) {
	dsn := "root:123456@tcp(127.0.0.1)/sql_test"
	db, err = sqlx.Connect("mysql", dsn)
	if err != nil {
		return
	}
	return
}

type user struct {
	Id   int
	Name string
	Age  int
}

func main() {
	err := initDB()
	if err != nil {
		fmt.Printf("connect err: %v\n", err)
		return
	}
}

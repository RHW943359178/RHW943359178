package main

import (
	"fmt"
	_ "src/github.com/go-sql-driver/mysql"
	"src/github.com/jmoiron/sqlx"
)

var db *sqlx.DB

func initDB() (err error) {
	dsn := `root:123456@tcp(127.0.0.1)/sql_test`
	db, err = sqlx.Connect("mysql", dsn)
	if err != nil {
		fmt.Printf("Connect DB failed, err: %v\n", err)
		return
	}
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)
	return
}

type user struct {
	Id   int
	Name string
	Age  int
}

//	查询单条数据
func queryRow() {
	sqlStr := `select id, name, age from user where id = ?;`
	var u user
	err := db.Get(&u, sqlStr, 1)
	if err != nil {
		fmt.Printf("get failed, err: %v\n", err)
		return
	}
	fmt.Printf("u: %#v\n", u)
}

//	查询多条数据
func queryRows(n int) {
	sqlStr := `select id, name, age from user where id > ?;`
	var users []user
	err := db.Select(&users, sqlStr, n)
	if err != nil {
		fmt.Printf("query failed, err: %v\n", err)
		return
	}
	fmt.Printf("users: %#v\n", users)

}

func main() {
	err := initDB()
	if err != nil {
		fmt.Printf("connect err: %v\n", err)
		return
	}
	//queryRow()
	queryRows(0)
}

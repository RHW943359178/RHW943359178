package main

import (
	"database/sql"
	"fmt"
	_ "src/github.com/go-sql-driver/mysql"
)

//	声明数据库连接对象
var db *sql.DB

//	声明用户列表的结构体
type user struct {
	id   int
	name string
	age  int
}

//	定义一个初始化数据库的函数
func initDB() (err error) {
	dsn := "root:123456@tcp(127.0.0.1)/sql_test"
	db, err = sql.Open("mysql", dsn)
	if err != nil {
		return
	}
	err = db.Ping()
	if err != nil {
		return
	}
	return err
}

//	查询单条sql
func queryRow(id int) {
	sqlStr := `select id, name, age from user where id=?`
	var u user
	err := db.QueryRow(sqlStr, id).Scan(&u.id, &u.name, &u.age)
	if err != nil {
		fmt.Printf("scan failed, err: %v\n", err)
		return
	}
	fmt.Printf("%#v\n", u)
}

func main() {
	err := initDB()
	if err != nil {
		fmt.Printf("init database falied, err: %v\n", err)
		return
	}
	fmt.Println("数据库连接成功")
	queryRow(1)
}

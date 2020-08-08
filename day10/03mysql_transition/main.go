package main

import (
	"database/sql"
	"fmt"
	_ "src/github.com/go-sql-driver/mysql"
)

var db *sql.DB //	是一个连接池对象

//	Go	连接mysql示例
func initDB() (err error) {
	//	数据库信息
	dsn := "root:123456@tcp(127.0.0.1)/sql_test"
	//	连接数据库
	db, err = sql.Open("mysql", dsn) //	不会校验用户名和密码正确
	if err != nil {                  //	dsn格式不正确时会报错
		return
	}
	err = db.Ping() //	尝试连接数据库 =》 校验的标准
	if err != nil {
		return
	}
	fmt.Println("数据库连接成功")
	return
}

func transaction() {
	tx, err := db.Begin() //	开启事务
	if err != nil {
		if tx != nil {
			_ = tx.Rollback() //	回滚
		}
		fmt.Printf("begin trans failed, err； %v\n", err)
		return
	}
	sqlStr1 := `update user set age=age-2 where id = ?;`
	sqlStr2 := `update user set age=age+2 where idd = ?;`
	//	执行sql1
	_, err = tx.Exec(sqlStr1, 1)
	if err != nil {
		_ = tx.Rollback()
		fmt.Printf("exec sql1 failed, err: %v\n", err)
		return
	}
	//	执行sql2
	_, err = tx.Exec(sqlStr2, 2)
	if err != nil {
		fmt.Printf("exec sql2 failed, err: %v\n", err)
		return
	}
	//	上面两步SQL都执行成功，就提交本次事务
	err = tx.Commit()
	if err != nil {
		//	要回滚
		_ = tx.Rollback()
		fmt.Println("提交出错了，要回滚")
		return
	}
	fmt.Println("执行事务成功！")
}

func main() {
	err := initDB()
	if err != nil {
		fmt.Printf("init DB failed, err: %v\n", err)
		return
	}
	transaction()
}

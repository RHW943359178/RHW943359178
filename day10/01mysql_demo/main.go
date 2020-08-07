package main

import (
	"database/sql"
	"fmt"
	_ "src/github.com/go-sql-driver/mysql"
)

// Go 连接mysql示例

func main() {
	//	数据库信息
	dsn := "root:123456@tcp(127.0.0.1)/goday10"
	//	连接数据库
	db, err := sql.Open("mysql", dsn) //	不会校验用户名和密码是否正确
	if err != nil {                   //	dsn格式不正确时会报错
		fmt.Printf("dsn: %s fialed, err: %v\n", dsn, err)
		return
	}
	err = db.Ping() //	尝试连接数据库 =》 校验的标准
	if err != nil {
		fmt.Printf("open %s fialed, err: %v\n", dsn, err)
		return
	}
	fmt.Println("连接数据库成功")

}

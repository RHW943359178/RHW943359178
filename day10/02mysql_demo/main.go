package main

import (
	"database/sql"
	"fmt"
	_ "src/github.com/go-sql-driver/mysql"
)

var db *sql.DB //	是一个连接池对象

// Go 连接mysql示例
func initDB() (err error) {
	//	数据库信息
	dsn := "root:123456@tcp(127.0.0.1)/sql_test"
	//	连接数据库
	db, err = sql.Open("mysql", dsn) //	不会校验用户名和密码是否正确
	if err != nil {                  //	dsn格式不正确时会报错
		return
	}
	err = db.Ping() //	尝试连接数据库 =》 校验的标准
	if err != nil {
		return
	}
	//	设置数据库连接的最大连接数
	//db.SetMaxOpenConns(10)
	return
}

type user struct {
	id   int
	name string
	age  int
}

//	查询
func query(id int) {
	var u1 user
	//	1.写查询单条记录的sql查询语句
	sqlStr := `select id, name, age from user where id=?;`
	//	2.执行
	rowObj := db.QueryRow(sqlStr, id) //	从连接池里拿一个连接出来去数据库查询单条记录
	//	3.拿到结果
	err := rowObj.Scan(&u1.id, &u1.name, &u1.age) //	必须对 rowObj 对象调用Scan方法，scan方法内置释放数据库连接到连接池
	if err != nil {
		fmt.Printf("query data failed: %v\n", err)
		return
	}
	//for i := 1; i < 12; i++ {
	//
	//	db.QueryRow(sqlStr, 1)
	//	fmt.Printf("已经进行了第%v次查询\n", i)
	//}
	fmt.Printf("u1: %#v\n", u1)
}

//	查询多条
func queryMore(n int) {
	//	1. SQL语句
	sqlStr := `select id, name, age from user where id > ?;`
	//	2. 执行
	rows, err := db.Query(sqlStr, n)
	if err != nil {
		fmt.Printf("exec %s query failed, err: %v\n", sqlStr, err)
		return
	}
	//	关闭rows
	defer rows.Close()
	//	循环取值
	for rows.Next() {
		var u1 user
		err := rows.Scan(&u1.id, &u1.name, &u1.age)
		if err != nil {
			fmt.Printf("scan failed, err: %v\n", err)
			return
		}
		fmt.Printf("u1: %#v\n", u1)
	}
}

//	插入数据
func insert() {
	//	1. 写SQL语句
	sqlStr := `insert into user(name, age) values("图朝阳", 48);`
	//	2. exec
	ret, err := db.Exec(sqlStr)
	if err != nil {
		fmt.Printf("insert failed, err: %v\n", err)
		return
	}
	//	如果是插入数据的操作，能够拿到插入数据的id
	id, err := ret.LastInsertId()
	if err != nil {
		fmt.Printf("get id failed, err: %v\n", err)
		return
	}
	fmt.Println("id: ", id)
}

//	更新操作
func updateRow(newAge int, id int) {
	sqlStr := `update user set age=? where id = ?;`
	ret, err := db.Exec(sqlStr, newAge, id)
	if err != nil {
		fmt.Printf("update failed, err: %v\n", err)
		return
	}
	n, err := ret.RowsAffected()
	if err != nil {
		fmt.Printf("get id failed, err: %v\n", err)
		return
	}
	fmt.Printf("更新了%d行的数据\n", n)
}

//	删除操作
func deleteRow(id int) {
	sqlStr := `delete from user where id = ?;`
	ret, err := db.Exec(sqlStr, id)
	if err != nil {
		fmt.Printf("delete failed, err: %v\n", err)
		return
	}
	n, err := ret.RowsAffected()
	if err != nil {
		fmt.Printf("get id failed, err: %v\n", err)
		return
	}
	fmt.Printf("删除了%d行数据\n", n)
}

func main() {
	err := initDB()
	if err != nil {
		fmt.Printf("init DB failed, err: %v\n", err)
		return
	}
	//query(2)
	//queryMore(2)
	//insert()
	//updateRow(39, 5)
	deleteRow(5)
}

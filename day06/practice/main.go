package main

import "fmt"

//	ini配置文件解析器

//	创建结构体
type mysqlConfig struct {
	Address  string `ini:"address"`
	Port     string `ini:"port"`
	Username string `ini:"username"`
	Password string `ini:"password"`
}

func loadIni(x interface{}) {

}

func main() {
	//var mc mysqlConfig
	//loadIni(&mc)
	//fmt.Println(mc.Address, mc.Port, mc.Password, mc.Username)
	fmt.Println()
}

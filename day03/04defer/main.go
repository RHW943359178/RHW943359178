package main

import "fmt"

//	defer

func main() {
	deferDemo()
}

//	defer 多用于函数结束之前释放资源（文件句柄，数据库连接，socket连接）
func deferDemo() {
	fmt.Println("start")
	defer fmt.Println("哈哈哈")	//	defer 把后面的语句延时加载到函数即将返回的时候再执行
	defer fmt.Println("黑恶黑")	//	一个函数中可以有多个 defer 语句
	defer fmt.Println("呜呜呜")	//	多个 defer 语句按照先进后出的顺序执行
	fmt.Println("end")
}
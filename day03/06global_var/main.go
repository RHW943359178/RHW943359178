package main

import "fmt"

//	变量的作用域

var x = 100	//	定义一个全局变量

func f1() {
	//	函数中查找变量的顺序
	//	1.先在函数内部查找
	//	2.找不到之后再往函数外面找，一直找到全局
	fmt.Println(x)
}

func main() {
	f1()
}
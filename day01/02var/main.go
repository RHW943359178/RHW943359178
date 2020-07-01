package main

import "fmt"

//	声明变量
//	Go语言中推荐使用小驼峰命名法
// var name string
// var age int
// var isOk bool

//	批量声明
var (
	name string // ''
	age  int    // 0
	isOk bool   //	false
)

func main() {
	// var sex string
	name = "理想"
	age = 16
	isOk = true
	//	Go语言中非变量声明必须使用，不使用就编译不过去
	fmt.Printf("name: %s", name) //	%s: 占位符 使用 name 这个变量的值去替换占位符

	//	声明变量时同时赋值

	//	类型推导（根据值来判断该变量是什么类型）

	//	简短变量声明（只能在函数中使用）
}

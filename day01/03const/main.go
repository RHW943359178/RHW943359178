package main

import "fmt"

//	常量
//	定义了常量之后不能修改
//	在程序运行期间不会改变的量
const pi = 3.14

//	批量声明变量

//	批量声明常量时，如果某一行没有赋值，则默认和上一行一致
const (
	n1 = 100
	n2
	n3
)

func main() {
	fmt.Println("n1:", n1)
	fmt.Println("n2:", n2)
	fmt.Println("n3:", n3)
}

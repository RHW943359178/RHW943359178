package main

import "fmt"

//	匿名函数

//var f1 = func (x, y int) {
//	fmt.Println(x + y)
//}

func main() {
	//	函数的内部没有办法声明带有名字的函数
	f1 := func (x, y int) {
		fmt.Println(x + y)
	}
	f1(10, 20)

	//	如果只是调用一次的函数，可以简写成立即调用函数
	func (x, y int) {
		fmt.Println(x + y)
	}(1, 2)
}
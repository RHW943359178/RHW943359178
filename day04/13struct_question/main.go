package main

import "fmt"

type myInt int

func (m myInt) hello() {
	fmt.Println("我是一个int")
}

func main() {
	m := myInt(100) //	强制类型转换
	m.hello()
	test()
}

func test() {
	//	声明一个 int32 类型的变量，它的值是10
	//	方法1
	//var x int32
	//x = 10
	////	方法2
	//var x int32 = 10
	////	方法3
	//var x = int32(10)
	////	方法4
	//x := int32(10)

	//	注意：x := int(10) 默认表示的是 int 类型
}

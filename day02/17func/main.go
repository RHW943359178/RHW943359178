package main

import "fmt"

//	函数

//	函数存在的意义：
//	函数是一段代码的封装

//	把一段代码封装出来，到一个函数中，然后通过函数名调用

//	求和（有参数，有返回值）
func sum(x int, y int)(ret int) {
	return x + y
}

//	没有返回值
func  f1(x int, y int)  {
	fmt.Print(x + y)
}

//	没有返回值，没有参数
func f2() {
	fmt.Print("f2")
}

//	没有参数但是有返回值
func f3() int{
	return 3
}

//	参数可以命名也可以不命名
//	命名的返回值就相当于再函数中声明了一个变量
func f4(x int, y int)(ret int) {
	ret = x + y
	return	//	命名之后，return之后可以默认不加返回值，返回值就默认是 ret
}

//	多个返回值
func f5()(int, int) {
	return 1, 2
}

//	参数类型的简写：当参数中连续多个参数的类型一致时，我们可以将前面的那个参数类型省略
func f6(x, y int) int {
	return x + y
}

//	可变长参数
//	可变长参数必须放在函数参数的最后
//	Go语言中没有默认参数这个概念

func f7(x string, y ...int){
	fmt.Println(x)
	fmt.Println(y)	//	y的类型是切片，具体的类型就是 ...后跟的类型
}


func main() {
	r := sum(10, 20)
	fmt.Println(r)

	//	多个返回值接收
	m, n := f5()
	fmt.Print(m, n)

	f7("下雨了")
	f7("下雨了", 1, 2, 3, 4, 5, 6, 7)

}

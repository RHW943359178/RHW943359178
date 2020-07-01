package main

import "fmt"

//	day03	复习

func main() {
	var ages [30]int	//	声明了一个变量 ages，它是 [30]int类型
	fmt.Println(ages)
	//array()
	//forArray()
	planarArray()
}

//	数组初始化的三种方式
func array() {
	var ages [30]int
	//	方式一：指定数组的固定长度
	ages = [30]int{1, 2, 3, 4, 5, 6}
	fmt.Println(ages)

	//	方式二：扩展运算符，自动计算数组长度
	var ages2 = [...]int{1, 3, 5, 7, 8, 9}
	fmt.Println(ages2)

	//	方式三：指定指定下标的元素
	var ages3 = [...]int{1: 100, 99: 200}
	fmt.Println(ages3)

}

//	数组的循环
func forArray()  {
	var a = [...]string{"北京", "上海", "深圳", "广州"}
	//	方式一：for + 长度循环
	for i := 0; i < len(a); i++ {
		fmt.Println(a[i])
	}

	//	方式二：for + range 循环，可以同时获得索引和值
	for key, value := range a {
		fmt.Printf("%v: %v\n", key, value)
	}
}

//	二维数组
func planarArray() {
	var a [3][2]int	//	[[1 2] [3 4] [5 6]]

	a = [3][2]int{
		[2]int{1, 2},
		[2]int{3, 4},
		[2]int{5, 6},
	}
	fmt.Println(a)
	//	多维数组只有最外层能使用扩展运算符 ...
	//	数组是值类型
	x := [3]int{1, 2, 3}
	y := x	//	将 x 的值拷贝了一份给 y
	y[1] = 200	//	修改的是副本 y， 不会影响 x
	fmt.Println(x)
	f1(x)
	fmt.Println(x)
}

func f1(a [3]int) {
	//	Go 语言中函数传递的都是值（Ctrl C 和 Ctrl V）
	a[1] = 100	//	此处修改的是副本的值，不会影响 x 本来的值
}



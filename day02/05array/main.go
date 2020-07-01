package main

import "fmt"

//	数组

//	存放元素的容器
//	必须制定存放的元素的类型和容量（长度）
//	数组的长度是数组类型的一部分

func main() {
	var a1 [3]bool //	[true false true]
	var a2 [4]bool //	[true true false false]

	fmt.Printf("a1:%T a2:%T\n", a1, a2)

	//	数组的初始化
	//	如果不初始化：默认元素都是零值（布尔值：false，整形和浮点型都是0，字符串： ""）
	fmt.Println(a1, a2)

	//	1.初始化方式1
	a1 = [3]bool{true, true, true}

	//	初始化方式2
	//	根据初始值自动推断数组的长度是多少
	a10 := [...]int{1, 2, 3, 4, 5, 6, 7}
	fmt.Println(a10)

	//	初始化方式3（一）
	a3 := [5]int{0, 1}
	fmt.Println(a3)

	//	初始化方式3（二）根据索引来初始化
	a4 := [5]int{0: 1, 4: 2}
	fmt.Println(a4)

	//	数组的遍历
	citys := [...]string{"北京", "上海", "深圳"}
	//	1. 根据索引遍历
	for i := 0; i < len(citys); i++ {
		fmt.Println(citys[i])
	}
	//	2.for range遍历
	for i, v := range citys {
		fmt.Println(i, v)
	}

	//	多维数组
	//	[[1 2] [3 4] [5 6]]
	var a11 [3][2]int
	a11 = [3][2]int{
		[2]int{1, 2},
		[2]int{3, 4},
		[2]int{5, 6},
	}
	fmt.Println(a11)

	//	多维数组的遍历
	for _, v := range a11 {
		fmt.Println(v)
		for _, v2 := range v {
			fmt.Println(v2)
		}
	}

	//	数组是值类型
	b1 := [3]int{1, 2, 3} //	b1: [1 2 3]
	b2 := b1              //	b2: [1 2 3] 类似于Ctrl+C  Ctrl+V 不会改变原来的值
	b2[0] = 100           //	b2: [100 2 3]
	fmt.Println(b1, b2)   //	b1: [1 2 3], b2: [100 2 3]

	// add()
	index()
}

func add() {
	//	1.求数组[1, 3, 5, 7, 8]所有元素的和
	//	2.找出数组中和为指定值的两个元素的下标，比如从数组[1, 3, 5, 7, 8]中找出和为8的两个元素的下标分别为(0,3)和(1,2)。

	var a = [5]int{1, 3, 5, 7, 8}
	var b = int(0)

	//	遍历
	for i := 0; i < len(a); i++ {
		b += a[i]
	}
	fmt.Println(b)
}

func index() {
	var a = [5]int{1, 3, 5, 7, 8}
	for i := 0; i < len(a); i++ {
		for j := 0; j < len(a); j++ {
			if a[i]+a[j] == 8 {
				fmt.Printf("(%d, %d)\n", i, j)
			}

		}
	}
}

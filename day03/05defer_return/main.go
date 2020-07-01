package main

import "fmt"

//	Go语言中的 return 不是原子操作，是分为两步
//	第一步：返回值赋值
//	第二步：真正的 RET 返回
//	函数中如果存在 defer ，那么 defer 的执行时机在 第一步 和 第二步 之间

func f1() int {
	x := 5
	defer func() {
		x++	//	修改的是 x， 不是返回值
	}()
	return x
}

func f2() (x int) {
	defer func() {
		x++
	}()
	return 5	//	返回值=x
}

func f3() (y int) {
	x := 5
	defer func() {
		x++	//	修改的是x
	}()
	return x	//	返回值 = x = 5
}
func f4() (x int) {
	defer func(x int) {
		x++	//	改变的是函数中的副本
	}(x)
	return 5	//	返回值 = x = 5
}
func main() {
	fmt.Println(f1())
	fmt.Println(f2())
	fmt.Println(f3())
	fmt.Println(f4())
}
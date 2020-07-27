package main

import (
	"fmt"
	"time"
)

func hello(i int) {
	fmt.Println("hello", i)
}

//	程序启动之后会创建一个主goroutine去执行
func main() {
	for i := 0; i < 100; i++ {
		//go hello(i)	//	开启一个单独的 gotoutine 去执行 hello 函数 （任务）
		go func(i int) {
			fmt.Println(i) //	用的是函数参数的那个i ，不是外面的 i
		}(i) //	闭包会导致多次打印i，所以说明goroutine是需要耗费资源的
	}
	fmt.Println("main")
	//	main 函数结束了 由 main 函数启动的 goroutine 也都结束了
	time.Sleep(time.Second)
}

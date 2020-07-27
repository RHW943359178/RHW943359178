package main

import "fmt"

var a []int
var b chan int //	需要指定通道中元素的类型

//b = make(chan int)	//	通道只有在初始化之后才能使用

func noBufChannel() {

	//b = make(chan int, 16)	//	通道只有在初始化之后才能使用
	go func() {
		x := <-b
		fmt.Println("后台goroutine从通道b中取到了", x)
	}()
	b <- 10 //	hang 住了
	fmt.Println("10发送到通道b中了...")
	fmt.Println(b)
}

func main() {

}

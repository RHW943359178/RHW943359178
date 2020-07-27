package main

import "fmt"

//	关闭通道

func main() {
	ch1 := make(chan int, 2)
	ch1 <- 10
	ch1 <- 20
	close(ch1)
	for x := range ch1 {
		fmt.Println(x)
	}
	<-ch1
	<-ch1
	x, ok := <-ch1
	fmt.Println(x, ok) //	在通道中如果值取完之后继续取值，不会报错，第一个参数会变成相应类型的零值，第二个参数会变为 false
}

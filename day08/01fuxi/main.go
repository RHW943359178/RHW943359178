package main

import (
	"fmt"
	"math/rand"
	"time"
)

func send(ch chan<- int) {
	for {
		n := rand.Intn(10)
		ch <- n
		time.Sleep(time.Second * 5)
	}
}

func main() {
	ch := make(chan int, 1)
	//ch <- 100
	//<-ch
	go send(ch)
	for {
		i, ok := <-ch
		fmt.Println(i, ok)
		time.Sleep(time.Second)
	}

}

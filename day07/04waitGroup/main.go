package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func f() {
	rand.Seed(time.Now().Unix()) //	保证内次执行的时候都不一样
	fmt.Println(time.Now().UnixNano())
	for i := 0; i < 5; i++ {
		r1 := rand.Int()
		r2 := rand.Intn(10) //	0 <= x < 10
		fmt.Println(r1, r2)
	}
}

func f1(i int) {
	defer wg.Done()
	time.Sleep(time.Millisecond * time.Duration(rand.Intn(300)))
	fmt.Println(i)
}

var wg sync.WaitGroup

func main() {
	//f()
	//	wg.Add(10)
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go f1(i)
	}
	//	如何知道这10个 goroutine 都结束了
	wg.Wait() //	等待 wg 的计数器减为0
}

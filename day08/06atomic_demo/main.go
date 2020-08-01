package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

var (
	x int64
	wg sync.WaitGroup
	lock sync.Mutex
)

func add() {
	//lock.Lock()
	//x++
	//lock.Unlock()
	atomic.AddInt64(&x, 1)
	wg.Done()
}

func main() {
	//for i := 0; i < 100000; i++ {
	//	wg.Add(1)
	//	go add()
	//}
	//wg.Wait()
	//fmt.Println(x)
	atomicEx()
}

func atomicEx() {
	ok := atomic.CompareAndSwapInt64(&x, 0, 200)
	fmt.Println(ok, x)
}
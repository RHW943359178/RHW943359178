package main

import (
	"fmt"
	"runtime"
	"sync"
)

var wg sync.WaitGroup

func a() {
	defer wg.Done()
	for i := 0; i < 10; i++ {
		fmt.Printf("A:%d\n", i)
	}
}

func b() {
	defer wg.Done()
	for i := 0; i < 10; i++ {
		fmt.Printf("B:%d\n", i)
	}
}

func main() {
	fmt.Println(runtime.NumCPU())
	//runtime.GOMAXPROCS(2)	//	默认CPU的逻辑核心数，默认跑满整个CPU
	//wg.Add(2)
	//go a()
	//go b()
	//wg.Wait()
}

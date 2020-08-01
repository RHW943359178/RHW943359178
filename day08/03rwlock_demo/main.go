package main

import (
	"fmt"
	"sync"
	"time"
)

var (
	x    = 0
	wg sync.WaitGroup
	lock sync.Mutex
	rLock sync.RWMutex
)

//	当读的操作远大于写的操作的时候，读写互斥锁的效率要高于互斥锁

func read() {
	defer wg.Done()
	//rLock.RLock()
	lock.Lock()
	fmt.Println(x)
	time.Sleep(time.Millisecond)
	//rLock.RUnlock()
	lock.Unlock()
}

func write() {
	defer wg.Done()
	//rLock.RLock()
	lock.Lock()
	x++
	time.Sleep(time.Millisecond * 5)
	//rLock.RUnlock()
	lock.Unlock()
}

func main() {
	start := time.Now()
	for i := 0; i < 10; i++ {
		go write()
		wg.Add(1)
	}
	for i := 0; i < 1000; i++ {
		go read()
		wg.Add(1)
	}
	wg.Wait()
	fmt.Println(time.Now().Sub(start))
}

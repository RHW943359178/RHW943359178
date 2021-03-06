package main

import (
	"fmt"
	"strconv"
	"sync"
	"time"
)

//	Go内置的map不是并发安全的
//var (
//	m = make(map[string]int)
//	lock sync.Mutex
//)
//
//
//func get(key string) int {
//	return m[key]
//}
//
//func set(key string, value int) {
//	m[key] = value
//}
//
//func main() {
//	wg := sync.WaitGroup{}
//	start := time.Now()
//	for i := 0; i < 100; i++ {
//		wg.Add(1)
//		go func(n int) {
//			key := strconv.Itoa(n)
//			lock.Lock()
//			set(key, n)
//			lock.Unlock()
//			fmt.Printf("k=%v, v=%v\n", key, get(key))
//			wg.Done()
//		}(i)
//	}
//	wg.Wait()
//	fmt.Println(time.Now().Sub(start))
//}

var (
	m2 sync.Map
	wg sync.WaitGroup
)

func main() {
	start := time.Now()
	for i := 0; i < 20; i++ {
		wg.Add(1)
		go func(n int) {
			key := strconv.Itoa(n)
			m2.Store(key, n)          //	必须使用sync.Map内置的Store方法设置键值对
			value, ok := m2.Load(key) //	必须使用sync.Map内置的Load方法根据key取值
			if ok {
				fmt.Printf("k=%v, v=%v\n", key, value)
			}
			wg.Done()
		}(i)
	}
	wg.Wait()
	fmt.Println(time.Now().Sub(start))
}

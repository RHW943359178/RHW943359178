package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

//使用goroutine和channel实现一个计算int64随机数各位数和的程序。
//开启一个goroutine循环生成int64类型的随机数，发送到jobChan
//开启24个goroutine从jobChan中取出随机数计算各位数的和，将结果发送到resultChan
//主goroutine从resultChan取出结果并打印到终端输出

type job struct {
	value int64
}

type result struct {
	job *job
	sum int64
}

var jobChan = make(chan *job, 100)
var resultChan = make(chan *result, 100)
var count = make(chan int, 24)
var wg sync.WaitGroup

func zhoulin(zl chan<- *job) {
	//循环生成int64类型的随机数
	for {
		x := rand.Int63()
		newJob := &job{
			value: x,
		}
		zl <- newJob
		time.Sleep(time.Millisecond * 1000)
	}
}

func baolude(zl <-chan *job, bld chan<- *result) {
	for {
		sum := int64(0)
		job := <-zl
		n := job.value
		for n > 0 {
			sum += n % 10
			n = n / 10
		}
		newResult := &result{
			job: job,
			sum: sum,
		}
		bld <- newResult
	}
}

func main() {
	//wg.Add(1)
	go zhoulin(jobChan)
	//	开启24个goroutine
	//wg.Add(24)
	for i := 0; i < 24; i++ {
		count <- i
		go baolude(jobChan, resultChan)

	}
	//从resultChan取出结果并打印到终端输出
	//for result := range resultChan {
	//	fmt.Printf("value: %d sum: %d\n", result.job.value, result.sum)
	//}
	//fmt.Printf("count: %#v\n", count)
	for a := range count {
		fmt.Printf("value: %d\n", a)
	}
}
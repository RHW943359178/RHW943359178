package main

import (
	"flag"
	"fmt"
	"os"
	"runtime/pprof"
	"time"
)

//	一段有有问题的代码
func logicCode() {
	var c chan int
	for {
		select {
		case v := <-c: //	阻塞，c未被初始化，不管是读或者取，都会被阻塞
			fmt.Printf("recv from chan, value:%v\n", v)
		default:
			//time.Sleep(time.Microsecond * 500)
		}
	}
}

func main() {
	var isCPUPprof bool //	是否开启CPUprofile的标志位
	var isMemPprof bool //	是否开启内存profile的标志位

	flag.BoolVar(&isCPUPprof, "cpu", false, "turn cpu pprof on")
	flag.BoolVar(&isMemPprof, "mem", false, "turn cpu pprof on")
	flag.Parse()

	if isCPUPprof {
		file, err := os.Create("./cpu.pprof")
		if err != nil {
			fmt.Printf("create cpu pprof failed, err:%v\n", err)
			return
		}
		err2 := pprof.StartCPUProfile(file)
		if err2 != nil {
			fmt.Println("StartCPUProfile err: ", err)
			return
		}
		defer pprof.StopCPUProfile()
	}
	for i := 0; i < 8; i++ {
		go logicCode()
	}
	time.Sleep(time.Second * 20)
	if isMemPprof {
		file, err := os.Create("./mem.pprof")
		if err != nil {
			fmt.Printf("create mem pprof failed, err:%v\n", err)
			return
		}
		pprof.WriteHeapProfile(file)
		file.Close()
	}
}

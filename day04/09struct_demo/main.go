package main

import "fmt"

//	结构体占用一块连续的内存空间

type x struct {
	a int8 //	8bit = 1byte
	b int8
	c int8
}

func main() {
	m := x{
		a: int8(10),
		b: int8(20),
		c: int8(30),
	}
	fmt.Printf("%p\n", &(m.a)) //	0xc00000a0e0
	fmt.Printf("%p\n", &(m.b)) //	0xc00000a0e1
	fmt.Printf("%p\n", &(m.c)) //	0xc00000a0e2
}

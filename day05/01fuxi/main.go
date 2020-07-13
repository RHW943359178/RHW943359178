package main

import "fmt"

//	复习结构体

func main() {
	var a = struct {
		x int
		y int
	}{10, 20}
	fmt.Println(a)
}

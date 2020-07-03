package main

import "fmt"

func yuanshuai(name string) {
	fmt.Println("Hello", name)
}

func lixiang(f func(string), name string) {
	f(name)
	//fmt.Println("aa")
}

func zhoulin() func(int, int) int {
	return func(x, y int) int {
		return x + y
	}
}

func main() {
	lixiang(yuanshuai, "理想")
	ret := zhoulin()
	fmt.Printf("%T\n", ret)
	sum := ret(10, 20)
	fmt.Println(sum)
}

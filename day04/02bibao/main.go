package main

import "fmt"

//	闭包

func low(f func()) {
	f()
}

func yuanshuai(name string) {
	fmt.Println("hello", name)
}

//	实现 low(yuanshuai)

func bi(f func(string), name string) func() {
	return func() {
		f(name)
	}
}

func main() {
	fc := bi(yuanshuai, "元帅")
	low(fc)
}

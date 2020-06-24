package main

import "fmt"

func main() {
	//	元素类型为 map 的切片
	var s1 = make([]map[int]string, 10, 10)

	//	对 map 进行初始化
	s1[0] = make(map[int]string, 1)
	s1[0][10] = "沙河"
	fmt.Println(s1)

	test()
}

func test() {
	//	值为切片类型的 map
	var m1 = make(map[string][]int, 10)
	m1["沙河"] = []int{10, 20, 30}
	fmt.Println(m1)
}
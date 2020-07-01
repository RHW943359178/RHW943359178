package main

import "fmt"

//	切片（Slice）复习

func main() {
	var s1 []int	//	没有分配内存，丛刻为 nil
	fmt.Println(s1)
	//slice()
	sliceCopy()
}


func slice() {
	var s1 []int
	//	初始化方式一：
	s1 = []int{1, 2, 3}
	fmt.Println(s1)

	//	初始化方式二：make 关键字
	s2 := make([]bool, 2, 4)
	fmt.Println(s2)
	fmt.Println(len(s2), cap(s2))

	//	make 关键字一定会为 切片分配内存，即使长度为0，切片也不会等于 nil
	s3 := make([]int, 0, 4)
	fmt.Println(s3)
	fmt.Println(s3 == nil)
}

func sliceCopy() {
	//	切片不存值，拷贝之后指向的都是同一个底层数组，所以操作的也是底层数组
	s1 := []int{1, 2, 3}	//	[1 2 3]
	s2 := s1
	fmt.Println(s2)
	s2[1] = 100
	fmt.Println(s2)
	fmt.Println(s1)
}
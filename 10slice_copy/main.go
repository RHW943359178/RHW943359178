package main

import "fmt"

//	copy

func main() {
	// a1 := []int{1, 3, 5}
	// a2 := a1 //	赋值
	// // var a3 []int	//	=> 此时 a3 是nil，无内存空间所以无法保存

	// var a3 = make([]int, 3, 5) //	使用 copy时，拷贝的目标切片如果长度小于被拷贝的长度，则取拷贝切片的长度，如果大于，则补0

	// copy(a3, a1) //	copy
	// fmt.Println(a1, a2, a3)

	// //	copy只是把切片里面的值拿出来放进另外一个切片，改变原来切片的值不会影响已经copy的切片
	// a1[0] = 100
	// fmt.Println(a1, a2, a3)

	// //	将 a1中索引为1 的3这个元素删除掉
	// a1 = append(a1[:1], a1[2:]...)
	// fmt.Println(a1)
	// fmt.Println(cap(a1))

	x1 := [...]int{1, 3, 5} //	数组
	s1 := x1[:]             //	切片
	fmt.Println(s1, len(s1), cap(s1))

	//	1. 切片不会保存具体的值
	//	2. 切片对应一个底层的数组
	//	3. 底层数组都是占用一块连续的内存

	//	修改了底层数组，切片删除了一个元素，但是数组一定是一块连续的内存地址，所以元素往左补齐，同时保留了原来的元素5
	s1 = append(s1[:1], s1[2:]...)
	fmt.Println(s1, len(s1), cap(s1))
	s1[0] = 100 //	修改底层数组
	fmt.Println(x1)

	test()
}

func test() {
	//	切片的练习
	var a = make([]int, 5, 10) //	创建切片，长度为5， 容量为10
	for i := 0; i < 10; i++ {
		a = append(a, i)
	}
	fmt.Println(a)

	//	结果：[0 0 0 0 0 0 1 2 3 4 5 6 7 8 9]
	//	分析： a 作为切片的值已经是[0 0 0 0 0 ], 在往后面循环追加，此时的切片的值就是 [0 0 0 0 0 0 1 2 3 4 5 6 7 8 9]，容量为20
}

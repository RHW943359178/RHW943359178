package main

import "fmt"

//	流程控制之跳出 for 循环

func main() {
	// for i := 0; i < 10; i++ {
	// 	if i == 5 {
	// 		break //	跳出for循环
	// 	}
	// 	fmt.Printf("%d\n", i)
	// }
	// fmt.Println("Over")

	//	当s = 5时，跳过此次循环（不执行for循环内部的打印），继续下一次循环
	for i := 0; i < 10; i++ {
		if i == 5 {
			continue //	继续下一次循环
		}
		fmt.Println(i)
	}
	fmt.Println("Over")
}

package main

import "fmt"

//	map

func main() {
	var m1 map[string]int

	println(m1 == nil) //	还没有初始化（没有在内存中开辟空间）

	m1 = make(map[string]int, 10) //	要估算好该map容量，避免在程序运行期间再动态扩容
	m1["理想"] = 18
	m1["num"] = 35
	fmt.Println(m1)

	fmt.Println(m1["num"])
	fmt.Println(m1["娜扎"]) //	如果不存在这个key，则拿到对应类型的零值

	//	约定俗成用 ok 接收返回的布尔值
	value, ok := m1["娜扎"]
	if !ok {
		fmt.Println("查无此key")
	} else {
		fmt.Println(value)
	}

	//	map 的遍历

	// for range 遍历
	for k, v := range m1 {
		fmt.Println(k, v)
	}
	//	只遍历 key，一个参数默认就是k
	for k := range m1 {
		fmt.Println(k)
	}

	//	只遍历value，匿名变量接收
	for _, v := range m1 {
		fmt.Println(v)
	}

	//	删除
	delete(m1, "num")
	fmt.Println(m1)
}

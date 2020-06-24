package main

import (
	"fmt"
	"strings"
)

//	1.写一个程序，统计一个字符串中每个单词出现的次数。比如：”how do you do”中how=1 do=2 you=1。

//	2.观察下面代码，写出最终的打印结果。
//func main() {
//	type Map map[string][]int
//	m := make(Map)
//	s := []int{1, 2}
//	s = append(s, 3)
//	fmt.Printf("%+v\n", s)
//	m["q1mi"] = s
//	s = append(s[:1], s[2:]...)
//	fmt.Printf("%+v\n", s)
//	fmt.Printf("%+v\n", m["q1mi"])
//}

func main() {
	test1()
}

func test1() {
	str := "how do you do how beautiful you are how are you"
	s := strings.Split(str, " ")
	fmt.Println(s)
	fmt.Printf("%T", s)
}
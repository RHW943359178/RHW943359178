package main

import (
	"fmt"
	"reflect"
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
	//test1()
	test2()
}

func test1() {
	str := "how do you do how beautiful you are how are you"
	s := strings.Split(str, " ")
	fmt.Println(s)
	//fmt.Printf("%T", s)
	//	思路：把放进 map 中

	strMap := make(map[string]int, 11)
	fmt.Println(strMap)

	for _, v := range s {
		//fmt.Println(v)
		_, ok := strMap[v]
		if !ok {
			strMap[v] = 1
		} else {
			strMap[v] += 1
		}
	}
	//fmt.Println(strMap)
	//	遍历 map
	for k, v := range strMap {
		fmt.Printf("%v=%d\n", k, v)
	}
}

func test2() {
	//	回文判断
	//	字符串从右至左和从左至右读是一样的，这就是回文
	//	上海自来水来自海上			山西运煤车煤运山西

	//	方式一：正序和倒叙循环比较字符串，如果相同，则是回文(通过切片比较)
	str := "上海自来水来自海上"

	strSlice := make([]string, 0, 10)
	lSlice := make([]string, 0, 10)
	rSlice := make([]string, 0, 10)

	for _, v := range str {
		strSlice = append(strSlice, string(v))
	}
	for i := 0; i < len(strSlice); i++ {
		lSlice = append(lSlice, strSlice[i])
	}
	for i := len(strSlice) - 1; i >= 0; i-- {
		rSlice = append(rSlice, strSlice[i])
	}

	if reflect.DeepEqual(lSlice, rSlice) {
		fmt.Println("是回文字符串")
	} else {
		fmt.Println("不是回文字符串")
	}
}


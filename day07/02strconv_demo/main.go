package main

import (
	"fmt"
	"strconv"
)

func main() {
	//从字符串中解析出对应的数据
	//str := "10000"
	////ret1 := int64(str)
	//
	//ret1, err := strconv.ParseInt(str, 10, 64)
	//if err != nil {
	//	fmt.Printf("parseint failed, err: %v", err)
	//	return
	//}
	//fmt.Printf("%#v %T\n", ret1, ret1)
	////	把数字转成字符串类型
	//i := int32(97)
	//
	//ret2 := fmt.Sprintf("%d", i)
	//fmt.Printf("%#v", ret2)

	AtoI()
}

func AtoI() {
	str := "10000"
	//	字符串转成int
	retInt, _ := strconv.Atoi(str)
	fmt.Printf("%#v, %T\n", retInt, retInt)

	i := 100
	retStr := strconv.Itoa(i)
	fmt.Printf("%#v, %T\n", retStr, retStr)
}

package main

import (
	"fmt"
	"time"
)

func main() {
	//now := time.Now()
	//fmt.Println(now)
	//fmt.Println(now.Year())
	//fmt.Println(now.Month())
	//fmt.Println(now.Day())
	//fmt.Println(now.Date())
	//fmt.Println(now.Hour())
	//fmt.Println(now.Minute())
	//fmt.Println(now.Second())
	////	时间戳
	//fmt.Println(now.Unix())
	//fmt.Println(now.UnixNano())
	////	将时间戳转为时间格式
	//ret := time.Unix(1595317710, 0)
	//fmt.Println(ret)
	//fmt.Println(ret.Year())
	//fmt.Println(ret.Day())
	ret := dateFormat("2006-01-02")
	fmt.Println(ret)

	//	2020/7/21 4:17
	ret2 := dateFormat("2006/1/2 3:4:5 PM")
	fmt.Println(ret2)
}

func dateFormat(t string) string {
	return time.Now().Format(t)
}

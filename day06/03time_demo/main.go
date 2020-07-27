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
	//ret := dateFormat("2006-01-02")
	//fmt.Println(ret)
	//
	////	2020/7/21 4:17
	//ret2 := dateFormat("2006/1/2 3:4:5 PM")
	//fmt.Println(ret2)
	//stringFormat()
	//sleepFormat()
	locationFormat()
}

func dateFormat(t string) string {
	return time.Now().Format(t)
}

//	按照对应的格式解析字符串类型的时间
func stringFormat() {
	timeStamp, err := time.Parse("2006-01-02", "2020-07-22")
	if err != nil {
		fmt.Printf("time Parse failed, err: %v\n", err)
		return
	}
	fmt.Println(timeStamp)
	fmt.Println(timeStamp.Unix())
}

//	sleep
func sleepFormat() {
	n := 5
	fmt.Println("开始sleep了")
	time.Sleep(time.Duration(n) * time.Second)
	fmt.Println("5秒已经结束")
}

func locationFormat() {
	now := time.Now() //	本地的时间
	//	按照东八区的时区和格式解析一个字符串格式的时间
	//	根据字符串加载时区
	loc, err := time.LoadLocation("Asia/Shanghai")
	if err != nil {
		fmt.Printf("load loc failed, err: %v\n", err)
		return
	}
	//	根据指定时区解析时间
	timeObj, err := time.ParseInLocation("2006-01-02 15:04:05", "2020-07-23 10:28:00", loc)
	if err != nil {
		fmt.Printf("parse time failed, err: %v\n", err)
		return
	}
	fmt.Println(timeObj)
	//	时间对象相减
	td := timeObj.Sub(now)
	fmt.Println(td)
}

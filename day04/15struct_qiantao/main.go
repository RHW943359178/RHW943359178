package main

import "fmt"

type address struct {
	province string
	city     string
}

type person struct {
	name    string
	age     int
	address //	匿名结构嵌套体
}

type company struct {
	name string
	addr address
}

func main() {
	p1 := person{
		name: "周琳",
		age:  90,
		address: address{
			province: "山东",
			city:     "威海",
		},
	}
	fmt.Println(p1)
	fmt.Println(p1.name, p1.address.city)
	//fmt.Println(p1.city)	//	非匿名结构嵌套体时，不能获取到
	fmt.Println(p1.city) //	现在自己的结构体里找这个字段，找不到的话再去匿名结构嵌套体里查找

}

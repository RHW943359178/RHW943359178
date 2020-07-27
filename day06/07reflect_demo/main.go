package main

import (
	"fmt"
	"reflect"
)

type myInt int64

func reflectType(x interface{}) {
	t := reflect.TypeOf(x)
	fmt.Printf("type: %v kind: %v\n", t.Name(), t.Kind())
}

func main() {
	//var a float32 = 3.14
	//reflectType(a)
	//var b int64 = 100
	//reflectType(b)
	var a *float32 //	指针
	var b myInt    //	自定义类型
	var c rune     //	类型别名
	reflectType(a)
	reflectType(b)
	reflectType(c)

	type person struct {
		name string
		age  int
	}
	type book struct {
		title string
	}
	d := person{
		name: "沙河小王子",
		age:  18,
	}
	e := book{
		title: "《跟小王子学Go语言》",
	}
	reflectType(d)
	reflectType(e)
}

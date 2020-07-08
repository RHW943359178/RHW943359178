package main

import "fmt"

//	结构体是值类型
//	一个结构体在内存是是连续的一块内存
type person struct {
	name   string
	gender string
}

//	go语言中函数的参数永远都是拷贝
func f(x person) {
	x.gender = "女" //	修改的是副本的gender，不会改变原本的值
}

//	如果想修改结构体的值，就需要传入结构体的内存地址
func f2(x *person) {
	x.gender = "女" //	语法糖，自动根据指针找对应的变量
}

func main() {
	var p person
	p.name = "周林"
	p.gender = "男"
	f(p)
	fmt.Println(p.gender) //	男

	f2(&p)
	fmt.Println(p.gender)

	//	结构体指针1
	var p2 = new(person)
	fmt.Printf("%T\n", p2)
	fmt.Printf("%p\n", p2)  //	p2保存的值就是一个内存地址, 实际就是new(person)的地址
	fmt.Printf("%p\n", &p2) //	这个才是p2的内存地址

	//	2.结构体指针2
	//	2.1 key-value初始化
	var p3 = person{ //	声明结构体的同时，初始化
		name:   "元帅",
		gender: "男",
	}
	fmt.Printf("%#v\n", p3)

	//	2.2 使用值列表的形式初始化（注意顺序要和结构体定义的字段顺序一致）
	p4 := person{
		"小王子",
		"男",
	}
	//	注意：两种方式不可以混用
	fmt.Printf("%#v\n", p4)
}

package main

import "fmt"

type person struct {
}

type cat struct {
}

type dog struct {
}

//	定义一个能叫的类型
type speaker interface {
	speak() //	只要实现了 speak 方法的变量都是 speaker 类型
}

func (p person) speak() {
	fmt.Println("啊啊啊")
}

func (c cat) speak() {
	fmt.Println("喵喵喵")
}

func (d dog) speak() {
	fmt.Println("汪汪汪")
}

func da(x speaker) {
	//	抽象概念，传进来什么就打什么
	x.speak()
}

func main() {
	var (
		p person
		c cat
		d dog
	)
	da(p)
	da(c)
	da(d)
}

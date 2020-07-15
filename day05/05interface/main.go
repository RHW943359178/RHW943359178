package main

import "fmt"

//	接口的实现

type animal interface {
	move()
	eat(string)
}

//	猫类
type cat struct {
	name string
	feet int8
}

//	鸡类
type chicken struct {
	feet int8
}

func (c cat) move() {
	fmt.Println("走猫步。。")
}

func (c cat) eat(food string) {
	fmt.Printf("猫吃%s。。\v", food)
}

func (ch chicken) move() {
	fmt.Println("跳着走。。")
}

func (ch chicken) eat(food string) {
	fmt.Printf("吃鸡%s。。\v", food)
}

func main() {
	var a1 animal //	定义一个接口类型的变量
	bc := cat{    //	定义一个 cat 类型的变量 bc
		name: "淘气",
		feet: 4,
	}
	kfc := chicken{
		feet: 4,
	}
	a1 = bc
	a1.eat("小黄鱼")
	fmt.Println(a1)
	a1 = kfc
	a1.eat("饲料")
	fmt.Printf("%T\n", a1)
}

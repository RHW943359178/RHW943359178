package main

import "fmt"

//	结构体中模拟实现其它语言中的继承

type animal struct {
	name string
}

//	给 animal 实现一个移动的方法
func (a animal) move() {
	fmt.Printf("%s会动\n", a.name)
}

//	狗类
type dog struct {
	feet   uint8
	animal //	animal 拥有的方法，dog此时也有了
}

//	给狗实现叫的方法
func (d dog) bark() {
	fmt.Printf("%s在叫：汪汪汪\n", d.name)
}

func main() {
	d := dog{
		feet: 4,
		animal: animal{
			name: "周琳",
		},
	}
	d.move()
	d.bark()
}

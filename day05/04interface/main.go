package main

import "fmt"

//	定义接口实现 run 方法
type car interface {
	run()
}

type ferrari struct {
	brand string
}

type porsche struct {
	brand string
}

func (f ferrari) run() {
	fmt.Printf("%v速度可以达到70码\n", f.brand)
}

func (p porsche) run() {
	fmt.Printf("%v速度可以达到90码\n", p.brand)
}

//	drive函数接收一个 car 类型的变量
func drive(c car) {
	c.run()
}

func main() {
	var (
		f = ferrari{
			brand: "法拉利",
		}
		p = porsche{
			brand: "保时捷",
		}
	)
	drive(f)
	drive(p)
}

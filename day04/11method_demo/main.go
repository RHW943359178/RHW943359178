package main

import "fmt"

//	方法
//	标识符：变量名、函数名、类型名、方法名
//	Go 语言中如果标识符首字母是大写的，就表示对外部可见（暴露的、共有的）

type dog struct {
	name string
}

type person struct {
	name string
	age  int
}

//	构造函数
func newPerson(name string, age int) *person {
	return &person{
		name: name,
		age:  age,
	}
}

//	构造函数
func newDog(name string) dog {
	return dog{
		name: name,
	}
}

// 方法是作用于特定类型的函数
//	接收者表示的是调用该方法的具体类型变量，多用类型名字的首字母小写
func (d dog) wang() {
	fmt.Printf("%s: 汪汪汪", d.name)
}

func (p person) run() {
	fmt.Printf("%s: 快点跑", p.name)
}

//	使用值接收者：传拷贝进去
func (p person) guonian() {
	p.age++
}

//	指针接收者：传内存地址
func (p *person) zhenguonian() {
	p.age++
}

func main() {
	//d1 := newDog("小黑")
	//d1.wang()
	p1 := newPerson("元帅", 18)
	//p1.run()
	fmt.Println(p1.age)
	//p1.guonian()
	p1.zhenguonian()
	fmt.Println(p1.age)
}

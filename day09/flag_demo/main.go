package main

import (
	"flag"
	"fmt"
)

//	flag获取命令行参数
func main() {
	//	创建一个标志位参数
	//name := flag.String("name", "王志", "请输入名字")
	//age := flag.Int("age", 200, "请输入真实年龄")
	//married := flag.Bool("married", false, "结婚了吗")
	//cTime := flag.Duration("ct", time.Second, "结婚多久了")

	var name string
	flag.StringVar(&name, "name", "王志", "请输入名字")

	//	使用flag
	flag.Parse() // 必须要加
	fmt.Println(name)
	//fmt.Println(*age)
	//fmt.Println(*married)
	//fmt.Println(*cTime)

}

package main

import (
	"fmt"
	"os"
)

/*
	函数版学生管理系统
	写一个系统能够查看\新增学生\删除学生
*/

type student struct {
	id   int64
	name string
}

func main() {
	for {
		//	1. 打印菜单
		fmt.Println("欢迎使用学生管理系统")
		fmt.Println(`
			1. 查看所有学生
			2. 新增学生
			3. 删除学生
			4. 退出
		`)
		fmt.Print("请输入操作指令")
		//	2. 等待用户选择要做什么
		var choice int
		fmt.Scanln(&choice)
		fmt.Printf("你选择了%d这个选项\n", choice)
		//	执行对应的函数
		switch choice {
		case 1:
			showAllStudent()
		case 2:
			addStudent()
		case 3:
			deleteStudent()
		case 4:
			os.Exit(1) //	退出
		default:
			fmt.Println("无效指令")
		}
	}
}

func showAllStudent() {

}
func addStudent() {

}
func deleteStudent() {

}

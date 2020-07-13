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

//	构造函数
func newStudent(id int64, name string) student {
	return student{
		id:   id,
		name: name,
	}
}

//	定义全局变量
var allStudents map[int64]string

func main() {
	allStudents = make(map[int64]string, 48) //	变量初始化
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
	//	循环遍历 map
	for k, v := range allStudents {
		fmt.Printf("学号：%v, 姓名: %v\n", k, v)
	}
}
func addStudent() {
	var (
		id   int64
		name string
	)
	//	1.获取用户输入
	fmt.Println("请输入用户学号：")
	fmt.Scanln(&id)
	fmt.Println("请输入用户姓名：")
	fmt.Scanln(&name)
	//	2.使用 student 的构造函数
	newStu := newStudent(id, name)
	//	3.判断新增的学号在原来 map 中是否存在
	_, ok := allStudents[id]
	if ok {
		fmt.Println("该学号已经存在！")
		return
	} else {
		allStudents[id] = newStu.name
	}
}
func deleteStudent() {
	//	获取用户输入
	var id int64
	fmt.Println("请输入需要删除的学生学号：")
	fmt.Scanln(&id)
	//	判断学号是否存在
	_, ok := allStudents[id]
	if ok {
		delete(allStudents, id)
	} else {
		fmt.Println("该学生不存在！请重试")
	}
}

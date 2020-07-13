package main

import (
	"encoding/json"
	"fmt"
)

//	结构体与json

//	1. 序列化：把 go语言中的结构体变量 --> json格式的字符串
//	2. 反序列化：把 json 格式的字符串 --> Go语言中能够识别的结构体变量

type person struct { //	在序列化的时候，字段首字母必须大写，否则数据无法导出
	Name string `json:"name"` //	通过反引号 `json:"name"` 的格式，可以将字段修改成需要的首字母小写的形式
	Age  int    `json:"age"`
}

func main() {
	p := person{
		Name: "周琳",
		Age:  9000,
	}
	//	序列化
	b, err := json.Marshal(p)
	if err != nil {
		fmt.Printf("marshal failed, err: %v\n", err)
		return
	}
	fmt.Printf("%v\n", string(b))
	//	反序列化
	str := `{"name":"理想","age":18}`
	var p2 person
	json.Unmarshal([]byte(str), &p2) //	函数传参永远都是传递副本拷贝，如果需要修改值徐奥传递指针
	fmt.Printf("%#v\n", p2)
}

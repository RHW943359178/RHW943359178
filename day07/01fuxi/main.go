package main

import (
	"encoding/json"
	"fmt"
)

//	反射的应用之一：json解析

type student struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func main() {
	str := `{"name":"周琳", "age":9000}`
	var stu1 student
	json.Unmarshal([]byte(str), &stu1)
	fmt.Printf("%#v\n", stu1)
}

package main

import (
	"fmt"
	"os"
)

func f() {
	var fileObj *os.File
	var err error
	fileObj, err = os.Open("./main.go")
	//	defer fileObj.Close()	//	如果err不为nil的话，则fileObj则为nil，此时调用Close方法会引发panic，所以这句代码不能出现在这个位置
	if err != nil {
		fmt.Printf("open file failed, err: %v", err)
		return
	}
	fmt.Println(*fileObj)
	defer fileObj.Close()
}

func main() {
	f()
}

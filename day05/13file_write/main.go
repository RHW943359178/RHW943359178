package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
)

func writeDemo() {
	fileObj, err := os.OpenFile("./xx.txt", os.O_APPEND|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		fmt.Printf("open file failed, err: %v\n", err)
		return
	}
	//	write
	fileObj.Write([]byte("zhoulinmengbile\n"))
	//	writeString
	fileObj.WriteString("周琳解释不了\n")
	fileObj.Close()
}

func writeDemo2() {
	fileObj, err := os.OpenFile("./xx.txt", os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		fmt.Printf("open file failed, err: %v", err)
		return
	}
	defer fileObj.Close()
	//	创建一个写的对象
	wr := bufio.NewWriter(fileObj)
	wr.WriteString("hello, 沙河") //	写到缓存中
	wr.Flush()                  //	将缓存中的内容写入文件
}

func writeDemo3() {
	str := "hello, 沙河沙河"
	err := ioutil.WriteFile("./xx.txt", []byte(str), 0644)
	if err != nil {
		fmt.Printf("write file failed, err: %v", err)
		return
	}
}

//	打开文件内容

func main() {
	//writeDemo2()
	writeDemo3()
}

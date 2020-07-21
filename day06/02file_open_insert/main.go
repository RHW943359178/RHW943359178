package main

import (
	"fmt"
	"io"
	"os"
)

func writeFileInsert() {
	//	打开原文件
	fileObj, err := os.OpenFile("./sb.txt", os.O_RDWR, 0644)
	if err != nil {
		fmt.Printf("open file failed, err: %v\n", err)
		return
	}
	//	打开临时文件
	tmpFile, err := os.OpenFile("./sb.tmp", os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Printf("create file failed, err: %v\n", err)
		return
	}
	defer tmpFile.Close()
	//	设置首次读取长度(是原文件的读取长度，先读后写)
	var ret [2]byte
	n, err := fileObj.Read(ret[:])
	if err != nil {
		fmt.Printf("read file failed, err: %v\n", err)
		return
	}
	//	写入临时文件
	tmpFile.Write(ret[:n])
	//	写入新增的内容
	var z = []byte{'c', 'b', 'a'}
	tmpFile.Write(z)
	//	紧接着把原文件的后续内容写入文件
	var x [1024]byte
	for {
		//	依旧是先读后写
		n, err := fileObj.Read(x[:])
		//	判断是否读完
		if err == io.EOF {
			tmpFile.Write(x[:n])
			break
		}
		if err != nil {
			fmt.Printf("read file failed, err: %v\n", err)
			return
		}
		tmpFile.Write(x[:n])
	}
	fileObj.Close()
	tmpFile.Close()
	//	修改名称
	os.Rename("./sb.tmp", "./sb.txt")
}

func main() {
	writeFileInsert()
}

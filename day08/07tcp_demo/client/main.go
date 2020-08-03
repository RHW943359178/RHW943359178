package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:20000")
	if err != nil {
		fmt.Println("connect err: ", err)
		return
	}
	defer conn.Close()
	inputReader := bufio.NewReader(os.Stdin)
	for {
		//	获取用户输入
		input, _ := inputReader.ReadString('\n')
		inputInfo := strings.TrimSpace(input)
		if inputInfo == "exit" {
			return
		}
		_, err := conn.Write([]byte(inputInfo))
		if err != nil {
			fmt.Println("write err: ", err)
		}
		var buf [512]byte
		n, err := conn.Read(buf[:])
		if err != nil {
			fmt.Println("read err: ", err)
			return
		}
		fmt.Println(string(buf[:n]))
	}
}

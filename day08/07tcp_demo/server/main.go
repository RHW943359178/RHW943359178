package main

import (
	"bufio"
	"fmt"
	"net"
)

//	tcp server

func main() {
	listener, err := net.Listen("tcp", "127.0.0.1:20000")
	if err != nil {
		fmt.Println("listen err: ", err)
		return
	}
	for {
		conn, err := listener.Accept() //	建立连接
		if err != nil {
			fmt.Println("connect err: ", err)
			break
		}
		go connect(conn) //	起一个goroutine
	}
}

func connect(conn net.Conn) {
	defer conn.Close() //	关闭连接
	for {
		var buf [128]byte
		reader := bufio.NewReader(conn)
		n, err := reader.Read(buf[:])
		if err != nil {
			fmt.Println("read err: ", err)
			break
		}
		rcvStr := string(buf[:n])
		fmt.Println("收到client端发来的数据：", rcvStr)
		//	把收到的数据再发回给client端
		conn.Write([]byte(rcvStr))
	}
}

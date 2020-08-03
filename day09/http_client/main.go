package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	response, err := http.Get("http://127.0.0.1:9090/xxx/?name=alex&age=18")
	if err != nil {
		fmt.Printf("get url failed, err:%v\n", err)
		return
	}
	//	从response中把服务端返回的数据读出来
	b, err := ioutil.ReadAll(response.Body) //	我在客户端读出来服务端返回的响应的body
	if err != nil {
		fmt.Printf("read response failed, err:%v\n", err)
		return
	}
	fmt.Println(string(b))
}

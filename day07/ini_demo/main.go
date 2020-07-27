package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"reflect"
	"strings"
)

type mysqlConfig struct {
	Address string `ini:"address"`
	Port string `ini:"port"`
	Username string `ini:"username"`
	Password string `ini:"password"`
}

type redisConfig struct {
	Host string `ini:"host"`
	Port string `ini:"port"`
	Password string `ini:"password"`
	Database string `ini:"database"`
}

//	Config
type Config struct {
	mysqlConfig `ini:"mysql"`
	redisConfig	`ini:"redis"`
}

func loadIni(fileName string, data interface{}) (err error) {
	//	0. 参数的校验
	//	0.1 传递进来的data参数必须是指针类型（因为需要在函数中对其赋值）
	t := reflect.TypeOf(data)
	fmt.Println(t, t.Name(), t.Kind())
	if t.Kind() != reflect.Ptr {
		err = errors.New("data param should be a pointer")	//	新建一个错误
		return
	}
	//	0.2 传递进来的参数必须是结构体类型指针（因为配置文件中各种键值对需要赋值给结构体的字段）
	if t.Elem().Kind() != reflect.Struct {
		err = errors.New("data param should be a struct pointer")	// 新建一个错误
		return
	}
	//	1.读文件得到字节类型的数据
	b, err := ioutil.ReadFile(fileName)
	if err != nil {
		return
	}
	//	string(b) //	将字节类型的文件内容转换成字符串
	lineSlice := strings.Split(string(b), "\r\n")
	fmt.Printf("%#v\n", lineSlice)
	//	2.一行一行得读数据
	for idx, line := range lineSlice {
		//	去掉字符串首尾的空格
		line = strings.TrimSpace(line)
		//	2.1 如果是注释就跳过
		if strings.HasPrefix(line, ";") || strings.HasPrefix(line, "#") {
			continue
		}
		if strings.HasPrefix(line, "[") {
			//	2.2 如果是[开头的就是表示节（section）
			if line[0] != '[' || line[len(line) - 1] != ']' {
				err = fmt.Errorf("line: %d syntax error", idx+1)
				return
			}
			//	把这一行首位的[]去掉，取到中间的内容把首尾的空格去掉拿到内容
			sectionName := strings.TrimSpace(line[1:len(line) - 1])
			if len(sectionName) == 0 {
				err = fmt.Errorf("line: %d syntax error", idx+1)
				return
			}
			//	根据字符串 sectionName 去data里面根据反射找到对应的结构体
			for i := 0; i < t.NumField(); i++ {
				field := t.Field(i)
				if section
			}
		} else {

		}
	}
	//	2.1如果是注释就跳过
	//	2.2 如果是[开头的就表示是节（section）
	//	2.3 如果不是[开头就是=分割的键值对

	return
}

func main() {
	var mc mysqlConfig
	err := loadIni("./conf.ini", &mc)
	if err != nil {
		fmt.Printf("load ini failed, err: %v\n", err)
		return
	}
	fmt.Println(mc.Address, mc.Port, mc.Password, mc.Username)
}
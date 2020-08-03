package split_string

import "strings"

//	Split 切割字符串
//	example
//	abc, b => [a c]

func Split(str string, sep string) []string {
	//	str: "babcbef"
	var ret []string
	index := strings.Index(str, sep)
	for index >= 0 {

		if str[:index] != "" {
			ret = append(ret, str[:index])
		}
		str = str[index+1:]
		index = strings.Index(str, sep)
	}
	ret = append(ret, str)
	return ret
}

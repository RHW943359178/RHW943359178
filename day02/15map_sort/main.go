package main

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano()) //	初始化随机数种子

	var scoreMap = make(map[string]int, 200)

	for i := 0; i < 100; i++ {
		key := fmt.Sprintf("stu%02d", i) //	生成 stu 开头的字符串
		value := rand.Intn(100)          //	生成 0~99 的随机整数
		scoreMap[key] = value
	}

	//fmt.Println(scoreMap)
	//	取出 map 中的所有key存入切片 keys
	var keys = make([]string, 0, 200)
	//fmt.Printf("keys=%v len(keys)=%d cap(keys)=%d\n", keys, len(keys), cap(keys))
	for key := range scoreMap {
		keys = append(keys, key)
	}
	//fmt.Println(keys)
	//	对切片排序
	sort.Strings(keys)
	//fmt.Println(keys)

	//	按照排序后的keys遍历map
	for _, key := range keys {
		fmt.Println(key, scoreMap[key])
	}

}

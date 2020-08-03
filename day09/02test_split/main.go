package main

import (
	"RHW943359178/day09/split_string"
	"fmt"
)

func main() {
	ret := split_string.Split("babcbef", "b")
	fmt.Printf("%#v\n", ret)
	ret2 := split_string.Split("bbb", "b")
	fmt.Printf("%#v\n", ret2)
	ret3 := split_string.Split("lklljjlkju", "b")
	fmt.Printf("%#v\n", ret3)
}

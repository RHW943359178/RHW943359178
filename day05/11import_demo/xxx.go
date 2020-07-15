package main //	只有main包才能编译成可执行化文件

import (
	calc "RHW943359178/day05/10calc" //	go中默认的导入方式是导入文件夹这一层，默认文件夹和包名相同，也可以取别名
	"fmt"
)

func init() {
	fmt.Println("自动执行")
}

func main() {
	ret := calc.Add(10, 20)
	fmt.Println(ret)
}

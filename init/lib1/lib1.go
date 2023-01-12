package lib1

import "fmt"

func Lib1Func() {
	fmt.Println("lib1.Func")
}

//首先调用init 函数，init函数中通常会执行一些初始化操作或者数据库操作或者环境变量的赋值
func init() {
	fmt.Println("lib1 init")
}

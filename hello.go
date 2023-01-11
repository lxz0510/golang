package main //主程序

import (
	"fmt"
	"time"
) // 导入包

func main() {
	//golang 不加分号
	fmt.Println("111")

	time.Sleep(1 * time.Second)

	fmt.Println("hello world!")
}

//go build 编译为二进制文件 （可执行程序）

//go run 编译+执行

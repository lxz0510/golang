package main

import "fmt"

func deferFunc() {
	fmt.Println("defer")
}

func retuenFunc() int {
	fmt.Println("return")
	return 0
}

func returnAndDefer() int {
	defer deferFunc()
	return retuenFunc()
}

func main() {
	// defer fmt.Println("end1")
	// defer fmt.Println("end2")

	// fmt.Println("1")
	// fmt.Println("2")
	// fmt.Println("3")

	returnAndDefer()
}

/*
1. defer作用是在当前函数关闭之前执行
2. 当有多个defer关键字时，维护一个栈进行处理。（先进后出
3. defer 和return的执行顺序，return先执行，defer后执行
*/

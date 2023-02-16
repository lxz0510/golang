package main

import (
	"fmt"
	"runtime"
	"time"
)

func a() {
	for i := 1; i < 10; i++ {
		fmt.Println("A:", i)
	}
}

func b() {
	for i := 1; i < 10; i++ {
		fmt.Println("B:", i)
	}
}

func main() {
	// go func(s string) {
	// 	//Goexit终止调用它的go程。
	// 	runtime.Goexit()
	// 	for i := 0; i < 2; i++ {
	// 		fmt.Println(s)
	// 	}
	// }("world")
	// // // 主协程
	// for i := 0; i < 2; i++ {
	// 	// 切一下，再次分配任务
	// 	//Gosched使当前go程放弃处理器，以让其它go程运行。它不会挂起当前go程，因此当前go程未来会恢复执行。
	// 	runtime.Gosched()

	// 	fmt.Println("hello")
	// }
	runtime.GOMAXPROCS(2)
	go a()
	go b()
	time.Sleep(time.Second)
}

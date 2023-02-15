package main

import (
	"fmt"
	"time"
)

func hello() {
	fmt.Println("Hello Goroutine!")
}

// 在程序启动时，Go程序就会为main()函数创建一个默认的goroutine。
func main() {
	// //串行
	// hello()
	// fmt.Println("main goroutine done!")
	//这里并没有输出hello中的信息，因为启动goroutine需要时间，在还没输出的时候main函数已经结束了，在main中启动的goroutine都会一并结束
	go hello()
	fmt.Println("main goroutine done!")
	time.Sleep(time.Second)
	//是因为我们在创建新的goroutine的时候需要花费一些时间，而此时main函数所在的goroutine是继续执行的。
}

//main函数所在的goroutine就像是权利的游戏中的夜王，其他的goroutine都是异鬼，夜王一死它转化的那些异鬼也就全部GG了。

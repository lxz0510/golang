package main

import (
	"fmt"
	"time"
)

// func worker(id int, wg *sync.WaitGroup) {
// 	//Done 表示一个协程的结束
// 	defer wg.Done()

// 	fmt.Printf("Worker %d starting\n", id)

// 	// 模拟工作
// 	for i := 0; i < 3; i++ {
// 		fmt.Printf("Worker %d working...\n", id)
// 	}

// 	fmt.Printf("Worker %d done\n", id)
// }

// func main() {
// 	//用于等待一组协程完成它们的任务。
// 	var wg sync.WaitGroup

// 	for i := 1; i <= 3; i++ {

// 		//新增一个协程
// 		wg.Add(1)
// 		go worker(i, &wg)
// 	}

// 	wg.Wait()
// 	fmt.Println("All workers done")
// }

//如果主协程退出了，那么其他未完成的协程也将被立即终止，不会继续执行
func longRunningTask() {
	for {
		fmt.Println("Long running task")
		time.Sleep(time.Second)
	}
}

func shortTask() {
	fmt.Println("Short task")
}

func main() {
	go longRunningTask()

	time.Sleep(time.Second * 5)

	go shortTask()

	time.Sleep(time.Second * 2)
}

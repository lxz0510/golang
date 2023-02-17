// 在代码中生硬的使用time.Sleep肯定是不合适的，Go语言中可以使用sync.WaitGroup来实现并发任务的同步.
package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func hello() {
	defer wg.Done()
	fmt.Printf("hello\n")
}

func main() {
	wg.Add(1)
	go hello()
	wg.Wait()

	fmt.Printf("Done\n")

}

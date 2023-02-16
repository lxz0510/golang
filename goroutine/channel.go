// 单纯地将函数并发执行是没有意义的。函数与函数间需要交换数据才能体现并发执行函数的意义。
//虽然可以使用共享内存进行数据交换，但是共享内存在不同的goroutine中容易发生竞态问题。为了保证数据交换的正确性，必须使用互斥量对内存进行加锁，这种做法势必造成性能问题。
//Go语言的并发模型是CSP（Communicating Sequential Processes），提倡通过通信共享内存而不是通过共享内存而实现通信。

//通过通信去共享内存，而不是通过共享内存去通信

//如果说goroutine是Go程序并发的执行体，channel就是它们之间的连接。channel是可以让一个goroutine发送特定值到另一个goroutine的通信机制。

//Go 语言中的通道（channel）是一种特殊的类型。通道像一个传送带或者队列，总是遵循先入先出（First In First Out）的规则，保证收发数据的顺序。每一个通道都是一个具体类型的导管，也就是声明channel的时候需要为其指定元素类型。
package main

import "fmt"

func revc(c chan int) {
	fmt.Println("1")

	ret := <-c
	fmt.Println("接收成功", ret)
}

func main() {
	// c := make(chan int)
	//无缓冲信道，要先有接收者，然后能够发送，不然会造成死锁。
	//无缓冲通道上的发送操作会阻塞，直到另一个goroutine在该通道上执行接收操作，这时值才能发送成功，两个goroutine将继续执行。
	//相反，如果接收操作先执行，接收方的goroutine将阻塞，直到另一个goroutine在该通道上发送一个值。
	// go revc(c)
	// fmt.Println("2")
	// c <- 10
	// fmt.Println("3")

	//make 一个有缓冲的信道，就像小区内的快递柜
	c := make(chan int, 1)
	c <- 10
	// value := <-c

	//快递柜装满了，不能继续push
	c <- 20
	// fmt.Println(value)

}

package main

import "fmt"

//交换失败，交换的只是局部变量中a和b的值

func swap(a, b int) {
	var temp int
	temp = a
	a = b
	b = temp

	fmt.Println("局部a->", a)
	fmt.Println("局部b->", b)
}

func main() {
	a := 1
	b := 2

	//swap
	swap(a, b)

	fmt.Println("a->", a)
	fmt.Println("b->", b)
}

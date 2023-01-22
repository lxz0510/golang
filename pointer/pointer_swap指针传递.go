package main

import "fmt"

func swap(a, b *int) {
	var temp int
	temp = *a
	*a = *b
	*b = temp

	fmt.Println("局部a->", *a)
	fmt.Println("局部b->", *b)
}

func main() {
	a := 1
	b := 2

	//swap
	swap(&a, &b)

	fmt.Println("a->", a)
	fmt.Println("b->", b)

}

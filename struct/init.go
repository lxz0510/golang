package main

import "fmt"

//定义一个结构体（把多种基础的数据类型组合在一起成为一个复杂的数据类型
type Book struct {
	title  string
	author string
}

func changeB1(b Book) {
	//值传递
	b.title = "change"
}

func changeB2(b *Book) {
	//指针传递
	b.title = "change"
}

func main() {
	var book1 Book

	book1.title = "war and peace"
	book1.author = "我"

	changeB1(book1)

	fmt.Println(book1)

	changeB2(&book1)

	fmt.Println(book1)

}

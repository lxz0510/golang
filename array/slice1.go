package main

import "fmt"

func main() {
	s1 := []int{1, 2, 3, 4}
	fmt.Printf("len = %d,slice = %v\n", len(s1), s1)
	var s2 []int
	//只是声明了这个切片的引用，但是还没有分配任何内存地址
	fmt.Printf("len = %d,slice = %v\n", len(s2), s2)

	//开辟10个内存空间，但是初始值为0
	// s2 = make([]int, 10)

	//最常见的一种方式
	s3 := make([]int, 10)

	fmt.Printf("len = %d,slice = %v\n", len(s3), s3)

	//判断一个切片是否为空
	if s2 == nil {
		fmt.Println("nil")
	} else {
		fmt.Println("not nil")
	}
}

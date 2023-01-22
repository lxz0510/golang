//切片容量的追加和截取
package main

import "fmt"

func main() {
	s1 := []int{100, 2, 3, 4}

	s2 := s1[0:2]

	fmt.Printf("len = %d,cap = %d,slice = %v\n", len(s2), cap(s2), s2)

	//***切片的截取只是改变了当前切片的头尾指针的位置，其引用都都是同一块内存地址，所以修改其中一个切片的值，内存地址中的值发生了改变，所有的切片都会呗改变

	//go中提供了copy函数来开辟出一块新的内存地址来做复制
	//先开辟新的内存地址
	s3 := make([]int, 3)
	copy(s3, s1)
	fmt.Printf("len = %d,cap = %d,slice = %v\n", len(s3), cap(s3), s3)

}

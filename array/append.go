//切片容量的追加和截取
package main

import "fmt"

func main() {
	//len长度 cap容量
	s1 := make([]int, 3, 5)
	fmt.Printf("len = %d,cap = %d,slice = %v\n", len(s1), cap(s1), s1)

	//追加元素

	s1 = append(s1, 1)
	fmt.Printf("len = %d,cap = %d,slice = %v\n", len(s1), cap(s1), s1)

	s1 = append(s1, 1)
	fmt.Printf("len = %d,cap = %d,slice = %v\n", len(s1), cap(s1), s1)

	//在动态数组中，当内存地址不够append操作时，go会自动扩充当前cap的两倍长度的地址
	s1 = append(s1, 1)
	fmt.Printf("len = %d,cap = %d,slice = %v\n", len(s1), cap(s1), s1)
}

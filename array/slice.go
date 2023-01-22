package main

import "fmt"

//引用传递
func change(arr []int) {
	for i := range arr {
		fmt.Println(i)
		arr[0] = 100
	}
}

func main() {
	a := []int{1} //动态数组，切片slice
	//动态数组本身就是指向当前数组的引用
	fmt.Println(a)

	// _表示匿名的变量
	change(a)
	fmt.Println(a)
}

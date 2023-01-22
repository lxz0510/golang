package main

import "fmt"

func main() {
	var myArr [10]int

	for i := 0; i < len(myArr); i++ {
		fmt.Println(myArr[i])
	}

	myArr2 := [10]int{1, 2, 3}
	for _, v := range myArr2 {
		fmt.Println(v)
	}
}

package main

import "fmt"

func main() {
	//var 定义，只是对m1进行了声明并没有初始化，并没有分配任何内存空间
	var m1 map[string]string
	if m1 == nil {
		fmt.Println("map 为空")

	}
	//make 分配内存空间
	m1 = make(map[string]string)
	m1["one"] = "1"
	m1["two"] = "2"
	m1["three"] = "3"
	//打印map
	fmt.Println(m1)

	//hash排序,并不是顺序排列

	//第二种声明方式
	m2 := make(map[string]string)
	fmt.Println(m2)

	//第三种声明
	m3 := map[string]string{
		"1": "one",
	}
	fmt.Println(m3)

	//添加
	m4 := make(map[string]string)
	m4["xiaolu"] = "pig"
	m4["xiaozhuang"] = "man"
	fmt.Println(m4)

	//遍历

}

package main

import "fmt"

//使用const 定义枚举类型

const (
	//可以在const()中使用一个关键字 iota 用于累加
	beijing  = 10 * iota //0
	shanghai             //1
	shenzhen
)

//iota 逐行累加1 （只有在const中进行累加效果）
const (
	a, b = iota + 1, iota + 2
	c, d
	e, f

	//公式可以变化，没有公式的行跟随离他最近的公式行
	g, h = iota + 10, iota + 11
	i, k
)

func main() {
	//常量（只读）
	const p int = 100

	fmt.Println(beijing)
	fmt.Println(shanghai)
	fmt.Println(shenzhen)
}

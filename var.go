package main

// 四种变量的声明方式

import "fmt"

/*
 1.Println 和Printf的区别  Println是非格式化 Printf格式化输出
*/

//关于声明全局变量
//一下三种声明都是可以的，:=不能被用来声明全局变量，只能用来在函数体内部声明局部变量

var C = 100
var A int = 100
var B = 100

//expected declaration, found Dsyntax
// D:= 100

func main() {
	//不初始化，默认值为0
	var a int
	fmt.Println("a = ", a)

	//声明一个变量的同时赋初始值
	var b int = 100
	fmt.Println("b = ", b)

	//声明一个变量的同时赋值，但是不声明类型，通过编译施给c 自动匹配
	var c = "100"
	fmt.Println("c = ", c)
	fmt.Printf("c value is %s, type of c = %T\n", c, c)

	//最常用的方法 省略var，直接自动匹配
	d := 100 //声明的同时赋值
	fmt.Println("d", d)

	g := 2.31
	fmt.Printf("g type is %T\n", g)

	//声明多个变量
	//在开发中我偏向分开声明变量

	var (
		o int    = 100
		p string = "123"
	)
	fmt.Printf("o type is %T\n p value is %s\n", o, p)
}

//问题遗留：int类型不能输出%s

package main

import "fmt"

//值传递

//	//b 局部变量 默认值为0
func changeVal(b int) { //执行p = a ，把当前a的值复制给内存b中的值 b = 10
	fmt.Println("b", b)
	b = 100
	//再把b中的值修改为100 ，a内存中的值依然没有改变
	fmt.Println(b)
}

func main() {

	a := 10
	fmt.Println(a)
	//执行p = a ，把当前a的值复制给内存b中的值 b = 10
	changeVal(a)
	fmt.Println("change后", a)

}

//*int 表示b是一个指向int类型的指针类型,也是一个变量，在内存中也占用地址
func changeVal1(b *int) {

	//&取地址
	//*取值 -> 取的是当前b内存中指向的地址中的值
	*b = 100
	//这里是关键，表示把当前b内存中的存取的内存地址指向的地方的内容修改为等号右边的值
}

func main1() {

	a := 10
	//需要传递一个指针类型，所以把a 的地址传递给函数
	changeVal1(&a)
	fmt.Println("change后", a)

}

package main

import "fmt"

//interface{} 表示一种万能类型，类似 ts中 的any

func myfunc(args ...interface{}) {
	for _, v := range args {
		//如何区分传进来的参数是什么数据类型呢，业务中对于不同的数据类型的传递参数可能会有不同的处理方式
		//golang中提供了 ·类型断言· 的机制
		_, ok := v.(string)
		if ok {
			fmt.Println("is string")
		} else {
			fmt.Println("not string")
		}
	}
}

func main() {
	myfunc(100, "sss")
}

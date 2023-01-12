package main

//目前使用gopath模式，由于项目没有创建在go root下没有不知道这里如何导入lib包，后续使用go mod重写
import (
	_ "lib1"      //匿名别名 导入但不使用，调用内部的方法使用
	. "lib2"      //直接导入到当前包中，直接使用包中的方法  Lib2Func() 最好不要使用，因为可能会存在不同包中同名方法的冲突
	myLib2 "lib2" //别名，方便开发中使用
)

func main() {
	// lib1.Lib1Func()
	myLib2.Lib2Func()
}

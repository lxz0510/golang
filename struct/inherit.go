//继承
package main

import "fmt"

type Human struct {
	name string
	sex  string
}

func (this *Human) Eat() {
	fmt.Println("human eat")
}
func (this *Human) Walk() {
	fmt.Println("human walk")
}

type SuperMan struct {
	Human //继承了Human类中的方法和属性
	level int
}

func (this *SuperMan) Walk() {
	fmt.Println("SuperMan walk")
}

func main() {
	zhangsan := Human{name: "zhangsan", sex: "man"}
	zhangsan.Eat()
	zhangsan.Walk()

	lisi := SuperMan{Human{name: "superrrrr", sex: "man"}, 100}
	fmt.Println(lisi)
	lisi.Walk()
}

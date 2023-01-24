package main

import "fmt"

//声明一个父类接口
type AnimalInterface interface {
	Sleep()
	GetName() string
	GetColor() string
}

//实现子类，实现接口的所有方法
type Cat struct {
	name  string
	color string
}

func (this *Cat) GetName() string {
	return this.name
}

func (this *Cat) GetColor() string {
	return this.color
}

func (this *Cat) Sleep() {
	fmt.Println("卜卜在睡觉")
}

type Dog struct {
	name  string
	color string
}

func (this *Dog) GetName() string {
	return this.name
}

func (this *Dog) GetColor() string {
	return this.color
}

func (this *Dog) Sleep() {
	fmt.Println("豆豆在睡觉")
}

func ShowAnimal(animal AnimalInterface) {
	animal.Sleep()
	fmt.Println("color", animal.GetColor())
	fmt.Println("name", animal.GetName())

}

func main() {

	//使用接口的调用
	// var animal AnimalInterface //接口的数据类型，父类的指针

	// //声明一个对象，把地址传递给父类指针
	// animal = &Cat{"卜卜", "white"}
	// //接口调用的方法其实就是当前指向的对象的方法，实现一个行为，多种表现的能力
	// animal.Sleep()

	c := Cat{"卜卜", "white"}
	//多态现象
	ShowAnimal(&c)

	//不使用接口的调用,
	// cat := Cat{name: "卜卜", color: "white"}
	// cat.Sleep()

	// dog := Dog{name: "豆豆", color: "white"}
	// dog.Sleep()
}

//多态的理解：多态是同一个行为具有多个不同表现形式或形态的能力。
// 都是动物在睡觉，但是对于不同的动物来说没有有不同的睡觉方式和时间，当有很多个动物类的时候，就可以把睡觉抽象为一个接口，
// 不同的动物在实现睡觉的方法是都要去重写接口中的睡觉方法以实现不同的表现形式

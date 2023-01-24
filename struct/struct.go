//go语言中的类就是 通过声明结构体并绑定方法
package main

import "fmt"

//定义一个英雄 Hero
type Hero struct {
	Name  string
	Level int
	Ad    int
}

//当前对象的copy
// func (this Hero) GetName() string {
// 	return this.Name
// }

// func (this Hero) SetLevel(level int) {
// 	this.Level = level
// }

//使用*Hero 调用的是当前对象的指针
func (this *Hero) GetName() string {
	return this.Name
}

func (this *Hero) SetLevel(level int) {
	this.Level = level
}

func main() {
	hero := Hero{Name: "shako", Level: 100, Ad: 10}
	heroName := hero.GetName()
	fmt.Println(heroName)
	hero.SetLevel(1000)
	fmt.Println(hero)
}

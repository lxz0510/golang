package main

import (
	"fmt"
	"reflect"
)

type Person struct {
	Name string
	Age  int
}

func main() {
	p := Person{"Alice", 30}
	v := reflect.ValueOf(p)

	fmt.Println("Type of v:", v.Type())

	for i := 0; i < v.NumField(); i++ {
		field := v.Field(i)
		fmt.Printf("Field %d: %s = %v\n", i, field.Type(), field.Interface())
	}
}

//这个示例定义了一个 Person 结构体，然后通过反射获取了这个结构体的值 p 的反射值 v。我们可以使用 v.Type() 来获取 v 的类型。在这个示例中，它将返回 reflect.Value 类型。
// 我们还可以使用 v.NumField() 来获取结构体字段的数量，然后通过 v.Field(i) 获取每个字段的反射值。我们可以使用 field.Type() 来获取字段的类型，使用 field.Interface() 来获取字段的值。

package main

import "fmt"

//接口，定义一种规范

type Car interface {
	GetName() string
	Run()
}

type BMW struct {
	Name string
}

func (b *BMW) GetName() string {
	return b.Name
}

func (b *BMW) Run() {
	fmt.Printf("bmw run")
}

func main() {
	var a interface{}	//空接口,可以保存任何类型变量
	var b int
	a = b	//可以赋值，因为b实现了a
	fmt.Printf("type of a %T\n", a)	//type of a int

	var c Car
	var bmw = BMW{
		Name: "my bmw",
	}
	c = &bmw
	fmt.Println(c.GetName())
	c.Run()
}

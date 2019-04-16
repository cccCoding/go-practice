package main

import "fmt"

//接口

type Student struct {
	Name string
	Age int
}

type Test interface {
	Print()
}

//实现Print函数就实现了Test接口
func (s Student) Print() {
	fmt.Println("Print", s)
}

func main() {
	var t Test
	fmt.Println(t)
	//t.Print()		//接口类型默认是一个指针，没有赋值，报错地址无效或空指针
	var stu = Student{
		Name:"stu1",
		Age:18,
	}

	t = stu
	t.Print()
	fmt.Println(t)
}

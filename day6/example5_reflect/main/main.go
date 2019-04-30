package main

import (
	"fmt"
	"reflect"
)

//反射，可以在运行时动态获取变量的相关信息

type Student struct {
	Name string
}

func test(b interface{}) {
	t := reflect.TypeOf(b)
	v := reflect.ValueOf(b)
	k := v.Kind()
	fmt.Println("type:", t)	//打印类型
	fmt.Println("value:", v) //打印值
	fmt.Println("kind:", k) //打印类别

	iv := v.Interface()	//转成interface
	stu, ok := iv.(Student)	//断言为Student
	if ok {
		fmt.Printf("%v %T\n", stu, stu)
	}
}

func main() {
	var a = Student{
		Name:"haha",
	}
	test(a)
}

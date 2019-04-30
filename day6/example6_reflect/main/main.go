package main

import (
	"encoding/json"
	"fmt"
	"reflect"
)

//反射，可以在运行时动态获取变量的相关信息

type Student struct {
	Name string `json:"name"`
	Age int
	B int
}

func (s Student) HaHa() {
	fmt.Println(s)
}

func (s *Student) HaHa1() {
	fmt.Println(s)
}

func test(b interface{}) {
	val := reflect.ValueOf(b)
	if val.Kind() != reflect.Ptr && val.Elem().Kind() != reflect.Struct {
		fmt.Println("expect struct")
		return
	}
	num := val.Elem().NumField()
	fmt.Printf("struct has %d fields\n", num)
	numM := val.Elem().NumMethod()
	fmt.Printf("struct has %d methods\n", numM)	//指针的方法不算。。

	val.Elem().Field(0).SetString("xxx")
	val.Method(0).Call(nil)
	val.Elem().Method(0).Call(nil)

	typ := reflect.TypeOf(b)
	fmt.Printf("tag:%s\n", typ.Elem().Field(0).Tag.Get("json"))
}

func main() {
	var a = Student{
		Name:"haha",
		Age:18,
	}

	result, _ := json.Marshal(a)
	fmt.Println("json result:", string(result))

	test(&a)
	fmt.Println(a)
}

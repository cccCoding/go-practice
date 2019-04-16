package main

import "fmt"

//方法

type Student struct {
	Name string
	Age int
}

func (s Student) set(name string, age int) {	//结构体，值类型，给副本赋值
	s.Name = name
	s.Age = age
}

func (s *Student) set2(name string, age int) {	//指针，赋值原变量成功
	s.Name = name
	s.Age = age
}

func main() {
	var s Student
	s.set("123", 456)
	fmt.Println(s)

	var s2 Student
	s2.set2("456", 789)
	fmt.Println(s2)
}

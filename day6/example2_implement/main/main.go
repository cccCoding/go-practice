package main

import (
	"fmt"
	"sort"
)

//接口，定义一种规范

type Student struct {
	Name string
}

type StudentArray []Student

func (a StudentArray) Len() int {
	return len(a)
}

func (a StudentArray) Less(i, j int) bool {
	return a[i].Name < a[j].Name
}

func (a StudentArray) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}

func main() {
	var stus StudentArray
	stu1 := Student{
		"123",
	}
	stu2 := Student{
		Name:"456",
	}
	stu3 := Student{
		Name:"124",
	}
	stus = append(stus, stu1)
	stus = append(stus, stu2)
	stus = append(stus, stu3)
	fmt.Println(stus)
	fmt.Println("---------------")
	sort.Sort(stus)
	fmt.Println(stus)
}

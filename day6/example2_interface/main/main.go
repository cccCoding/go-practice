package main

import (
	"fmt"
	"sort"
)

//感受接口之美

type Student struct {
	Name string
}

type StudentArray []Student
//实现以下三个方法，可以作为一种接口，传入sort.Sort方法进行排序
//sort.Sort方法通过以下三个方法可进行排序操作

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

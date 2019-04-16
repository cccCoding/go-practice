package main

import "fmt"

type Student struct {
	Name string
	Age int32
	score float32
}

func main() {
	var stu Student
	stu.Age = 18
	stu.Name = "hua"
	stu.score = 80

	fmt.Println(stu)
	fmt.Printf("Name:%p\n", &stu.Name)
	fmt.Printf("Age:%p\n", &stu.Age)
	fmt.Printf("score:%p\n", &stu.score)
	//{hua 18 80}
	//Name:0xc000088020
	//Age:0xc000088030
	//score:0xc000088038
	//字段在内存中是连续的，"hua"占十个字节，Age18占4个字节（int32）

	var stu1 = &Student{	//结构体初始化
		Age: 20,
		Name: "hei",
	}
	fmt.Println(stu1)
}

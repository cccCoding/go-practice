package main

import (
	"encoding/json"
	"fmt"
)

//tag，字段原信息，编译器会将其和字段保持在一起，可以通过反射接口取到

type People struct {
	Score float32
}

type Student struct {
	Name string `json:"name"`	//表示在json里面小写
	Age int `json:"student_age"`
	int		//匿名字段
	People	//匿名字段实现继承
	Score float32	//冲突情况，默认访问自己的
}

func main() {
	var stu = Student{
		Name: "haha",
		Age: 18,
	}
	stu.int = 10	//使用匿名字段
	stu.Score = 99.9	//emmm
	stu.People.Score = 88.8
	data, err := json.Marshal(stu)
	if err != nil {
		fmt.Println("json encode stu failed, err:", err)
		return
	}
	fmt.Println(string(data))	//打印{"name":"haha","student_age":18}
	fmt.Println(stu)

}

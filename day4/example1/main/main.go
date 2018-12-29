package main

import (
	"fmt"
	"github.com/pkg/errors"
)

func main() {
	var i int
	fmt.Println(i)

	j := new(int)	//new用来分配值类型内存，返回的是指针
	fmt.Println(j)

	*j = 100
	fmt.Println(*j)

	var a [5]int	//数组
	fmt.Println(a)

	var b []int		//切片
	b = append(b, 10, 11)	//append 用来追加元素到slice中
	fmt.Println(b)

	test()
	test1()
}

func test() {
	//捕获错误并打印
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()
	b := 0
	a := 100 / b
	fmt.Println(a)
}

func test1() {
	err := initConfig()
	if err != nil {
		panic(err)
	}
}

func initConfig() (err error) {
	return errors.New("init config failed")
}

package main

import "fmt"

//数组，同一种数据类型的固定长度的序列
//长度是数组类型的一部分，var a [5]int 和 var a [10]int是不同的类型
//数组是值类型，改变副本的值不会改变本身的值

func main() {
	test()

	var a [5]int
	test2(a)
	fmt.Println(a)

	test3(&a)
	fmt.Println(a)
}

func test() {
	var a [5]int
	b := a
	b[0]= 10
	fmt.Println(a)
}

func test2(arr [5]int) {
	arr[1] = 10
}

func test3(arr *[5]int) {
	(*arr)[1] = 10
}

package main

import "fmt"

func main() {
	test()
}

func test() {
	s1 := new([]int)	//分配内存，返回指针
	fmt.Println(s1)

	s2 := make([]int, 10)	//分配内存，返回切片slice
	fmt.Println(s2)

	//(*s1)[0] = 10	//报错，没有初始化
	//先make初始化
	*s1 = make([]int, 5)
	(*s1)[0] = 10
	fmt.Println(s1)

	s2[0] = 100
	fmt.Println(s2)
}

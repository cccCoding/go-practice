package main

import "fmt"

type op_func func(int, int) int	//自定义类型，是两个int作为参数，一个int作为返回值的func类型

func add(a, b int) int {
	return a + b
}

func operator(op op_func, a, b int) int {
	return op(a, b)
}

func test(a, b int) int {
	result := func(m, n int) int {	//将匿名函数赋值给result
		return m + n
	}
	return result(a, b)
}

func test2(a, b int) int {
	result := func(m, n int) int {	//将匿名函数执行后的值赋给result
		return m + n
	}(a, b)	//匿名函数直接调用
	return result
}

func main() {
	sum := operator(add, 100, 200)		//func(函数也是一种类型)
	fmt.Println(sum)

	sum = operator(func(i int, i2 int) int {	//匿名函数
		return i * i2
	}, 5, 6)
	fmt.Println(sum)
}

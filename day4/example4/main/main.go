package main

import (
	"fmt"
	"strings"
)

//闭包，一个函数和与其相关的引用环境组合而成的实体

func main() {
	var f = Adder()
	fmt.Println(f(1))
	fmt.Println(f(20))
	fmt.Println(f(300))

	func1 := makeSuffixFunc(".bmp")
	func2 := makeSuffixFunc(".jpg")
	fmt.Println(func1("test"))
	fmt.Println(func2("test"))
}

func Adder() func(int) int {
	var x int
	return func(delta int) int {
		x += delta
		return x
	}
}

func makeSuffixFunc(suffix string) func(string) string {
	return func(name string) string {
		if !strings.HasSuffix(name, suffix) {
			return name + suffix
		}
		return name
	}
}

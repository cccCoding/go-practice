package main

import (
	"fmt"
	"sort"
)

func testIntSort() {
	var a = [...]int{1,6,2,5,7}
	sort.Ints(a[:])		//只能排序切片，数组是值类型，只改变函数内拷贝不改变原数组
	fmt.Println(a)
}

func testStringSort() {
	var a = [...]string{"ag", "a", "fdg", "aaa"}
	sort.Strings(a[:])
	fmt.Println(a)
}

func main() {
	testIntSort()
	testStringSort()
}

package main

import "fmt"

//切片
func main() {
	testSlice()
	modifyString()
}

func testSlice() {
	var slice []int
	var arr = [...]int{1,2,3,4,5}
	slice = arr[1:3]
	fmt.Println(slice)
	fmt.Println(len(slice))
	fmt.Println(cap(slice))
}

func modifyString() {
	s := "你啊号hello world"	//字符串是不可变的
	s1:= []rune(s)		//要改变字符串先转成字符数组
	s1[1] = '0'
	str := string(s1)	//改变了字符数组后再转成字符串
	fmt.Println(str)
}



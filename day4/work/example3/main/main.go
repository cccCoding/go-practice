package main

import (
	"fmt"
)

func main() {
	//输入一个字符串，判断是否为回文（正反读相同）
	var str string
	fmt.Scanf("%s", &str)
	if isHuiWen(str) {
		fmt.Printf("%s is huiwen\n", str)
	} else {
		fmt.Printf("%s is not huiwen\n", str)
	}

	if isHuiWen2(str) {
		fmt.Printf("%s is huiwen!\n", str)
	} else {
		fmt.Printf("%s is not huiwen!\n", str)
	}
}

func isHuiWen(str string) bool {
	for i := 0; i < len(str); i++ {		//string遍历是字节集遍历
		if i == len(str)/2 {
			break
		}
		if str[i] != str[len(str)-1-i] {
			return false
		}
	}
	return true
}

func isHuiWen2(str string) bool {	//支持中文
	t := []rune(str)	//转为字符集
	length := len(t)
	for i, v := range t {
		fmt.Printf("i=%d,v=%c\n", i, v)
		if i == length/2 {
			break
		}
		if t[i] != t[length-1-i] {
			return false
		}
	}
	return true
}
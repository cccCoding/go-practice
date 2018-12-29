package main

import (
	"fmt"
)

func reverse(str string) string {
	var result string
	strLen := len(str)
	for i := 0; i < strLen; i++ {
		result = result + fmt.Sprintf("%c", str[strLen-i-1])
	}
	return result
}

func main() {
	var str1 = "hello"
	str2 := "world"

	str3 := fmt.Sprintf("%s %s", str1, str2)
	fmt.Println(str3)

	n := len(str3)
	fmt.Printf("len(str3)=%d\n", n)

	substr := str3[:5]
	substr2 := str3[6:]
	fmt.Println(substr, substr2)

	r := reverse(str3)
	fmt.Println(r)
}

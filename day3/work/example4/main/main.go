package main

import (
	"fmt"
	"strconv"
)

//判断一个数是否是水仙花数

func main() {
	var str string
	fmt.Scanf("%s", &str)

	var result = 0
	for i := 0; i < len(str); i++ {
		num := int(str[i] - '0')	//字符相减，参考字符编码
		result += (num*num*num)
	}
	number, err := strconv.Atoi(str)
	if err != nil {
		fmt.Printf("can not convert %s to int\n", str)
		return
	}
	if result == number {
		fmt.Printf("%d is shuixianhuashu", number)
	} else {
		fmt.Printf("%d is not shuixianhuashu", number)
	}
}

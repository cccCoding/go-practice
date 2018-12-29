package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

//计算两个大树相加的和，这两个大数会超过int64的表示范围

func main() {
	reader := bufio.NewReader(os.Stdin)
	line, _, err := reader.ReadLine()
	if err != nil {
		fmt.Println("read from console err:", err)
		return
	}
	strSlice := strings.Split(string(line), "+")
	if len(strSlice) != 2 {
		fmt.Println("please input a+b")
	}
	strNumber1 := strings.TrimSpace(strSlice[0])
	strNumber2 := strings.TrimSpace(strSlice[1])
	fmt.Println(multi(strNumber1, strNumber2))
}

func multi(str1, str2 string) (result string) {
	if len(str1) == 0 && len(str2) == 0 {
		result = "0"
		return
	}

	index1 := len(str1) - 1
	index2 := len(str2) - 1
	var left int

	for index1 >= 0 && index2 >= 0 {
		m := str1[index1] - '0'
		n := str2[index2] - '0'
		sum := int(m) + int(n) + left
		if sum > 9 {
			left = 1
		} else {
			left = 0
		}
		c := (sum % 10) + '0'
		result = fmt.Sprintf("%c%s", c, result)
		index1--
		index2--
	}

	for index1 >= 0 {
		m := str1[index1] - '0'
		sum := int(m) + left
		if sum > 9 {
			left = 1
		} else {
			left = 0
		}
		c := (sum % 10) + '0'
		result = fmt.Sprintf("%c%s", c, result)
		index1--
	}

	for index2 >= 0 {
		n := str2[index2] - '0'
		sum := int(n) + left
		if sum > 9 {
			left = 1
		} else {
			left = 0
		}
		c := (sum % 10) + '0'
		result = fmt.Sprintf("%c%s", c, result)
		index2--
	}

	if left == 1 {
		result = fmt.Sprintf("1%s", result)
	}

	return
}

package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	//输入一行字符，分别统计其中英文字母、空格、数字和其他字符个数
	reader := bufio.NewReader(os.Stdin)
	result, _, err := reader.ReadLine()
	if err != nil {
		fmt.Println("read from console err:", err)
		return
	}
	wc, sc, nc, oc := count(string(result))
	fmt.Println("wc=%d,sc=%d,nc=%d,oc=%d\n", wc, sc, nc, oc)
}

func count(str string) (wordCount, spaceCount, numberCount, otherCount int) {
	t := []rune(str)
	for _, v := range t {
		switch {
		case v >= 'a' && v <= 'z':
			fallthrough
		case v >= 'A' && v <= 'Z':
			wordCount++
		case v == ' ':
			spaceCount++
		case v >= '0' && v <= '9':
			numberCount++
		default:
			otherCount++
		}
	}
	return
}
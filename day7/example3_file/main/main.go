package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

type CharCount struct {
	chCount int
	numCount int
	spaceCount int
	otherCount int
}

func main() {
	file, err := os.Open("/tmp/test.txt")
	if err != nil {
		fmt.Println("read file failed, err:", err)
		return
	}
	defer file.Close()

	var count CharCount

	reader := bufio.NewReader(file)
	for {
		str, err := reader.ReadString('\n')
		if err == io.EOF {			//读完了会报这个错误
			fmt.Println("read string end")
			break
		}
		if err != nil {
			fmt.Println("read string failed, err:", err)
			return
		}
		fmt.Printf("read line success, result:%s", str)

		runeArr := []rune(str)
		for _, v := range runeArr {
			switch {
			case v >= 'a' && v <= 'z':
				fallthrough
			case v >= 'A' && v <= 'Z':
				count.chCount++
			case v == ' ' || v == '\t':
				count.spaceCount++
			case v >= '0' && v <= '9':
				count.numCount++
			default:
				count.otherCount++
			}
		}
	}
	fmt.Println(count)

}
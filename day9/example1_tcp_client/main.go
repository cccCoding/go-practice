package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:50000")
	if err != nil {
		fmt.Println("Error dialing", err.Error())
		return
	}
	defer conn.Close()
	inputReader := bufio.NewReader(os.Stdin)
	//从命令行读取数据写到服务端
	for {
		input, err := inputReader.ReadString('\n')
		if err != nil {
			fmt.Println("Error ReadString", err.Error())
			return
		}
		trimmedInput := strings.Trim(input, "\r\n")
		if trimmedInput == "Q" {	//输入 Q 则退出链接
			return
		}
		_, err = conn.Write([]byte(trimmedInput))
		if err != nil {
			fmt.Println("Error Write", err.Error())
			return
		}
	}
}

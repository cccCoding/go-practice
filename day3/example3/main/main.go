package main

import (
	"fmt"
	"os"
)

func read() {
	//假设打开一个数据库操作
	conn, err := openConn()
	defer func() {
		if conn != nil {
			conn.Close()
		}
	}()
}

func main() {
	var i int = 0
	//defer 把函数压进栈里，等程序返回再执行	//先进后出

	//特性方便用于资源释放
	file, _ := os.Open("path")
	defer file.Close()

	defer fmt.Println(i)	//defer语句中的变量，在defer声明的时候就决定了
	defer fmt.Println("let us find out")	//defer后声明，先执行

	i = 10
	fmt.Println(i)
}

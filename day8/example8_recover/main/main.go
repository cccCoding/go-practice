package main

import (
	"fmt"
	"time"
)

func test() {
	defer func() {		//捕获异常,协程挂了进程不挂，不影响其他协程操作
		if err := recover(); err!= nil {
			fmt.Println("panic:", err)
		}
	}()
	var m map[string]int
	m["1"] = 1
}

func calc() {
	for {
		fmt.Println("calc...")
		time.Sleep(time.Second)
	}
}

func main() {
	go calc()
	go test()
	time.Sleep(time.Second*10)
}

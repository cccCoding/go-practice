package main

import (
	"fmt"
	"time"
)

//递归

func main() {
	recusive(10)
}

func recusive(n int) {
	fmt.Println("hello")
	time.Sleep(time.Second)
	if n < 0 {
		return
	}
	recusive(n-1)
}

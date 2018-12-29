package main

import "fmt"

func add(arg... int) (length, sum int) {	//可变参数，0个或多个参数
	length = len(arg)
	for _, n := range arg {
		sum += n
	}
	return
}

func main() {
	len, sum := add(1,3,5,7)
	fmt.Printf("len=%d, sum=%d\n", len, sum)
}

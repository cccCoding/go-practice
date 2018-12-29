package main

import(
	"go-practice/day1/calc"
	"fmt"
)

func main()  {
	var pipe chan int
	pipe = make(chan int, 1)
	go calc.Add(100, 200, pipe)

	sum := <- pipe
	fmt.Println("sum=", sum)
}
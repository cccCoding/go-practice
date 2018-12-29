package main

import (
	"fmt"
)

func modify(a int) {
	a = 10
}

func modify1(a *int) {
	*a = 15
}

func main() {
	a := 5
	b := make(chan int, 1)

	fmt.Println("a=", a)
	fmt.Println("b=", b) //print address

	modify(a)
	fmt.Println("a=", a)

	modify1(&a)
	fmt.Println("a=", a)

	//swap
	a1 := 1
	a2 := 2
	a1, a2 = a2, a1
	fmt.Println("a1=", a1, "a2=", a2)
}

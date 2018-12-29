package main

import (
	"fmt"
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().Unix())
}

func main() {
	for i := 0; i < 10; i++ {
		a := rand.Int()
		fmt.Println(a)
	}

	for i := 0; i < 10; i++ {
		a := rand.Intn(100)
		fmt.Println(a)
	}

	for i := 0; i < 10; i++ {
		a := rand.Float32()
		fmt.Println(a)
	}

	var str = `
hello,
	hi,
	xixi,
		haha
	`
	fmt.Println(str)

	var b byte = 'c'
	fmt.Printf("%c\n", b)
	fmt.Println(b)
}

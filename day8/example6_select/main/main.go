package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int, 10)
	ch2 := make(chan int, 10)
	go func() {
		for i := 0; i < 10; i++ {
			ch <- i
			time.Sleep(time.Second)
			ch2 <- i*i
			time.Sleep(time.Second)
		}
	}()

	for {
		select {
		case v := <-ch:
			fmt.Println(v)
		case v := <-ch2:
			fmt.Println(v)
		default:
			fmt.Println("get data timeout")
			time.Sleep(time.Second)
		}
	}
}

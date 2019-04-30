package main

import "fmt"

func send(ch chan<- int, exitCh chan bool) {	//ch声明为只写
	for i := 0; i < 10; i++ {
		ch <- i
	}
	close(ch)
	exitCh <- true
}

func recv(ch <-chan int, exitCh chan bool) {	//ch声明为只读
	for {
		v, ok := <- ch
		if !ok {
			break
		}
		fmt.Println(v)
	}
	exitCh <- true
}

func main() {
	ch := make(chan int, 10)
	exitCh := make(chan bool, 2)
	go send(ch, exitCh)
	go recv(ch, exitCh)
	for i := 0; i < 2; i++ {
		<-exitCh
	}
}

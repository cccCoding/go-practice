package main

import (
	"fmt"
)

func calc(taskChan, resChan chan int, exitChan chan bool) { //判断素数，是则加到resChan中
	for v := range taskChan {
		flag := true
		for i := 2; i < v; i++ {
			if v%i == 0 {
				flag = false
				break
			}
		}
		if flag {
			resChan <- v
		}
	}
	exitChan <- true
}

func main() {
	intChan := make(chan int, 100)
	resultChan := make(chan int, 100)
	exitChan := make(chan bool, 8)
	go func() {
		for i := 0; i < 10000; i++ {
			intChan <- i
		}
		close(intChan)
	}()

	for i := 0; i < 8; i++ {	//8个goroutine并发执行calc
		go calc(intChan, resultChan, exitChan)
	}

	go func() {
		//等待所有计算的goroutine全部退出，关闭resultChan
		for i := 0; i < 8; i++ {
			<-exitChan
		}
		close(resultChan)
	}()

	//当resultChan关闭后，该循环会退出
	for v := range resultChan {
		fmt.Println(v)
	}
}

package main

import (
	"fmt"
	"sync"
	"time"
)

//了解并发与并行
//并发，单核多线程，看起来多线程程序都在执行，但同时只有一个在执行
//并行，多核多程序，多个线程程序都在执行

var (
	m = make(map[int]uint64)
	lock sync.Mutex
)

type task struct {
	n int
}

func calc(t *task) {	//计算阶乘
	var sum uint64 = 1
	for i := 1; i < t.n; i++ {
		sum *= uint64(i)
	}
	lock.Lock()
	m[t.n] = sum
	lock.Unlock()
}

func main() {
	for i := 0; i < 100; i++ {
		t := &task{n:i}
		go calc(t)
	}

	time.Sleep(10*time.Second)
	lock.Lock()
	fmt.Println(m)
	lock.Unlock()

	//go build -race day8/example1_goroutine/main/main.go
	//-race 看是否有资源竞争
}

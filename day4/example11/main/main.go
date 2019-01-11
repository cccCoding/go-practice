package main

import (
	"fmt"
	"math/rand"
	"sync"
	"sync/atomic"
	"time"
)

//go build -race day4/example11/main/main.go 	加上参数race可以检查有没有数据竞争

var lock sync.Mutex		//互斥锁
var rwLock sync.RWMutex	//读写锁，适用于读多写少

func testLock() {
	var a map[int]int
	a = make(map[int]int, 5)
	a[1] = 1
	for i := 0; i < 2; i++ {
		go func(b map[int]int) {
			lock.Lock()
			b[1] = rand.Intn(100)	//多个协程修改一个数据，需要加锁
			lock.Unlock()
		}(a)
	}
	lock.Lock()
	fmt.Println(a)	//读的时候也加上
	lock.Unlock()
}

func testRWLock() {
	var a map[int]int
	a = make(map[int]int, 5)
	a[1] = 1
	var count int64
	for i := 0; i < 2; i++ {
		go func(b map[int]int) {
			rwLock.Lock()
			//lock.Lock()
			b[1] = rand.Intn(100)	//多个协程修改一个数据，需要加锁
			time.Sleep(10*time.Millisecond)
			rwLock.Unlock()
			//lock.Unlock()
		}(a)
	}
	for i := 0; i < 100; i++ {
		go func(b map[int]int) {
			for {
				rwLock.RLock()
				//lock.Lock()
				time.Sleep(time.Millisecond)
				rwLock.RUnlock()
				//lock.Unlock()
				atomic.AddInt64(&count, 1)	//原子操作,写
			}
		}(a)
	}
	time.Sleep(3*time.Second)
	fmt.Println(atomic.LoadInt64(&count))		//原子操作，读

	//使用读写锁时，读取的次数count为217213
	//使用互斥锁时，count为2189
	//显然，在读多写少的情况下，用读写锁效率高，能读取更多次数
}

func main() {
	//testLock()
	testRWLock()
}

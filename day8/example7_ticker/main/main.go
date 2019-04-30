package main

import (
	"fmt"
	"time"
)

func main() {
	//一秒打印一次时间
	//t := time.NewTicker(time.Second)
	//for v := range t.C {
	//	fmt.Println(v)
	//}

	//超时控制
	t := time.NewTicker(time.Second)
	select {
	case <-t.C:
		fmt.Println("after 1s")
	}
	t.Stop()	//要关闭，防止内存泄漏
}

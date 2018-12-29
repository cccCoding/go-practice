package main

import "fmt"

//3,对于一个数n，求n的阶乘之和,即：1！+2！+3！+...+n！

func sum(n int) uint64 {
	var s uint64 = 1
	var sum uint64
	for i := 1; i <= n; i++ {
		s = s * uint64(i)
		fmt.Printf("%d! = %v\n", i, s)
		sum += s
	}
	return sum
}

func main() {
	var n int
	fmt.Scanf("%d", &n)

	fmt.Println("sum=", sum(n))
}

package main

import (
	"fmt"
	"math"
)

//1,判断101-200之间有多少个素数，并输出所有素数

func isPrime(n int) bool {
	for i := 2; i <= int(math.Sqrt(float64(n))); i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func main() {
	var n int
	var m int
	fmt.Scanf("%d%d", &n, &m)	//从终端传一个数到n

	var count int
	for i := n; i < m; i++ {
		if isPrime(i) {
			count++
			fmt.Printf("%d is sushu\n", i)
		}
	}
	fmt.Printf("number of prime between %d and %d is %d", n, m, count)
}

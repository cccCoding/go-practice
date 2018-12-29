package main

import "fmt"

func main() {
	//1000以内所有完数（一个数等于他的因子之和）
	for i := 1; i < 1000; i++ {
		if isPerfect(i) {
			fmt.Println(i)
		}
	}
}

func isPerfect(n int) bool {
	sum := 0
	for i := 1; i < n; i++ {
		if n%i == 0 {
			sum += i
		}
	}
	return sum == n
}

package main

import "fmt"

//冒泡排序
func bubbleSort(a []int) {
	var compareTime int
	var exchangeTime int
	for i := 0; i < len(a); i++ {
		for j := 0; j < len(a)-i-1; j++ {
			if a[j] > a[j+1] {
				a[j], a[j+1] = a[j+1], a[j]
				exchangeTime++
			}
			compareTime++
		}
	}
	fmt.Printf("bubbleSort compareTime=%d, exchangeTime=%d\n", compareTime, exchangeTime)
}

func main() {
	b := [...]int{6, 4, 5, 10, 1, 2, 999, 9999, 88}
	fmt.Println(b)
	bubbleSort(b[:])
	fmt.Println(b)
}

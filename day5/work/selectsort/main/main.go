package main

import "fmt"

//选择排序，每次把最小或者最大的数选出来
func selectSort(a []int) {
	var compareTime int
	var exchangeTime int
	for i := 0; i < len(a); i++ {
		for j := i+1; j < len(a); j++ {
			if a[i] > a[j] {
				a[i], a[j] = a[j], a[i]
				exchangeTime++
			}
			compareTime++
		}
	}
	fmt.Printf("selectSort compareTime=%d, exchangeTime=%d\n", compareTime, exchangeTime)
}

func main() {
	b := [...]int{6, 4, 5, 10, 1, 2, 999, 9999, 88}
	fmt.Println(b)
	selectSort(b[:])
	fmt.Println(b)
}

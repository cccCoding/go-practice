package main

import "fmt"

//插入排序，类似抓牌,从手里一张牌开始抓，每抓一张牌从后往前插
func selectSort(a []int) {
	var compareTime int
	var exchangeTime int
	for i := 1; i < len(a); i++ {	//从手里有一张牌开始
		for j := i; j > 0; j-- {	//从后往前插
			if a[j-1] > a[j] {
				a[j-1], a[j] = a[j], a[j-1]
				exchangeTime++
			}
			compareTime++
		}
	}
	fmt.Printf("insertSort compareTime=%d, exchangeTime=%d\n", compareTime, exchangeTime)
}

func main() {
	b := [...]int{6, 4, 5, 10, 1, 2, 999, 9999, 88}
	fmt.Println(b)
	selectSort(b[:])
	fmt.Println(b)
}

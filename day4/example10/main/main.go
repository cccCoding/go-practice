package main

import (
	"fmt"
	"sort"
)

func testMapSort() {
	var a map[int]int
	a = make(map[int]int, 5)
	a[5] = 1
	a[3] = 1
	a[9] = 1
	a[10] = 1
	a[2] = 1
	var keys []int
	for k, v := range a {
		fmt.Println(k, v)
		keys = append(keys, k)
	}
	fmt.Println("------sort------")
	sort.Ints(keys)
	for _, v := range keys {
		fmt.Println(v, "1")
	}
}

func main() {
	testMapSort()
}

package main

import "fmt"

//2,打印出100-999之间所有的水仙花数，水仙花数指一个三位数，其各位数字立方和等于该数

func isShuiXianHuaShu(i int) bool {
	bai := (i / 100) % 10
	shi := (i / 10) % 10
	ge := i % 10
	if (bai*bai*bai+ shi*shi*shi+ ge*ge*ge) == i {
		return true
	} else {
		return false
	}
}

func main() {
	for i := 100; i < 1000; i++ {
		if isShuiXianHuaShu(i) {
			fmt.Printf("%d is shuixianhuashu!\n", i)
		}
	}
}

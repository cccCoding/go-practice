package main

import "fmt"

//快速排序,用了递归的思想，把一个大数组的问题切分成多个小数组的子问题
//确定一个元素的最终位置，其左边的元素都比他小，右边都比他大
//该切片被该元素切为两部分，这两部分同时再去确认内部某个元素的最终位置，切割，重复操作
//当被切分出来的子数组长度为1时，有序，就不用再去切分了
func quickSort(a []int, left, right int) {
	if left >= right {
		return
	}

	//先选取一个元素,确定val所在的位置
	val := a[left]
	k := left
	for i := left+1; i <= right; i++ {
		if a[i] < val {	//保证比val小的数在其左边
			a[k] = a[i]
			a[i] = a[k+1]
			k++
		}
	}
	a[k] = val	//确定val所在的位置
	//例子
	//k=0,val=6
	//6 7 8 2 4 99 88
	//i = 3, 2<val
	//2 7 8 7 4 99 88
	//k=1
	//i = 4, 4<val
	//2 4 8 7 8 99 88
	//k = 2
	//2 4 6 7 8 99 88

	quickSort(a, left, k-1)
	quickSort(a, k+1, right)
}

func main() {
	b := [...]int{6, 4, 5, 10, 1, 2, 999, 9999, 88}
	fmt.Println(b)
	quickSort(b[:], 0, len(b)-1)
	fmt.Println(b)
}

package main

import "fmt"

//切片
func main() {
	testSlice()
	modifyString()
	testSlice2()
	testCopy()
}

func testCopy() {
	fmt.Println("testcopy-----------------------")
	var a = []int{1,2,3,4,5,6}
	b := make([]int, 1)
	copy(b, a)			//拷贝第一个元素过来，不追加
	fmt.Println(b)
}

func testSlice2() {
	fmt.Println("test2-----------------------")
	var a = [5]int{1,2,3,4,5}
	s := a[1:]
	fmt.Printf("s=%p a[1]=%p\n", s, &a[1])	//相同//通过数组得到的切片地址为数组对应index元素的地址
	fmt.Printf("len=%d,cap=%d\n", len(s), cap(s))

	s[1] = 100
	fmt.Printf("s=%p a[1]=%p\n", s, &a[1])	//相同，切片还是原来的切片
	fmt.Printf("len=%d,cap=%d\n", len(s), cap(s))
	s = append(s, 10)
	fmt.Printf("s=%p a[1]=%p\n", s, &a[1])	//不同，重新申请了内存，切片内部指向另一个数组
	fmt.Printf("len=%d,cap=%d\n", len(s), cap(s))		//指向的另一个数组长度翻倍了
	fmt.Println(s)
}

func testSlice() {
	var slice []int
	var arr = [...]int{1,2,3,4,5}
	slice = arr[1:3]
	fmt.Println(slice)
	fmt.Println(len(slice))
	fmt.Println(cap(slice))
}

func modifyString() {
	s := "你啊号hello world"	//字符串是不可变的
	s1:= []rune(s)		//要改变字符串先转成字符数组
	s1[1] = '0'
	str := string(s1)	//改变了字符数组后再转成字符串
	fmt.Println(str)
}



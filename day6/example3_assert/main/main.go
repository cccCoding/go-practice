package main

import "fmt"

//类型断言

func Test(a interface{}) {
	if b, ok := a.(int); ok {	//类型断言，tip：也可以断言为某个接口
		b += 3
		fmt.Println("ok", b)
	} else {
		fmt.Println("not ok", a)
	}
}

func classifier(items ...interface{}) {	//类型判断
	for i, item := range items {
		switch item.(type) {	//switch帮助断言
		case bool:
			fmt.Println("bool", i)
		case int:
			fmt.Println("int", i)
		default:
			fmt.Println("unknown", i)
		}
	}
}

func main() {
	var a = 3
	var b = "3"
	Test(a)
	Test(b)
	classifier(a, b)
}

package main

import "fmt"

type student struct {
	name string
}

func main() {
	var stuChan chan interface{}
	stuChan = make(chan interface{}, 10)
	stu := student{name:"stu1"}
	stuChan <- &stu

	var a interface{}
	a = <- stuChan
	var stu2 *student
	stu2, ok := a.(*student)
	if  !ok {
		fmt.Println("can't convert")
		return
	}
	fmt.Println(stu2)
}

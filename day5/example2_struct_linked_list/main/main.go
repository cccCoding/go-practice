package main

import (
	"fmt"
	"math/rand"
)

//链表

type Student struct {
	Name string
	Age int
	Score float32
	next *Student
}

func trans(p *Student) {
	for p != nil {
		fmt.Println(*p)
		p = p.next
	}
}

func insertTail(p *Student) {
	var tail = p
	for i := 0; i < 6; i++ {
		var stu = &Student{
			Name: fmt.Sprintf("student%d", i),
			Age:rand.Intn(100),
			Score:rand.Float32()*100,
		}
		tail.next = stu
		tail = stu
	}
}

func insertHead(p **Student) {
	for i := 0; i < 6; i++ {
		var stu = Student{
			Name: fmt.Sprintf("student%d", i),
			Age:rand.Intn(100),
			Score:rand.Float32()*100,
		}
		stu.next = *p
		*p = &stu
	}
}

func delNode(p *Student) {
	var prev = p
	for p != nil {
		if p.Name == "student3" {
			prev.next = p.next
			break
		}
		prev = p
		p = p.next
	}
}

func main() {
	var head = &Student{
		Name: "hua",
		Age: 18,
		Score: 81,
	}

	//insertTail(head)	//head尾部插入十个Student

	insertHead(&head)	//头部插入

	trans(head)

	fmt.Println()

	delNode(head)

	trans(head)
}

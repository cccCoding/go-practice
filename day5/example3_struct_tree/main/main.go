package main

import "fmt"

//二叉树

type Student struct {
	Name string
	Age int
	left *Student
	right *Student
}

func trans(root *Student) {
	if root == nil {
		return
	}
	fmt.Println(root)
	trans(root.left)
	trans(root.right)
}

func main() {
	var root = &Student{
		Name: "hua",
		Age: 18,
	}

	var left1 = new(Student)
	left1.Name = "left1"

	var right1 = new(Student)
	right1.Name = "right1"

	var left2 = new(Student)
	left2.Name = "left2"

	root.left = left1
	root.right = right1
	left1.left = left2

	trans(root)
}

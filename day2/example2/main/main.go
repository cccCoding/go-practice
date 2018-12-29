package main

import(
	"fmt"
	a "awesomeProject/day2/example2/add"
)

func init() {
	fmt.Println("main init")
}

func main() {
	fmt.Println("name=", a.Name)
	fmt.Println("age=", a.Age)
}


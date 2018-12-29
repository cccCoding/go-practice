package main

import (
	"fmt"
	"os"
)

func main() {
	var goos = os.Getenv("GOOS")
	fmt.Printf("the operating system is:%s\n", goos)
	path := os.Getenv("PATH")
	fmt.Printf("Path is %s\n", path)
}

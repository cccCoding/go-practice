package main

import (
	"github.com/cccCoding/go-practice/day6/example4_link/link"
)

func main() {
 	var intLink link.Link
	for i := 0; i < 10; i++ {
		intLink.InsertTail(i)
		//intLink.InsertHead(i)
	}
	intLink.Trans()
}

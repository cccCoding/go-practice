package main

import (
	"fmt"
	"github.com/cccCoding/go-practice/day7/example1_interface/balance"
	"math/rand"
	"os"
	"time"
)

func main() {
	insts := make([]*balance.Instance, 0)
	for i := 0; i < 4; i++ {
		host := fmt.Sprintf("192.168.%d.%d", rand.Intn(255), rand.Intn(255))
		one := balance.NewInstance(host, 8080)
		insts = append(insts, one)
	}

	var balanceName = "random"
	if len(os.Args) > 1 {
		balanceName = os.Args[1]
	}

	for {
		inst, err := balance.DoBalance(balanceName, insts)
		if err != nil {
			fmt.Println("do balance err:", err)
			return
		}
		fmt.Println(inst)
		time.Sleep(time.Second)
	}
}

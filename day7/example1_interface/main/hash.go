package main

import (
	"errors"
	"fmt"
	"github.com/cccCoding/go-practice/day7/example1_interface/balance"
	"hash/crc32"
	"math/rand"
)

//实现自己的负载均衡算法,不用修改原来的库

func init() {
	balance.RegisterBalancer("hash", &HashBalance{})
}

type HashBalance struct {
	key string
}

func (p *HashBalance) DoBalance(insts []*balance.Instance, key ...string) (inst *balance.Instance, err error) {
	if len(insts) == 0 {
		err = errors.New("No instance")
		return
	}
	var k = fmt.Sprintf("%d", rand.Int())
	if len(key) > 0 {
		k = key[1]
	}
	crcTable := crc32.MakeTable(crc32.IEEE)
	hash := crc32.Checksum([]byte(k), crcTable)
	lens := len(insts)
	index := int(hash) % lens
	inst = insts[index]
	return
}

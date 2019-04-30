package balance

import (
	"github.com/pkg/errors"
	"math/rand"
)

func init() {
	RegisterBalancer("random", &RandomBalance{})
}

type RandomBalance struct {}

func (rb *RandomBalance) DoBalance(insts []*Instance, key ...string) (inst *Instance, err error) {
	if len(insts) == 0 {
		err = errors.New("No instance")
		return
	}
	lens := len(insts)
	index := rand.Intn(lens)
	inst = insts[index]
	return
}

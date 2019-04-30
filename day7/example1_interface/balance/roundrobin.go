package balance

import (
	"github.com/pkg/errors"
)

//轮询

func init() {
	RegisterBalancer("roundrobin", &RoundBobinBalance{})
}

type RoundBobinBalance struct {
	currIndex int
}

func (rb *RoundBobinBalance) DoBalance(insts []*Instance, key ...string) (inst *Instance, err error) {
	if len(insts) == 0 {
		err = errors.New("No instance")
		return
	}
	lens := len(insts)
	inst = insts[rb.currIndex]
	rb.currIndex = (rb.currIndex+1) % lens
	return
}

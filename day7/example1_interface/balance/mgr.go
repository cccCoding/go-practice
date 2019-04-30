package balance

import "fmt"

//支持可扩展
type BalanceMgr struct {
	allBalancer map[string]Balancer
}

var (
	mgr = BalanceMgr{
		allBalancer:make(map[string]Balancer),
	}
)

func (p *BalanceMgr) registerBalancer(name string, b Balancer) {
	p.allBalancer[name] = b
}

func RegisterBalancer(name string, b Balancer) {
	mgr.registerBalancer(name, b)
}

func DoBalance(name string, insts []*Instance) (inst *Instance, err error) {
	if balancer, ok := mgr.allBalancer[name]; !ok {
		err = fmt.Errorf("Not found %s balancer", name)
	} else {
		fmt.Printf("use %s balancer\n", name)
		inst, err = balancer.DoBalance(insts)
	}
	return
}

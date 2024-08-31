package pool

import (
	"errors"
	"net/http"
	"sync"

	uuid "github.com/google/uuid"

	lc500Vm "github.com/yammerjp/lc500/worker/vm"
)

type Pool struct {
	vms map[string]*lc500Vm.VM
	mu  sync.Mutex
}

func NewPool() *Pool {
	return &Pool{
		vms: make(map[string]*lc500Vm.VM),
		mu:  sync.Mutex{},
	}
}

func (p *Pool) get(key string) *lc500Vm.VM {
	p.mu.Lock()
	defer p.mu.Unlock()
	if vm, ok := p.vms[key]; ok {
		return vm
	}
	vm := lc500Vm.NewVM()
	p.vms[key] = vm
	return vm
}

func (p *Pool) put(key string, vm *lc500Vm.VM) {
	p.mu.Lock()
	defer p.mu.Unlock()
	p.vms[key] = vm
}

func (p *Pool) delete(key string) {
	p.mu.Lock()
	defer p.mu.Unlock()
	delete(p.vms, key)
}

func (p *Pool) Generate() string {
	uuid, err := uuid.NewRandom()
	if err != nil {
		panic(err)
	}

	vm := lc500Vm.NewVM()
	p.put(uuid.String(), vm)
	return uuid.String()
}

func (p *Pool) CompileScript(key string, scriptStr string) error {
	vm := p.get(key)
	if vm == nil {
		return errors.New("vm not found")
	}
	return vm.CompileScript(scriptStr)
}

func (p *Pool) SetContext(key string, context *lc500Vm.VMContext) error {
	vm := p.get(key)
	if vm == nil {
		return errors.New("vm not found")
	}
	return vm.SetContext(context)
}

func (p *Pool) Run(key string, w http.ResponseWriter) error {
	vm := p.get(key)
	if vm == nil {
		return errors.New("vm not found")
	}
	vm.SetResponseWriter(w)
	return vm.RunScript()
}

func (p *Pool) Dispose(key string) {
	vm := p.get(key)
	if vm == nil {
		return
	}
	vm.Dispose()
	p.delete(key)
}

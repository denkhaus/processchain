package shared

import (
	"sync"
)

type HandlerContext struct {
	store map[string]interface{}
	mu    sync.Mutex
}

func (p *HandlerContext) Set(key string, value interface{}) {
	p.mu.Lock()
	defer p.mu.Unlock()
	p.store[key] = value
}

func (p *HandlerContext) Get(key string) interface{} {
	p.mu.Lock()
	defer p.mu.Unlock()
	if val, ok := p.store[key]; ok {
		return val
	}

	return nil
}

// func (p *HandlerContext) GetEntityContext() *EntityContext {
// 	ctx := p.Get("EntityContext")

// 	if ctx != nil {
// 		return ctx.(*EntityContext)
// 	}

// 	return nil
// }

type Handler func(value interface{}) error
type Handlers []Handler

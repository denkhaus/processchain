package shared

import (
	"fmt"

	"github.com/pkg/errors"
)

var (
	ErrPropertyAlreadySet = func(key string) error {
		return errors.Errorf("property with key %s already defined", key)
	}
)

type ModuleContext struct {
	name string
	data map[string]interface{}
}

func NewModuleContext(name string) *ModuleContext {
	ctx := ModuleContext{
		name: name,
		data: make(map[string]interface{}),
	}
	return &ctx
}

func (p *ModuleContext) MustGet(field string) interface{} {
	if val, ok := p.data[field]; ok {
		return val
	}
	panic(fmt.Sprintf("ModuleContext: MustGet: field %s undefined", field))
}

func (p *ModuleContext) MustString(field string) string {
	if value, ok := p.MustGet(field).(string); ok {
		return value
	}

	panic(fmt.Sprintf("ModuleContext: MustString: field %s not of type string", field))
}

func (p *ModuleContext) MustBool(field string) bool {
	if value, ok := p.MustGet(field).(bool); ok {
		return value
	}

	panic(fmt.Sprintf("ModuleContext: MustBool: field %s not of type bool", field))
}

func (p *ModuleContext) MustInt64(field string) int64 {
	if value, ok := p.MustGet(field).(int64); ok {
		return value
	}

	panic(fmt.Sprintf("ModuleContext: MustInt64: field %s not of type int64", field))
}

func (p *ModuleContext) Set(key string, value interface{}) *ModuleContext {
	if _, ok := p.data[key]; ok {
		panic(fmt.Sprintf("ModuleContext: %s", ErrPropertyAlreadySet(key)))
	}

	p.data[key] = value
	return p
}

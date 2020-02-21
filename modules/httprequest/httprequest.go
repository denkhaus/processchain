package httprequest

import "github.com/denkhaus/processchain/interfaces"

type HttpRequestOptionable interface {
	WithOptions(opts ...Option) HttpRequestStartable
}

type HttpRequestProceedable interface {
	Get() interfaces.Readable
	Delete() interfaces.Readable
	Post() interfaces.Readable
	Patch() interfaces.Readable
	Put() interfaces.Readable
}

type HttpRequestStartable interface {
	HttpRequestOptionable
	HttpRequestProceedable
}

type Chain struct {
	interfaces.Chain
}

func (p Chain) Delete() interfaces.Readable {
	return p
}

func Inherit(parent interfaces.Chain) Chain {
	chain := Chain{
		parent,
	}

	return chain
}

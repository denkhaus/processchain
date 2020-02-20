package httprequest

import "github.com/denkhaus/processchain/interfaces"

type Chain struct {
	interfaces.Chain
}

func Inherit(parent interfaces.Chain) interfaces.Chain {
	chain := Chain{
		parent,
	}

	return chain
}

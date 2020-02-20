package interfaces

import "github.com/denkhaus/processchain/shared"

type Chain interface {
	WithContext(ctx *shared.ModuleContext) interface{}
}

type Readable interface {
	ReadResult(shared.ReaderHandler) Proceedable
}

type Combinable interface {
	Or(or ...Combinable) Combinable
	And(or ...Combinable) Combinable
	Not(not ...Combinable) Combinable
}

type Catchable interface {
	Catch(fn shared.ErrorHandler) Executable
}

type Alternative interface {
	Catchable
	Else(fns ...shared.Handler) Catchable
}

type Optionable interface {
	WithOptions(options shared.Option) Proceedable
}
type Executable interface {
	Execute() shared.ChainState
}

type Proceedable interface {
	Then(fns ...shared.Handler) Alternative
}

type Startable interface {
	Proceedable
	If(comb Combinable) Proceedable
}

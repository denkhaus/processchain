package processchain

import "github.com/denkhaus/processchain/shared"

type Combinable interface {
	Or(or ...Combinable) Combinable
	And(or ...Combinable) Combinable
	Not(not ...Combinable) Combinable
}

type Alternative interface {
	Else(fns ...shared.Handler) Catchable
	Catch(fn shared.ErrorHandler) Executable
}

type Catchable interface {
	Catch(fn shared.ErrorHandler) Executable
}

type Executable interface {
	Execute() shared.ChainHandledState
}

type Proceedable interface {
	Then(fns ...shared.Handler) Alternative
}

type Startable interface {
	If(comb Combinable) Proceedable
}

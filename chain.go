package processchain

import (
	"github.com/denkhaus/processchain/shared"
	"github.com/juju/errors"
	"github.com/lann/builder"
)

var (
	log logrus.FieldLogger = logrus.New().WithField("package", "processchain")
)

type chain builder.Builder

type ActionData struct {
	// EntityDescriptor shared.EntityDescriptor
	// Operation        shared.Operation
	// FieldOperation   shared.Operation
	ErrorHandlers shared.ErrorHandlers
	Then          shared.Handlers
	Else          shared.Handlers
	Conditions    shared.EvalFuncs
	//	FieldName        string
	Or  []ActionData
	And []ActionData
	Not []ActionData
}

func (b chain) FromGraphQL(gql string) Proceedable {
	return builder.Set(b, "Operation", "created").(Proceedable)
}

func (b chain) Or(or ...Combinable) Combinable {
	data := []interface{}{}
	for _, o := range or {
		data = append(data, builder.GetStruct(o))
	}
	return builder.Append(b, "Or", data...).(Combinable)
}

func (b chain) And(and ...Combinable) Combinable {
	data := []interface{}{}
	for _, a := range and {
		data = append(data, builder.GetStruct(a))
	}
	return builder.Append(b, "And", data...).(Combinable)
}

func (b chain) Not(not ...Combinable) Combinable {
	data := []interface{}{}
	for _, n := range not {
		data = append(data, builder.GetStruct(n))
	}
	return builder.Append(b, "Not", data...).(Combinable)
}

func (b chain) Catch(fn shared.ErrorHandler) Executable {
	return builder.Append(b, "ErrorHandlers", fn).(Executable)
}

func (b chain) Then(fns ...shared.Handler) Alternative {
	data := []interface{}{}
	for _, fn := range fns {
		data = append(data, fn)
	}
	return builder.Append(b, "Then", data...).(Alternative)
}

func (b chain) Else(fns ...shared.Handler) Catchable {
	data := []interface{}{}
	for _, fn := range fns {
		data = append(data, fn)
	}
	return builder.Append(b, "Else", data...).(Catchable)
}

var actionChain = builder.Register(chain{}, ActionData{})

func WithOptions(options ...sharedOption) Startable {
	return comb.(Proceedable)
}

// func OnNodeCreated() Combinable {
// 	return actionChain.(Selectable).OnNodeCreated()
// }
// func OnNodeUpdated() Combinable {
// 	return actionChain.(Selectable).OnNodeUpdated()
// }
// func OnNodeDeleted() Combinable {
// 	return actionChain.(Selectable).OnNodeDeleted()
// }
// func OnFieldCreated(field string) Combinable {
// 	return actionChain.(Selectable).OnFieldCreated(field)
// }
// func OnFieldUpdated(field string) Combinable {
// 	return actionChain.(Selectable).OnFieldUpdated(field)
// }
// func OnFieldDeleted(field string) Combinable {
// 	return actionChain.(Selectable).OnFieldDeleted(field)
// }

package processchain

import (
	"github.com/denkhaus/processchain/interfaces"
	"github.com/denkhaus/processchain/shared"
	"github.com/lann/builder"
	"github.com/sirupsen/logrus"
)

var (
	log logrus.FieldLogger = logrus.New().WithField("package", "processchain")
)

type chain builder.Builder

func (b chain) Or(or ...interfaces.Combinable) interfaces.Combinable {
	data := []interface{}{}
	for _, o := range or {
		data = append(data, builder.GetStruct(o))
	}
	return builder.Append(b, "Or", data...).(interfaces.Combinable)
}

func (b chain) And(and ...interfaces.Combinable) interfaces.Combinable {
	data := []interface{}{}
	for _, a := range and {
		data = append(data, builder.GetStruct(a))
	}
	return builder.Append(b, "And", data...).(interfaces.Combinable)
}

func (b chain) Not(not ...interfaces.Combinable) interfaces.Combinable {
	data := []interface{}{}
	for _, n := range not {
		data = append(data, builder.GetStruct(n))
	}
	return builder.Append(b, "Not", data...).(interfaces.Combinable)
}

func (b chain) Catch(fn shared.ErrorHandler) interfaces.Executable {
	return builder.Set(b, "ErrorHandler", fn).(interfaces.Executable)
}

func (b chain) Then(handler shared.Handler) interfaces.Alternative {
	return builder.Append(b, "Then", handler).(interfaces.Alternative)
}

func (b chain) WithContext(ctx *shared.ModuleContext) interface{} {
	return builder.Set(b, "Context", ctx)
}

func (b chain) ReadResult(shared.ReaderHandler) interfaces.Proceedable {
	return b
}

var defaultChain = builder.Register(chain{}, ActionData{})

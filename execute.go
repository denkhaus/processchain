package processchain

import (
	"github.com/denkhaus/processchain/shared"
	"github.com/juju/errors"
	"github.com/lann/builder"
)

type ActionData struct {
	Context      *shared.ModuleContext
	ErrorHandler shared.ErrorHandler
	Then         shared.Handlers
	Conditions   shared.EvalFuncs
	Or           []ActionData
	And          []ActionData
	Not          []ActionData
}

func (b chain) handleError(err error) error {
	if h, ok := builder.Get(b, "ErrorHandler"); ok {
		h.(shared.ErrorHandler)(err)
		return nil
	}
	return nil
}

func (b chain) Execute() error {
	data := builder.GetStruct(b).(ActionData)
	if err := data.Context.Evaluate(); err != nil {
		return b.handleError(errors.Annotate(err, "Evaluate"))
	}

	result := data.Context.Result
	for _, handle := range data.Then {
		res, err := handle(result)
		if err != nil {
			return b.handleError(
				errors.Annotate(err, "HandleEvent [then]"),
			)
		}
		result = res
	}

	return nil
}

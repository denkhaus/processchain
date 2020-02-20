package processchain

import (
	"github.com/juju/errors"
	"github.com/lann/builder"
)

func (b chain) handleError(err error) {
	if ehs, ok := builder.Get(b, "ErrorHandlers"); ok {
		handlers := ehs.(shared.ErrorHandlers)
		for _, handle := range handlers {
			handle(err)
		}
	} else {
		panic(errors.Annotate(err, "no catch handler found"))
	}
}

func (b chain) Execute(ctx goka.Context, m *shared.EventContext) shared.ChainHandledState {
	data := builder.GetStruct(b).(ActionData)
	if len(data.Then) == 0 {
		b.handleError(errors.New("no handler defined"))
		return shared.ChainHandledStateThenFailed
	}

	hCtx := shared.HandlerContext{
		GokaContext:      ctx,
		EntityDescriptor: data.EntityDescriptor,
		EventContext:     m,
	}

	if data.Match(m) {
		for _, handle := range data.Then {
			if err := handle(&hCtx); err != nil {
				b.handleError(errors.Annotate(err, "HandleEvent [then]"))
				return shared.ChainHandledStateThenFailed
			}
		}

		return shared.ChainHandledStateThen
	}

	if len(data.Else) == 0 {
		return shared.ChainHandledStateUnhandled
	}

	for _, handle := range data.Else {
		if err := handle(&hCtx); err != nil {
			b.handleError(errors.Annotate(err, "HandleEvent [else]"))
			return shared.ChainHandledStateElseFailed
		}
	}

	return shared.ChainHandledStateElse
}

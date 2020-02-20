package processchain

import (
	"github.com/denkhaus/processchain/interfaces"
	"github.com/denkhaus/processchain/modules/httprequest"
	"github.com/denkhaus/processchain/shared"
)

type HttpRequestOptionable interface {
	WithOptions(opts ...httprequest.Option) HttpRequestStartable
}

type HttpRequestProceedable interface {
	Get() interfaces.Proceedable
	Delete() interfaces.Proceedable
	Post() interfaces.Proceedable
	Patch() interfaces.Proceedable
	Put() interfaces.Proceedable
}

type HttpRequestStartable interface {
	HttpRequestOptionable
	HttpRequestProceedable
}

func HttpRequest(url string) HttpRequestStartable {
	ctx := shared.NewModuleContext("httprequest").Set("url", url)
	chain := httprequest.Inherit(defaultChain.(interfaces.Chain))
	return chain.WithContext(ctx).(HttpRequestStartable)
}

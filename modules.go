package processchain

import (
	"github.com/denkhaus/processchain/interfaces"
	"github.com/denkhaus/processchain/modules/httprequest"
	"github.com/denkhaus/processchain/shared"
)

func HttpRequest(url string) httprequest.HttpRequestStartable {
	ctx := shared.NewModuleContext("httprequest").Set("url", url)
	chain := httprequest.Inherit(defaultChain.(interfaces.Chain))
	return chain.WithContext(ctx).(httprequest.HttpRequestStartable)
}

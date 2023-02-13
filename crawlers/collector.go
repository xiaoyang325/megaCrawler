package crawlers

import (
	"time"

	"github.com/gocolly/colly/v2"
)

type HTMLCallback func(element *colly.HTMLElement, ctx *Context)
type XMLCallback func(element *colly.XMLElement, ctx *Context)
type CollyHTMLPair struct {
	callback colly.HTMLCallback
	selector string
}
type XMLPair struct {
	callback XMLCallback
	selector string
}
type CollectorConstructor struct {
	parallelLimit    int
	domainGlob       string
	timeout          time.Duration
	startingURLs     []string
	robotTxt         string
	htmlHandlers     []CollyHTMLPair
	xmlHandlers      []XMLPair
	responseHandlers []func(response *colly.Response, ctx *Context)
	launchHandler    func()
}

func retryRequest(r *colly.Request, maxRetries int) int {
	retriesLeft := maxRetries
	if x, ok := r.Ctx.GetAny("retriesLeft").(int); ok {
		retriesLeft = x
	}
	if retriesLeft > 0 {
		r.Ctx.Put("retriesLeft", retriesLeft-1)
		_ = r.Retry()
	}
	return retriesLeft
}

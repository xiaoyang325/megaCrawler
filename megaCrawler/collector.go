package megaCrawler

import (
	"github.com/gocolly/colly/v2"
	"time"
)

type CollectorConstructor struct {
	parallelLimit    int
	domainGlob       string
	timeout          time.Duration
	startingUrls     []string
	robotTxt         string
	htmlHandlers     map[string]func(element *colly.HTMLElement, ctx *Context)
	xmlHandlers      map[string]func(element *colly.XMLElement, ctx *Context)
	responseHandlers []func(response *colly.Response, ctx *Context)
	launchHandler    func()
	UrlData          chan UrlData
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

package megaCrawler

import (
	"github.com/gocolly/colly/v2"
	"time"
)

type CollectorConstructor struct {
	parallelLimit int
	domainGlob    string
	timeout       time.Duration
	startingUrls  []string
	robotTxt      string
	htmlHandlers  map[string]colly.HTMLCallback
	xmlHandlers   map[string]colly.XMLCallback
	UrlData       chan UrlData
}

func (cc *CollectorConstructor) SetStartingUrls(urls []string) *CollectorConstructor {
	cc.startingUrls = urls
	return cc
}

func (cc *CollectorConstructor) FromRobotTxt(url string) *CollectorConstructor {
	cc.robotTxt = url
	return cc
}

func (cc *CollectorConstructor) SetTimeout(timeout time.Duration) *CollectorConstructor {
	cc.timeout = timeout
	return cc
}

func (cc *CollectorConstructor) SetDomain(domain string) *CollectorConstructor {
	cc.domainGlob = domain
	return cc
}

func (cc *CollectorConstructor) OnHTML(querySelector string, callback colly.HTMLCallback) *CollectorConstructor {
	cc.htmlHandlers[querySelector] = callback
	return cc
}

func (cc *CollectorConstructor) OnXML(querySelector string, callback colly.XMLCallback) *CollectorConstructor {
	cc.xmlHandlers[querySelector] = callback
	return cc
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

package megaCrawler

import (
	"github.com/gocolly/colly/v2"
	"github.com/gocolly/colly/v2/extensions"
	"time"
)

type CollectorConstructor struct {
	parallelLimit int
	domainGlob    string
	timeout       time.Duration
	htmlHandlers  map[string]colly.HTMLCallback
	xmlHandlers   map[string]colly.XMLCallback
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

func (cc *CollectorConstructor) GetCollector() (c *colly.Collector, ok error) {
	c = colly.NewCollector(
		colly.ParseHTTPErrorResponse(),
		colly.Async(true),
	)
	extensions.RandomUserAgent(c)

	err := c.Limit(&colly.LimitRule{
		RandomDelay: 5 * time.Second,
		DomainGlob:  cc.domainGlob,
		Parallelism: 16,
	})
	if err != nil {
		return nil, err
	}

	for selector, htmlCallback := range cc.htmlHandlers {
		c.OnHTML(selector, htmlCallback)
	}
	for selector, xmlCallback := range cc.xmlHandlers {
		c.OnXML(selector, xmlCallback)
	}

	c.OnError(func(r *colly.Response, err error) {
		time.Sleep(1)
		_ = retryRequest(r.Request, 10)
	})
	return
}

func retryRequest(r *colly.Request, maxRetries int) int {
	retriesLeft := maxRetries
	if x, ok := r.Ctx.GetAny("retriesLeft").(int); ok {
		retriesLeft = x
	}
	if retriesLeft > 0 {
		r.Ctx.Put("retriesLeft", retriesLeft-1)
		r.Retry()
	}
	return retriesLeft
}

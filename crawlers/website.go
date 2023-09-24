package crawlers

import (
	"encoding/json"
	"io"
	"math/rand"
	"net/http"
	"net/url"
	"strings"
	"sync"
	"time"

	"megaCrawler/crawlers/commands"
	"megaCrawler/crawlers/config"
	"megaCrawler/crawlers/tester"

	"github.com/go-co-op/gocron"
	"github.com/gocolly/colly/v2"
	"github.com/gocolly/colly/v2/extensions"
	tld "github.com/jpillora/go-tld"
	"github.com/schollz/progressbar/v3"
	"github.com/temoto/robotstxt"
)

type WebsiteEngine struct {
	ID          string
	BaseURL     tld.URL
	IsRunning   bool
	Disabled    bool
	bar         *progressbar.ProgressBar
	Scheduler   *gocron.Scheduler
	LastUpdate  time.Time
	Collector   CollectorConstructor
	Config      *config.Config
	ProgressBar string
	WG          *sync.WaitGroup
	URLChannel  chan urlData
	Test        *tester.Tester
}

type urlData struct {
	URL      *url.URL
	PageType PageType
}

func (w *WebsiteEngine) Visit(url string, pageType PageType) {
	if url == "" {
		return
	}

	u, err := w.BaseURL.Parse(url)
	if err != nil {
		return
	}
	topLevel, err := tld.Parse(u.String())
	if err != nil {
		return
	}
	if topLevel.Domain != w.BaseURL.Domain || topLevel.TLD != w.BaseURL.TLD {
		return
	}

	w.URLChannel <- urlData{URL: u, PageType: pageType}
}

func (w *WebsiteEngine) SetStartingURLs(urls []string) *WebsiteEngine {
	w.Collector.startingURLs = urls
	return w
}

func (w *WebsiteEngine) FromRobotTxt(url string) *WebsiteEngine {
	w.Collector.robotTxt = url
	return w
}

func (w *WebsiteEngine) SetTimeout(timeout time.Duration) *WebsiteEngine {
	w.Collector.timeout = timeout
	return w
}

func (w *WebsiteEngine) SetDomain(domain string) *WebsiteEngine {
	w.Collector.domainGlob = domain
	return w
}

func (w *WebsiteEngine) OnHTML(selector string, callback func(element *colly.HTMLElement, ctx *Context)) *WebsiteEngine {
	w.Collector.htmlHandlers = append(w.Collector.htmlHandlers, CollyHTMLPair{func(element *colly.HTMLElement) {
		if w.Test != nil && w.Test.Done {
			return
		}
		callback(element, element.Request.Ctx.GetAny("ctx").(*Context))
	}, selector})
	return w
}

func (w *WebsiteEngine) OnXML(selector string, callback func(element *colly.XMLElement, ctx *Context)) *WebsiteEngine {
	w.Collector.xmlHandlers = append(w.Collector.xmlHandlers, XMLPair{callback, selector})
	return w
}

func (w *WebsiteEngine) OnResponse(callback func(response *colly.Response, ctx *Context)) *WebsiteEngine {
	w.Collector.responseHandlers = append(w.Collector.responseHandlers, callback)
	return w
}

func (w *WebsiteEngine) OnLaunch(callback func()) *WebsiteEngine {
	w.Collector.launchHandler = callback
	return w
}

func (w *WebsiteEngine) getCollector() (c *colly.Collector, ok error) {
	cc := w.Collector
	c = colly.NewCollector(
		colly.Async(true),
	)
	extensions.RandomUserAgent(c)
	extensions.Referer(c)

	if Proxy != nil {
		err := c.SetProxy(Proxy.String())
		if err != nil {
			return nil, err
		}
	}

	err := c.Limit(&colly.LimitRule{
		RandomDelay: 25 * time.Second,
		DomainGlob:  cc.domainGlob,
		Parallelism: cc.parallelLimit,
	})

	c.SetRequestTimeout(cc.timeout)
	if err != nil {
		return nil, err
	}

	for _, htmlCallback := range cc.htmlHandlers {
		c.OnHTML(htmlCallback.selector, htmlCallback.callback)
	}

	for _, xmlCallback := range cc.xmlHandlers {
		c.OnXML(xmlCallback.selector, func(element *colly.XMLElement) {
			if w.Test != nil && w.Test.Done {
				return
			}
			xmlCallback.callback(element, element.Request.Ctx.GetAny("ctx").(*Context))
		})
	}

	for _, handler := range cc.responseHandlers {
		c.OnResponse(func(response *colly.Response) {
			handler(response, response.Ctx.GetAny("ctx").(*Context))
		})
	}

	c.OnError(func(r *colly.Response, err error) {
		if err.Error() == "Too many requests" {
			time.Sleep(time.Duration(rand.Intn(10)) * time.Second)
		}
		RetryRequest(r.Request, err, w)
	})
	return
}

func RetryRequest(r *colly.Request, err error, w *WebsiteEngine) {
	defer func() {
		if r := recover(); r != nil {
			Sugar.Debug(r)
		}
	}()
	if w.Test != nil && w.Test.Done {
		return
	}
	left := retryRequest(r, 10)

	if left == 0 {
		_ = w.bar.Add(1)
		w.WG.Done()
		if err != nil {
			Sugar.Errorf("Max retries exceed for %s: %s", r.URL.String(), err.Error())
		}
	} else {
		time.Sleep(time.Duration(rand.Intn(10)) * time.Second)
		if err != nil {
			Sugar.Debugf("Website error tries %d for %s: %s", left, r.URL.String(), err.Error())
		}
	}
}

func (w *WebsiteEngine) processURL() (err error) {
	defer func() {
		if r := recover(); r != nil {
			Sugar.Debug(r)
		}
	}()

	c, err := w.getCollector()
	if err != nil {
		return
	}
	w.URLChannel = make(chan urlData)

	if w.Test != nil {
		go func() {
			defer func() {
				if r := recover(); r != nil {
					Sugar.Debug(r)
				}
			}()

			for {
				time.Sleep(10 * time.Microsecond)
				if w.Test.Done {
					for {
						w.WG.Done()
					}
				}
			}
		}()
	}

	c.OnScraped(func(response *colly.Response) {
		defer func() {
			if r := recover(); r != nil {
				Sugar.Debug(r)
			}
		}()

		if w.Test != nil && w.Test.Done {
			w.WG.Done()
			return
		}
		if strings.Contains(response.Ctx.Get("title"), "Internal server error") {
			time.Sleep(10 * time.Second)
			_ = response.Request.Retry()
			return
		}
		ctx := response.Ctx.GetAny("ctx").(*Context)
		ctx.CrawlTime = time.Now()
		go func() {
			defer func() {
				if r := recover(); r != nil {
					Sugar.Debug(r)
				}
			}()

			if !ctx.process(w.Test) {
				Sugar.Debugw("Empty Page", spread(*ctx)...)
				if w.Test == nil {
					newCtx := newContext(urlData{URL: response.Request.URL, PageType: ctx.PageType}, w)
					response.Ctx.Put("ctx", &newCtx)
					RetryRequest(response.Request, nil, w)
				}
			} else {
				_ = w.bar.Add(1)
				w.WG.Done()
			}
		}()
	})

	go func() {
		for {
			k := <-w.URLChannel
			if w.Test != nil && w.Test.Done {
				return
			}
			if k.URL == nil {
				break
			}
			ctx := colly.NewContext()

			newCtx := newContext(k, w)
			ctx.Put("ctx", &newCtx)
			err := c.Request("GET", k.URL.String(), nil, ctx, nil)
			if err != nil {
				continue
			}
			if w.Test != nil && w.Test.Done {
				return
			}
			w.WG.Add(1)
			w.bar.ChangeMax64(w.bar.GetMax64() + 1)
		}
	}()

	for _, startingURL := range w.Collector.startingURLs {
		w.Visit(startingURL, Index)
	}

	if w.Collector.launchHandler != nil {
		w.WG.Add(1)

		go func() {
			w.Collector.launchHandler()
			defer func() {
				if r := recover(); r != nil {
					Sugar.Debug(r)
				}
			}()
			w.WG.Done()
		}()
	}

	if w.Collector.robotTxt != "" {
		resp, err := http.Get(w.Collector.robotTxt)
		if err != nil {
			return err
		}
		robots, err := robotstxt.FromResponse(resp)
		if err != nil {
			return err
		}
		err = resp.Body.Close()
		if err != nil {
			return err
		}
		if len(robots.Sitemaps) > 0 {
			for _, sitemap := range robots.Sitemaps {
				u, err := w.BaseURL.Parse(sitemap)
				if err != nil {
					continue
				}
				w.Visit(u.String(), Index)
			}
		}
	}

	time.Sleep(5 * time.Second)
	w.WG.Wait()
	if w.Test != nil && !w.Test.Done {
		w.Test.WG.Done()
		w.Test.Done = true
		Sugar.Infof("w.Test finished, all urls processed")
	}
	return
}

func newContext(k urlData, w *WebsiteEngine) Context {
	return Context{
		PageType:   k.PageType,
		Authors:    []string{},
		Image:      []string{},
		Video:      []string{},
		Audio:      []string{},
		File:       []string{},
		Link:       []string{},
		Tags:       []string{},
		Keywords:   []string{},
		SubContext: []*Context{},
		URL:        k.URL.String(),
		Host:       k.URL.Host,
		Website:    w.ID,
		CrawlTime:  time.Time{},
	}
}

func StartEngine(w *WebsiteEngine, test bool) {
	if !test && w.Test != nil {
		return
	}
	if w.IsRunning {
		Sugar.Info("Already running id \"" + w.ID + "\"")
		return
	}
	Sugar.Infow("Starting engine", "id", w.ID)
	w.IsRunning = true
	_ = w.bar.Set(0)
	w.bar.ChangeMax(0)
	w.bar.Reset()
	err := w.processURL()
	if err != nil {
		Sugar.Errorw("Error when processing url", "id", w.ID, "err", err)
	}
	w.IsRunning = false
	Sugar.Info("Finished engine \"" + w.ID + "\"")
}

func (w *WebsiteEngine) toStatus() (s commands.WebsiteStatus) {
	_, next := w.Scheduler.NextRun()
	return commands.WebsiteStatus{
		ID:          w.ID,
		BaseURL:     w.BaseURL.String(),
		IsRunning:   w.IsRunning,
		NextIter:    next,
		ProgressBar: w.ProgressBar,
		Bar:         w.bar.String(),
		Name:        w.Config.Name,
		IterPerSec:  w.bar.State().KBsPerSecond * 1024,
	}
}

func (w *WebsiteEngine) toJSON() (b []byte, err error) {
	k := w.toStatus()
	b, err = json.Marshal(k)
	return
}

func NewEngine(id string, baseURL tld.URL) (we *WebsiteEngine) {
	we = &WebsiteEngine{
		WG:         &sync.WaitGroup{},
		ID:         id,
		BaseURL:    baseURL,
		LastUpdate: time.Unix(0, 0),
		Collector: CollectorConstructor{
			domainGlob:    baseURL.String(),
			parallelLimit: 1,
			timeout:       10 * time.Second,
			htmlHandlers:  []CollyHTMLPair{},
			xmlHandlers:   []XMLPair{},
		},
		Scheduler: gocron.NewScheduler(time.Local),
		bar: progressbar.NewOptions(
			0,
			progressbar.OptionSetWriter(io.Discard),
			progressbar.OptionEnableColorCodes(true),
			progressbar.OptionShowIts(),
			progressbar.OptionShowCount(),
			progressbar.OptionSetDescription("[progress] scrapping the internet..."),
			progressbar.OptionSetTheme(progressbar.Theme{
				Saucer:        "[green]=[reset]",
				SaucerHead:    "[green]>[reset]",
				SaucerPadding: " ",
				BarStart:      "[",
				BarEnd:        "]",
			}),
		),
	}
	return
}

// 美国媒体：
//1.Radio Free Asia：https://www.rfa.org/
//2.Benar News：https://www.benarnews.org/
//3.The Defense Post: https://www.thedefensepost.com
//英国媒体：
//1.Reuters：https://www.reuters.com/
//美国智库：
//1.USNI News：https://news.usni.org/
//2.The Maritime Executive：https://maritime-executive.com/
//3..Navy Recognition：https://navyrecognition.com/
//日本媒体：
//1.Kyodo：https://english.kyodonews.net/
//2.The Yomiuri Shimbun：https://japannews.yomiuri.co.jp/

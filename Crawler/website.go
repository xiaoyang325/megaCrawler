package Crawler

import (
	"encoding/json"
	"github.com/go-co-op/gocron"
	"github.com/gocolly/colly/v2"
	"github.com/gocolly/colly/v2/extensions"
	tld "github.com/jpillora/go-tld"
	"github.com/schollz/progressbar/v3"
	"github.com/temoto/robotstxt"
	"io/ioutil"
	"math/rand"
	"megaCrawler/Crawler/commands"
	"megaCrawler/Crawler/config"
	"net/http"
	"net/url"
	"strings"
	"sync"
	"time"
)

type WebsiteEngine struct {
	Id           string
	BaseUrl      tld.URL
	IsRunning    bool
	Disabled     bool
	bar          *progressbar.ProgressBar
	doneLaunch   bool
	Scheduler    *gocron.Scheduler
	LastUpdate   time.Time
	UrlProcessor CollectorConstructor
	Config       *config.Config
	ProgressBar  string
	WG           *sync.WaitGroup
	UrlData      chan urlData
}

type urlData struct {
	Url      *url.URL
	PageType PageType
}

func (w *WebsiteEngine) Visit(url string, pageType PageType) {
	if url == "" {
		return
	}

	u, err := w.BaseUrl.Parse(url)
	if err != nil {
		return
	}
	topLevel, _ := tld.Parse(u.String())
	if topLevel.Domain != w.BaseUrl.Domain || topLevel.TLD != w.BaseUrl.TLD {
		return
	}

	w.UrlData <- urlData{Url: u, PageType: pageType}
}

func (w *WebsiteEngine) SetStartingUrls(urls []string) *WebsiteEngine {
	w.UrlProcessor.startingUrls = urls
	return w
}

func (w *WebsiteEngine) FromRobotTxt(url string) *WebsiteEngine {
	w.UrlProcessor.robotTxt = url
	return w
}

func (w *WebsiteEngine) SetTimeout(timeout time.Duration) *WebsiteEngine {
	w.UrlProcessor.timeout = timeout
	return w
}

func (w *WebsiteEngine) SetDomain(domain string) *WebsiteEngine {
	w.UrlProcessor.domainGlob = domain
	return w
}

func (w *WebsiteEngine) OnHTML(selector string, callback func(element *colly.HTMLElement, ctx *Context)) *WebsiteEngine {
	w.UrlProcessor.htmlHandlers = append(w.UrlProcessor.htmlHandlers, CollyHTMLPair{func(element *colly.HTMLElement) {
		if Test != nil && Test.Done {
			return
		}
		callback(element, element.Request.Ctx.GetAny("ctx").(*Context))
	}, selector})
	return w
}

func (w *WebsiteEngine) OnXML(selector string, callback func(element *colly.XMLElement, ctx *Context)) *WebsiteEngine {
	w.UrlProcessor.xmlHandlers = append(w.UrlProcessor.xmlHandlers, XMLPair{callback, selector})
	return w
}

func (w *WebsiteEngine) OnResponse(callback func(response *colly.Response, ctx *Context)) *WebsiteEngine {
	w.UrlProcessor.responseHandlers = append(w.UrlProcessor.responseHandlers, callback)
	return w
}

func (w *WebsiteEngine) OnLaunch(callback func()) *WebsiteEngine {
	w.UrlProcessor.launchHandler = callback
	return w
}

func (w *WebsiteEngine) getCollector() (c *colly.Collector, ok error) {
	cc := w.UrlProcessor
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
		RandomDelay: 5 * time.Second,
		DomainGlob:  cc.domainGlob,
		Parallelism: Threads,
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
			if Test != nil && Test.Done {
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
	left := retryRequest(r, 10)
	if Test != nil && Test.Done {
		return
	}

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

func (w *WebsiteEngine) processUrl() (err error) {
	c, err := w.getCollector()
	if err != nil {
		return
	}
	w.UrlData = make(chan urlData)

	c.OnScraped(func(response *colly.Response) {
		if Test != nil && Test.Done {
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
			if !ctx.process() {
				Sugar.Debugw("Empty Page", spread(*ctx)...)
				response.Ctx.Put("ctx", newContext(urlData{Url: response.Request.URL, PageType: ctx.PageType}, w))
				RetryRequest(response.Request, nil, w)
			} else {
				_ = w.bar.Add(1)
				w.WG.Done()
			}
		}()
	})

	go func() {
		for true {
			k := <-w.UrlData
			if Test != nil && Test.Done {
				return
			}
			if k.Url == nil {
				break
			}
			ctx := colly.NewContext()

			ctx.Put("ctx", newContext(k, w))
			err := c.Request("GET", k.Url.String(), nil, ctx, nil)
			if err != nil {
				continue
			}
			w.WG.Add(1)
			w.bar.ChangeMax64(w.bar.GetMax64() + 1)
		}
	}()

	for _, startingUrl := range w.UrlProcessor.startingUrls {
		w.Visit(startingUrl, Index)
	}

	if w.UrlProcessor.launchHandler != nil {
		go func() {
			w.UrlProcessor.launchHandler()
			w.WG.Done()
		}()

		w.WG.Add(1)
	}

	if w.UrlProcessor.robotTxt != "" {
		resp, err := http.Get(w.UrlProcessor.robotTxt)
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
				u, err := w.BaseUrl.Parse(sitemap)
				if err != nil {
					continue
				}
				w.Visit(u.String(), Index)
			}
		}
	}

	time.Sleep(5 * time.Second)
	w.WG.Wait()
	if Test != nil && !Test.Done {
		Test.WG.Done()
		Test.Done = true
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
		Url:        k.Url.String(),
		Host:       k.Url.Host,
		Website:    w.Id,
		CrawlTime:  time.Time{},
	}
}

func StartEngine(w *WebsiteEngine, test bool) {
	if !test && Test != nil {
		return
	}
	if w.IsRunning {
		Sugar.Info("Already running id \"" + w.Id + "\"")
		return
	}
	Sugar.Infow("Starting engine", "id", w.Id)
	w.IsRunning = true
	_ = w.bar.Set(0)
	w.bar.ChangeMax(0)
	w.bar.Reset()
	err := w.processUrl()
	if err != nil {
		Sugar.Errorw("Error when processing url", "id", w.Id, "err", err)
	}
	w.IsRunning = false
	Sugar.Info("Finished engine \"" + w.Id + "\"")
}

func (w *WebsiteEngine) toStatus() (s commands.WebsiteStatus) {
	_, next := w.Scheduler.NextRun()
	return commands.WebsiteStatus{
		Id:          w.Id,
		BaseUrl:     w.BaseUrl.String(),
		IsRunning:   w.IsRunning,
		NextIter:    next,
		ProgressBar: w.ProgressBar,
		Bar:         w.bar.String(),
		Name:        w.Config.Name,
		IterPerSec:  w.bar.State().KBsPerSecond * 1024,
	}
}

func (w *WebsiteEngine) toJson() (b []byte, err error) {
	k := w.toStatus()
	b, err = json.Marshal(k)
	return
}

func NewEngine(id string, baseUrl tld.URL) (we *WebsiteEngine) {
	we = &WebsiteEngine{
		WG:         &sync.WaitGroup{},
		Id:         id,
		BaseUrl:    baseUrl,
		LastUpdate: time.Unix(0, 0),
		UrlProcessor: CollectorConstructor{
			domainGlob:    baseUrl.String(),
			parallelLimit: 16,
			timeout:       10 * time.Second,
			htmlHandlers:  []CollyHTMLPair{},
			xmlHandlers:   []XMLPair{},
		},
		Scheduler: gocron.NewScheduler(time.Local),
		bar: progressbar.NewOptions(
			0,
			progressbar.OptionSetWriter(ioutil.Discard),
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

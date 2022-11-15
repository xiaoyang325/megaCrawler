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

func (w *WebsiteEngine) ApplyTemplate(template Template) *WebsiteEngine {
	w.UrlProcessor.htmlHandlers = combineSlice(w.UrlProcessor.htmlHandlers, template.htmlHandlers)
	w.UrlProcessor.xmlHandlers = combineSlice(w.UrlProcessor.xmlHandlers, template.xmlHandlers)
	return w
}

func (w *WebsiteEngine) getCollector() (c *colly.Collector, ok error) {
	cc := w.UrlProcessor
	c = colly.NewCollector(
		colly.Async(true),
	)
	extensions.RandomUserAgent(c)
	extensions.Referer(c)

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
		left := RetryRequest(r.Request, 10)

		if left == 0 {
			_ = w.bar.Add(1)
			w.WG.Done()
			Sugar.Errorf("Max retries exceed for %s: %s", r.Request.URL.String(), err.Error())
		} else {
			time.Sleep(time.Duration(rand.Intn(10)) * time.Second)
			Sugar.Debugf("Website error tries %d for %s: %s", left, r.Request.URL.String(), err.Error())
		}
	})

	if w.UrlProcessor.launchHandler != nil {
		c.OnRequest(func(request *colly.Request) {
			if w.doneLaunch {
				w.doneLaunch = true
				w.UrlProcessor.launchHandler()
			}
		})
	}
	return
}

func (w *WebsiteEngine) processUrl() (data []*Context, err error) {
	c, err := w.getCollector()
	if err != nil {
		return
	}
	w.UrlData = make(chan urlData)
	data = []*Context{}

	c.OnScraped(func(response *colly.Response) {
		if strings.Contains(response.Ctx.Get("title"), "Internal server error") {
			time.Sleep(10 * time.Second)
			_ = response.Request.Retry()
			return
		}
		_ = w.bar.Add(1)
		ctx := response.Ctx.GetAny("ctx").(*Context)
		ctx.CrawlTime = time.Now()
		go func() {
			if !ctx.process() {
				Sugar.Debugw("Empty Page", spread(*ctx)...)
				RetryRequest(response.Request, 10)
			} else {
				w.WG.Done()
			}
		}()
		data = append(data, ctx)
	})

	go func() {
		for true {
			k := <-w.UrlData
			if k.Url == nil {
				break
			}
			ctx := colly.NewContext()

			ctx.Put("ctx", &Context{
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
			})
			w.WG.Add(1)
			err := c.Request("GET", k.Url.String(), nil, ctx, nil)
			if err != nil {
				w.WG.Done()
				continue
			}
			w.bar.ChangeMax64(w.bar.GetMax64() + 1)
		}
	}()

	for _, startingUrl := range w.UrlProcessor.startingUrls {
		w.Visit(startingUrl, Index)
	}

	if w.UrlProcessor.robotTxt != "" {
		resp, err := http.Get(w.UrlProcessor.robotTxt)
		if err != nil {
			return nil, err
		}
		robots, err := robotstxt.FromResponse(resp)
		if err != nil {
			return nil, err
		}
		err = resp.Body.Close()
		if err != nil {
			return nil, err
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
	close(w.UrlData)
	return
}

func startEngine(w *WebsiteEngine) {
	if w.IsRunning {
		Sugar.Info("Already running id \"" + w.Id + "\"")
		return
	}
	Sugar.Info("Starting engine ", w.Id)
	w.IsRunning = true
	_ = w.bar.Set(0)
	w.bar.ChangeMax(0)
	w.bar.Reset()
	data, err := w.processUrl()
	if err != nil {
		Sugar.Error("Error when processing url for id \"" + w.Id + "\": " + err.Error())
	}
	Sugar.Infof("Processed %d data from \"%s\" in %s", len(data), w.Id, shortDur(time.Duration(w.bar.State().SecondsSince)*time.Second))
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

package megaCrawler

import (
	"encoding/json"
	"github.com/go-co-op/gocron"
	"github.com/gocolly/colly/v2"
	"github.com/gocolly/colly/v2/extensions"
	"github.com/schollz/progressbar/v3"
	"io/ioutil"
	"math/rand"
	"megaCrawler/megaCrawler/commandImpl"
	"megaCrawler/megaCrawler/config"
	"net/url"
	"os"
	"strings"
	"time"
)

type websiteEngine struct {
	Id           string
	BaseUrl      url.URL
	IsRunning    bool
	bar          *progressbar.ProgressBar
	Scheduler    *gocron.Scheduler
	LastUpdate   time.Time
	UrlProcessor CollectorConstructor
	Config       *config.Config
	ProgressBar  string
}

type UrlData struct {
	Url     string
	LastMod time.Time
}

type SiteInfo struct {
	Title   string
	Content string
	Author  string
	LastMod time.Time
}

func (w *websiteEngine) AddUrl(url string, lastMod time.Time) {
	w.UrlProcessor.websiteData <- UrlData{Url: url, LastMod: lastMod}
}

func (w *websiteEngine) GetCollector() (c *colly.Collector, ok error) {
	cc := w.UrlProcessor
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
		if err.Error() == "Bad Gateway" || err.Error() == "Not Found" || err.Error() == "Forbidden" {
			w.bar.Add(1)
			return
		}
		if err.Error() == "Too many requests" {
			time.Sleep(time.Duration(rand.Intn(10)) * time.Second)
		}
		left := retryRequest(r.Request, 10)
		if left == 0 {
			w.bar.Add(1)
			Logger.Errorf("Max retries exceed for %s: %s", r.Request.URL.String(), err.Error())
		}
	})
	return
}

func (w *websiteEngine) processUrl() (data []SiteInfo, err error) {
	c, err := w.GetCollector()
	w.UrlProcessor.websiteData = make(chan UrlData)
	timeMap := map[string]time.Time{}
	data = []SiteInfo{}

	c.OnScraped(func(response *colly.Response) {
		if strings.Contains(response.Ctx.Get("title"), "Internal server error") {
			time.Sleep(10 * time.Second)
			response.Request.Retry()
			return
		}
		w.bar.Add(1)
		if response.Ctx.Get("title") == "" || response.Ctx.Get("content") == "" {
			println(response.Request.URL.String(), response.Ctx.Get("title"), response.Ctx.Get("content"))
			return
		}
		timeMutex.RLock()
		k := timeMap[response.Request.URL.String()]
		timeMutex.RUnlock()
		if k.Equal(time.Unix(0, 0)) && response.Ctx.GetAny("time") != nil {
			k = response.Ctx.GetAny("time").(time.Time)
		}
		data = append(data, SiteInfo{
			Title:   standardizeSpaces(response.Ctx.Get("title")),
			Content: standardizeSpaces(response.Ctx.Get("content")),
			Author:  standardizeSpaces(response.Ctx.Get("author")),
			LastMod: k,
		})
	})

	c.OnError(func(response *colly.Response, err error) {

	})

	go func() {
		for true {
			k := <-w.UrlProcessor.websiteData

			if k.Url == "" {
				break
			}

			if w.LastUpdate.Before(k.LastMod) {
				u, err := w.BaseUrl.Parse(k.Url)
				if err != nil {
					continue
				}
				err = c.Visit(u.String())
				if err != nil {
					continue
				}
				timeMutex.Lock()
				timeMap[k.Url] = k.LastMod
				timeMutex.Unlock()
				w.bar.ChangeMax64(w.bar.GetMax64() + 1)
			}
		}
	}()

	for _, startingUrl := range w.UrlProcessor.startingUrls {
		err = c.Visit(startingUrl)
		if err != nil {
			continue
		}
	}

	c.Wait()
	close(w.UrlProcessor.websiteData)
	return
}

func StartEngine(w *websiteEngine) {
	if w.IsRunning {
		Logger.Info("Already running id \"" + w.Id + "\"")
		return
	}
	Logger.Info("Starting engine \"" + w.Id + "\"")
	w.IsRunning = true
	w.bar.Set(0)
	w.bar.ChangeMax(0)
	w.bar.Reset()
	data, err := w.processUrl()
	if err != nil {
		Logger.Error("Error when processing url for id \"" + w.Id + "\": " + err.Error())
	}
	Logger.Infof("Processed %d data from \"%s\" in %s", len(data), w.Id, shortDur(time.Duration(w.bar.State().SecondsSince)*time.Second))
	err = saveToDB(data)
	if err != nil {
		Logger.Error("Error when saving to database for id \"" + w.Id + "\": " + err.Error())
	}
	w.IsRunning = false
	Logger.Info("Finished engine \"" + w.Id + "\"")
}

func (w *websiteEngine) ToStatus() (s commandImpl.WebsiteStatus) {
	_, next := w.Scheduler.NextRun()
	return commandImpl.WebsiteStatus{
		Id:          w.Id,
		BaseUrl:     w.BaseUrl.String(),
		IsRunning:   w.IsRunning,
		NextIter:    next,
		ProgressBar: w.ProgressBar,
		Bar:         w.bar.String(),
	}
}

func (w *websiteEngine) ToJson() (b []byte, err error) {
	k := w.ToStatus()
	b, err = json.Marshal(k)
	return
}

func NewEngine(id string, baseUrl url.URL) (we *websiteEngine) {
	we = &websiteEngine{
		Id:         id,
		BaseUrl:    baseUrl,
		LastUpdate: time.Unix(0, 0),
		UrlProcessor: CollectorConstructor{
			parallelLimit: 16,
			domainGlob:    "*",
			timeout:       10 * time.Second,
			htmlHandlers:  map[string]colly.HTMLCallback{},
			xmlHandlers:   map[string]colly.XMLCallback{},
		},
		IsRunning:   false,
		ProgressBar: "",
		Scheduler:   gocron.NewScheduler(time.Local),
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

func standardizeSpaces(s string) string {
	return strings.Join(strings.Fields(s), " ")
}

func saveToDB(data []SiteInfo) (err error) {
	file, err := os.Create("./json/iiss.json")
	if os.IsNotExist(err) {
		err = os.MkdirAll("./json/", 0700)
		if err != nil {
			return err
		}
		return saveToDB(data)
	}
	decoder := json.NewEncoder(file)
	err = decoder.Encode(&data)
	if err != nil {
		return err
	}
	return nil
}

package megaCrawler

import (
	"encoding/json"
	"errors"
	"github.com/go-co-op/gocron"
	"github.com/gocolly/colly/v2"
	"megaCrawler/megaCrawler/commandImpl"
	"megaCrawler/megaCrawler/config"
	"strconv"
	"strings"
	"time"
)

type websiteEngine struct {
	Id           string
	BaseUrl      string
	IsRunning    bool
	Scheduler    *gocron.Scheduler
	UrlProcessor CollectorConstructor
	UrlGetter    CollectorConstructor
	Config       *config.Config
	ProgressBar  string
}

type urlData struct {
	url     string
	lastMod time.Time
}

type SiteInfo struct {
	Title   string
	Content string
	Author  string
	LastMod time.Time
}

func (w *websiteEngine) collectUrl() (url []urlData, err error) {
	url = []urlData{}
	c, err := w.UrlGetter.GetCollector()
	if err != nil {
		L.Info("Could not get URLGetter for " + w.Id)
		return url, errors.New("could not get URLGetter for" + "w.Id")
	}

	c.OnRequest(func(request *colly.Request) {
		request.Ctx.Put("lastMod", time.Now())
	})

	c.OnScraped(func(response *colly.Response) {
		data := urlData{
			url:     response.Ctx.Get("url"),
			lastMod: response.Ctx.GetAny("lastMod").(time.Time),
		}
		url = append(url, data)
	})

	err = c.Visit(w.BaseUrl)
	if err != nil {
		return url, errors.New("Error Visiting starting url: " + err.Error())
	}
	c.Wait()
	return
}

func (w *websiteEngine) processUrl(url []urlData) (data []SiteInfo, err error) {
	c, err := w.UrlProcessor.GetCollector()
	timeMap := map[string]time.Time{}
	data = []SiteInfo{}

	c.OnScraped(func(response *colly.Response) {
		data = append(data, SiteInfo{
			Title:   response.Ctx.Get("title"),
			Content: standardizeSpaces(response.Ctx.Get("content")),
			Author:  response.Ctx.Get("author"),
			LastMod: timeMap[response.Request.URL.String()],
		})
	})

	for _, k := range url {
		c.Visit(k.url)
		timeMap[k.url] = k.lastMod
	}
	c.Wait()
	return
}

func StartEngine(w *websiteEngine) {
	if w.IsRunning {
		L.Info("Already running id \"" + w.Id + "\"")
		return
	}
	L.Info("Starting engine \"" + w.Id + "\"")
	w.IsRunning = true
	url, err := w.collectUrl()
	if err != nil {
		L.Error("Error when collecting url for id \"" + w.Id + "\" :" + err.Error())
	}
	L.Info("Collected " + strconv.Itoa(len(url)) + " url from \"" + w.Id + "\"")
	data, err := w.processUrl(url)
	if err != nil {
		L.Error("Error when processing url for id \"" + w.Id + "\" :" + err.Error())
	}
	L.Info("Processed " + strconv.Itoa(len(url)) + " url to " + strconv.Itoa(len(data)) + " data from \"" + w.Id + "\"")
	err = saveToDB(data)
	if err != nil {
		L.Error("Error when saving to database for id \"" + w.Id + "\" :" + err.Error())
	}
	w.IsRunning = false
	L.Info("Finished engine \"" + w.Id + "\"")
}

func (w *websiteEngine) ToStatus() (s commandImpl.WebsiteStatus) {
	_, next := w.Scheduler.NextRun()
	return commandImpl.WebsiteStatus{
		Id:          w.Id,
		BaseUrl:     w.BaseUrl,
		IsRunning:   w.IsRunning,
		NextIter:    next,
		ProgressBar: w.ProgressBar,
	}
}

func (w *websiteEngine) ToJson() (b []byte, err error) {
	k := w.ToStatus()
	b, err = json.Marshal(k)
	return
}

func NewEngine(id string, baseUrl string) (we *websiteEngine) {
	we = &websiteEngine{
		Id:      id,
		BaseUrl: baseUrl,
		UrlGetter: CollectorConstructor{
			parallelLimit: 16,
			domainGlob:    "*",
			timeout:       10 * time.Second,
			htmlHandlers:  map[string]colly.HTMLCallback{},
			xmlHandlers:   map[string]colly.XMLCallback{},
		},
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
	}
	return
}

func standardizeSpaces(s string) string {
	return strings.Join(strings.Fields(s), " ")
}

func saveToDB(data []SiteInfo) (err error) {
	return nil
}

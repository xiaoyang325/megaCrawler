package hudson

import (
	"github.com/gocolly/colly/v2"
	"megaCrawler/megaCrawler"
	"time"
)

func init() {
	s := megaCrawler.Register("hudson", "https://www.hudson.org/")
	s.UrlProcessor.OnXML("//urlset/url", func(e *colly.XMLElement) {
		s.AddUrl(e.ChildText("loc"), time.Now())
	}).SetStartingUrls([]string{"https://www.hudson.org/sitemap.xml"}).SetTimeout(20 * time.Second)

	s.UrlProcessor.OnHTML("meta[property=\"og:title\"]", func(element *colly.HTMLElement) {
		element.Request.Ctx.Put("title", element.Attr("content"))
	})

	s.UrlProcessor.OnHTML(".title", func(element *colly.HTMLElement) {
		element.Request.Ctx.Put("title", element.Text)
	})

	s.UrlProcessor.OnHTML(".article-body > p", func(element *colly.HTMLElement) {
		script := element.ChildText("script")
		if script != "" {
			return
		}
		str := element.Request.Ctx.Get("content")
		element.Request.Ctx.Put("content", str+"\n"+element.Text)
	})

	s.UrlProcessor.OnHTML(".publication-meta > time", func(element *colly.HTMLElement) {
		datetime := element.Attr("datetime")
		t, _ := time.Parse("2006-01-02", datetime)
		element.Request.Ctx.Put("time", t)
	})

	s.UrlProcessor.OnHTML(".author", func(element *colly.HTMLElement) {
		element.Request.Ctx.Put("author", element.Text)
	})
}

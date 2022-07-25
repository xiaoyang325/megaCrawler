package hudson

import (
	"github.com/gocolly/colly/v2"
	"megaCrawler/megaCrawler"
	"time"
)

func init() {
	s := megaCrawler.Register("hudson", "哈德森研究所", "https://www.hudson.org/").
		SetStartingUrls([]string{"https://www.hudson.org/sitemap.xml"}).
		SetTimeout(20 * time.Second)

	s.OnXML("//urlset/url", func(e *colly.XMLElement) {
		s.AddUrl(e.ChildText("loc"), time.Now())
	})

	s.OnHTML("meta[property=\"og:title\"]", func(element *colly.HTMLElement) {
		element.Request.Ctx.Put("title", element.Attr("content"))
	})

	s.OnHTML(".title", func(element *colly.HTMLElement) {
		element.Request.Ctx.Put("title", element.Text)
	})

	s.OnHTML(".article-body > p", func(element *colly.HTMLElement) {
		script := element.ChildText("script")
		if script != "" {
			return
		}
		str := element.Request.Ctx.Get("content")
		element.Request.Ctx.Put("content", str+"\n"+element.Text)
	})

	s.OnHTML(".publication-meta > time", func(element *colly.HTMLElement) {
		datetime := element.Attr("datetime")
		t, _ := time.Parse("2006-01-02", datetime)
		element.Request.Ctx.Put("time", t)
	})

	s.OnHTML(".author", func(element *colly.HTMLElement) {
		element.Request.Ctx.Put("author", element.Text)
	})
}

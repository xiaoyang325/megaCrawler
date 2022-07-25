package cnas

import (
	"github.com/gocolly/colly/v2"
	"megaCrawler/megaCrawler"
	"time"
)

func init() {
	s := megaCrawler.Register("cnas", "新美国安全中心", "https://www.cnas.org/").
		SetStartingUrls([]string{"https://www.cnas.org/sitemap.xml"}).
		SetTimeout(20 * time.Second)

	s.OnXML("//urlset/url", func(e *colly.XMLElement) {
		s.AddUrl(e.ChildText("loc"), time.Now())
	})

	s.OnXML("//sitemapindex/sitemap", func(e *colly.XMLElement) {
		s.AddUrl(e.Text, time.Now())
	})

	s.OnHTML("meta[property=\"og:title\"]", func(element *colly.HTMLElement) {
		element.Request.Ctx.Put("title", element.Attr("content"))
	})

	s.OnHTML(".page-title", func(element *colly.HTMLElement) {
		element.Request.Ctx.Put("title", element.Text)
	})

	s.OnHTML("div[id=\"mainbar\"]", func(element *colly.HTMLElement) {
		megaCrawler.SetContent(element, element.Text)
	})

	s.OnHTML(".margin-vertical", func(element *colly.HTMLElement) {
		megaCrawler.SetContent(element, element.Text)
	})

	s.OnHTML(".margin-bottom-1em", func(element *colly.HTMLElement) {
		t, _ := time.Parse("January 2, 2006", element.Text)
		element.Request.Ctx.Put("time", t)
	})

	s.OnHTML(".contributor", func(element *colly.HTMLElement) {
		element.Request.Ctx.Put("author", element.Text)
	})
}

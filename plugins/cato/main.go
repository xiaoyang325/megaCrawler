// Package cato should work functionally, but the website uses method to stop request scrapping
package cato

import (
	"github.com/gocolly/colly/v2"
	"megaCrawler/megaCrawler"
	"time"
)

func init() {
	s := megaCrawler.Register("cato", "卡托研究所", "https://www.cato.org/").
		SetStartingUrls([]string{"https://www.cato.org/sitemap.xml"}).
		SetTimeout(20 * time.Second)

	s.OnXML("//sitemapindex/sitemap", func(e *colly.XMLElement) {
		t, _ := time.Parse("2006-01-02T15:04:05-07:00", e.ChildText("lastmod"))
		s.AddUrl(e.ChildText("loc"), t)
	})

	s.OnXML("//urlset/url", func(e *colly.XMLElement) {
		t, _ := time.Parse("2006-01-02T15:04:05-07:00", e.ChildText("lastmod"))
		s.AddUrl(e.ChildText("loc"), t)
	})

	s.OnHTML("meta[property=\"og:title\"]", func(element *colly.HTMLElement) {
		megaCrawler.SetTitle(element, element.Attr("content"))
	})

	s.OnHTML(".body-text", func(element *colly.HTMLElement) {
		megaCrawler.SetContent(element, element.Text)
	})

	s.OnHTML(".authors", func(element *colly.HTMLElement) {
		megaCrawler.AppendAuthor(element, element.Text)
	})
}

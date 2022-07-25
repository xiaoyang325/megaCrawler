package cfr

import (
	"github.com/gocolly/colly/v2"
	"megaCrawler/megaCrawler"
	"time"
)

func init() {
	s := megaCrawler.Register("cfr", "美国外交关系协会", "https://www.cfr.org/").
		SetStartingUrls([]string{"https://www.cfr.org/sitemap.xml"}).
		SetTimeout(20 * time.Second)

	s.OnXML("//urlset/url", func(e *colly.XMLElement) {
		t, _ := time.Parse("2006-01-02T15:04:05-07:00", e.ChildText("lastmod"))
		s.AddUrl(e.ChildText("loc"), t)
	})

	s.OnXML("//sitemapindex/sitemap", func(e *colly.XMLElement) {
		t, _ := time.Parse("2006-01-02T15:04:05-07:00", e.ChildText("lastmod"))
		s.AddUrl(e.ChildText("loc"), t)
	})

	s.OnHTML("meta[property=\"og:title\"]", func(element *colly.HTMLElement) {
		element.Request.Ctx.Put("title", element.Attr("content"))
	})

	s.OnHTML(".body-content", func(element *colly.HTMLElement) {
		if element.Request.Ctx.Get("content") == "" {
			element.Request.Ctx.Put("content", element.Text)
		}
	})

	s.OnHTML(".podcast-body", func(element *colly.HTMLElement) {
		if element.Request.Ctx.Get("content") == "" {
			element.Request.Ctx.Put("content", element.Text)
		}
	})

	s.OnHTML(".article-header__link", func(element *colly.HTMLElement) {
		element.Request.Ctx.Put("author", element.Text)
	})
}

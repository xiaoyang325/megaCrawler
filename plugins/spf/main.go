package spf

import (
	"github.com/gocolly/colly/v2"
	"megaCrawler/megaCrawler"
	"time"
)

func init() {
	s := megaCrawler.Register("spf", "筱川和平基金会", "https://www.spf.org/").
		SetStartingUrls([]string{"https://www.spf.org/sitemap.xml"}).
		SetTimeout(20 * time.Second)

	s.OnXML("//urlset/url", func(e *colly.XMLElement) {
		t, err := time.Parse("2006-01-02T15:04:05-07:00Z", e.ChildText("lastmod"))
		if err != nil {
			t = time.Now()
		}
		s.AddUrl(e.ChildText("loc"), t)
	})

	s.OnXML("//sitemapindex/sitemap", func(e *colly.XMLElement) {
		t, err := time.Parse("2006-01-02T15:04:05-07:00Z", e.ChildText("lastmod"))
		if err != nil {
			t = time.Now()
		}
		s.AddUrl(e.ChildText("loc"), t)
	})

	s.OnHTML("meta[property=\"og:title\"]", func(element *colly.HTMLElement) {
		megaCrawler.SetTitle(element, element.Attr("content"))
	})

	s.OnHTML(".bTxt", func(element *colly.HTMLElement) {
		megaCrawler.AppendContent(element, element.Text)
	})

	s.OnHTML(".block", func(element *colly.HTMLElement) {
		megaCrawler.AppendContent(element, element.Text)
	})
}

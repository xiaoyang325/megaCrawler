package nrl

import (
	"github.com/gocolly/colly/v2"
	"megaCrawler/megaCrawler"
	"time"
)

func init() {
	s := megaCrawler.Register("nrl", "海军研究所", "https://www.nrl.navy.mil/").
		SetStartingUrls([]string{"https://www.nrl.navy.mil/SiteMap.aspx"}).
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
}

package cimsec

import (
	"github.com/gocolly/colly/v2"
	"megaCrawler/megaCrawler"
	"time"
)

func init() {
	s := megaCrawler.Register("cimsec", "国际海洋安全研究中心", "https://cimsec.org/").
		SetStartingUrls([]string{"https://cimsec.org/wp-sitemap.xml"}).
		SetTimeout(20 * time.Second)

	s.OnXML("//urlset/url", func(e *colly.XMLElement) {
		s.AddUrl(e.ChildText("loc"), time.Now())
	})

	s.OnXML("//sitemapindex/sitemap", func(e *colly.XMLElement) {
		s.AddUrl(e.ChildText("loc"), time.Now())
	})

	s.OnHTML(".entry-title", func(element *colly.HTMLElement) {
		megaCrawler.SetTitle(element, element.Text)
	})

	s.OnHTML(".entry-content", func(element *colly.HTMLElement) {
		megaCrawler.SetContent(element, element.Text)
	})

	s.OnHTML(".author", func(element *colly.HTMLElement) {
		megaCrawler.AppendAuthor(element, element.Text)
	})

	s.OnHTML(".entry-date", func(element *colly.HTMLElement) {
		t, err := time.Parse("2006-01-02T15:04:05-07:00", element.Attr("datetime"))
		if err == nil {
			t = time.Now()
		}
		megaCrawler.SetTime(element, t)
	})
}

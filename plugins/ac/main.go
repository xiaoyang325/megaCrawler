package ac

import (
	"github.com/gocolly/colly/v2"
	"megaCrawler/megaCrawler"
	"time"
)

func init() {
	s := megaCrawler.Register("ac", "大西洋理事会", "https://www.atlanticcouncil.org/").
		SetStartingUrls([]string{"https://www.atlanticcouncil.org/sitemap_index.xml"}).
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
		megaCrawler.SetTitle(element, element.Attr("content"))
	})

	s.OnHTML(".ac-single-post--content", func(element *colly.HTMLElement) {
		megaCrawler.SetContent(element, element.Text)
	})

	s.OnHTML(".gta-embed--tax--expert", func(element *colly.HTMLElement) {
		megaCrawler.SetAuthor(element, element.Text)
	})
}

package tp

import (
	"github.com/gocolly/colly/v2"
	"megaCrawler/megaCrawler"
	"time"
)

func init() {
	s := megaCrawler.Register("tp", "透明国际", "https://www.transparency.org").
		SetStartingUrls([]string{"https://www.transparency.org/en/sitemaps-1-sitemap.xml"}).
		SetTimeout(20 * time.Second)

	s.OnXML("//urlset/url", func(e *colly.XMLElement) {
		k, err := time.Parse("2006-01-02T15:04:05-07:00", e.ChildText("lastmod"))
		if err != nil {
			_ = megaCrawler.Logger.Errorf("Error when parsing %s to time: %s", e.ChildText("lastmod"), err.Error())
			k = time.Now()
		}
		s.AddUrl(e.ChildText("loc"), k)
	})

	s.OnXML("//sitemapindex/sitemap", func(e *colly.XMLElement) {
		k, err := time.Parse("2006-01-02T15:04:05-07:00", e.ChildText("lastmod"))
		if err != nil {
			_ = megaCrawler.Logger.Errorf("Error when parsing %s to time: %s", e.ChildText("lastmod"), err.Error())
			k = time.Now()
		}
		s.AddUrl(e.ChildText("loc"), k)
	})

	s.OnHTML("meta[property=\"og:title\"]", func(element *colly.HTMLElement) {
		megaCrawler.SetTitle(element, element.Attr("content"))
	})

	s.OnHTML("span.block.text-sm.font-bold", func(element *colly.HTMLElement) {
		megaCrawler.AppendAuthor(element, element.Text)
	})

	s.OnHTML("#content > article > div.container > div", func(element *colly.HTMLElement) {
		megaCrawler.AppendContent(element, element.Text)
	})
}

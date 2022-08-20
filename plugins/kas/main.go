package kas

import (
	"github.com/gocolly/colly/v2"
	"megaCrawler/megaCrawler"
	"time"
)

func init() {
	s := megaCrawler.Register("kas", "康拉德·阿登纳基金会", "https://www.kas.de/").
		SetStartingUrls([]string{"https://www.kas.de/sitemap.xml"}).
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
		s.AddUrl(e.ChildText("loc"), time.Now())
	})

	s.OnHTML("meta[property=\"og:title\"]", func(element *colly.HTMLElement) {
		megaCrawler.SetTitle(element, element.Attr("content"))
	})

	s.OnHTML(".c-page-main__text > p", func(element *colly.HTMLElement) {
		megaCrawler.AppendContent(element, element.Text)
	})
}

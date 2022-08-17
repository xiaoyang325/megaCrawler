package chh

import (
	"github.com/gocolly/colly/v2"
	"megaCrawler/megaCrawler"
	"time"
)

func init() {
	s := megaCrawler.Register("chh", "皇家国际事务研究所", "https://www.chathamhouse.org/").
		SetStartingUrls([]string{"https://www.chathamhouse.org/sitemap.xml"}).
		SetTimeout(20 * time.Second)

	s.OnXML("//urlset/url", func(e *colly.XMLElement) {
		t, err := time.Parse("2006-01-02T15:04:05-07:00", e.ChildText("lastmod"))
		if err != nil {
			t = time.Now()
			_ = megaCrawler.Logger.Errorf("Error when parsing %s to time: %s", e.ChildText("lastmod"), err.Error())
		}
		s.AddUrl(e.ChildText("loc"), t)
	})

	s.OnXML("//sitemapindex/sitemap", func(e *colly.XMLElement) {
		t, err := time.Parse("2006-01-02T15:04:05-07:00", e.ChildText("lastmod"))
		if err != nil {
			t = time.Now()
			_ = megaCrawler.Logger.Errorf("Error when parsing %s to time: %s", e.ChildText("lastmod"), err.Error())
		}
		s.AddUrl(e.ChildText("loc"), t)
	})

	s.OnHTML("meta[property=\"og:title\"]", func(element *colly.HTMLElement) {
		megaCrawler.SetTitle(element, element.Attr("content"))
	})

	s.OnHTML(".wysiwyg", func(element *colly.HTMLElement) {
		megaCrawler.AppendContent(element, element.Text)
	})

	s.OnHTML(".person-teaser__title > a > span", func(element *colly.HTMLElement) {
		megaCrawler.AppendAuthor(element, element.Text)
	})

}

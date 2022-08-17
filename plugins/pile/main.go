package pile

import (
	"github.com/gocolly/colly/v2"
	"megaCrawler/megaCrawler"
	"time"
)

func init() {
	s := megaCrawler.Register("pile", "彼得森国际经济研究所", "https://www.piie.com").
		SetStartingUrls([]string{"https://www.piie.com/sitemap.xml"}).
		SetTimeout(20 * time.Second)

	s.OnXML("//urlset/url", func(e *colly.XMLElement) {
		t, err := time.Parse("2006-01-02T15:04Z", e.ChildText("lastmod"))
		if err != nil {
			t = time.Now()
			println(err)
		}
		s.AddUrl(e.ChildText("loc"), t)
	})

	s.OnXML("//sitemapindex/sitemap", func(e *colly.XMLElement) {
		t, err := time.Parse("2006-01-02T15:04Z", e.ChildText("lastmod"))
		if err != nil {
			t = time.Now()
		}
		s.AddUrl(e.ChildText("loc"), t)
	})

	s.OnHTML("meta[property=\"og:title\"]", func(element *colly.HTMLElement) {
		megaCrawler.SetTitle(element, element.Attr("content"))
	})

	s.OnHTML(".field--body", func(element *colly.HTMLElement) {
		megaCrawler.SetContent(element, element.Text)
	})

	s.OnHTML(".field--contributor", func(element *colly.HTMLElement) {
		megaCrawler.AppendAuthor(element, element.Text)
	})

}

package ecfr

import (
	"github.com/gocolly/colly/v2"
	"megaCrawler/megaCrawler"
	"time"
)

func init() {
	s := megaCrawler.Register("ecfr", "欧洲外交关系委员会", "https://ecfr.eu/").
		SetStartingUrls([]string{"https://ecfr.eu/wp-sitemap.xml"}).
		SetTimeout(20 * time.Second)

	s.OnXML("//urlset/url", func(e *colly.XMLElement) {
		s.AddUrl(e.ChildText("loc"), time.Now())
	})

	s.OnXML("//sitemapindex/sitemap", func(e *colly.XMLElement) {
		s.AddUrl(e.ChildText("loc"), time.Now())
	})

	s.OnHTML("meta[property=\"og:title\"]", func(element *colly.HTMLElement) {
		megaCrawler.SetTitle(element, element.Attr("content"))
	})

	s.OnHTML(".card-main-link", func(element *colly.HTMLElement) {
		megaCrawler.AppendAuthor(element, element.Text)
	})

	s.OnHTML(".post-time", func(element *colly.HTMLElement) {
		t, err := time.Parse("2006-01-02 15:04:05", element.Attr("datetime"))
		if err != nil {
			t = time.Now()
			_ = megaCrawler.Logger.Errorf("Error when parsing %s to time: %s", element.Attr("datetime"), err.Error())
		}
		megaCrawler.SetTime(element, t)
	})

	s.OnHTML(".text-body", func(element *colly.HTMLElement) {
		megaCrawler.AppendContent(element, element.Text)
	})
}

package rusi

import (
	"github.com/gocolly/colly/v2"
	"megaCrawler/megaCrawler"
	"time"
)

func init() {
	s := megaCrawler.Register("rusi", "皇家军种国防研究所", "https://www.rusi.org/").
		SetStartingUrls([]string{"https://www.rusi.org/sitemap/sitemap-index.xml"}).
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

	s.OnHTML(".ArticleLayout-module--main--1vR5r > div", func(element *colly.HTMLElement) {
		megaCrawler.AppendContent(element, element.Text)
	})

	s.OnHTML(".Article-module--mainBody--3Ueil", func(element *colly.HTMLElement) {
		megaCrawler.AppendContent(element, element.Text)
	})

	s.OnHTML("span[aria-label=\"authored by\"]", func(element *colly.HTMLElement) {
		megaCrawler.AppendAuthor(element, element.Text)
	})

	s.OnHTML("span[aria-label=\"published date\"]", func(element *colly.HTMLElement) {
		t, err := time.Parse("2 January 2006", megaCrawler.StandardizeSpaces(element.Text))
		if err != nil {
			t = time.Now()
			_ = megaCrawler.Logger.Errorf("Error when parsing %s to time: %s", element.Text, err.Error())
		}
		megaCrawler.SetTime(element, t)
	})
}

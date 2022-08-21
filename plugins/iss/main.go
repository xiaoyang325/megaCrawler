package iss

import (
	"github.com/gocolly/colly/v2"
	"megaCrawler/megaCrawler"
	"time"
)

func init() {
	s := megaCrawler.Register("iss", "欧盟安全研究所", "https://www.iss.europa.eu/").
		SetStartingUrls([]string{"https://www.iss.europa.eu/publications"}).
		SetTimeout(20 * time.Second)

	s.OnHTML(".medium-teaser-right-side", func(element *colly.HTMLElement) {
		k, err := time.Parse("2006-01-02T15:04:05-07:00", element.ChildAttr(".date-display-single", "content"))
		if err != nil {
			_ = megaCrawler.Logger.Errorf("Error when parsing %s to time: %s", element.ChildAttr(".date-display-single", "content"), err.Error())
			k = time.Now()
		}
		s.AddUrl(element.ChildAttr(".field-name-title > span > a", "href"), k)
	})

	s.OnHTML(".pagination > li > a", func(element *colly.HTMLElement) {
		s.AddUrl(element.Attr("href"), time.Now())
	})

	s.OnHTML(".body", func(element *colly.HTMLElement) {
		megaCrawler.SetContent(element, element.Text)
	})

	s.OnHTML(".node-body", func(element *colly.HTMLElement) {
		megaCrawler.SetContent(element, element.Text)
	})

	s.OnHTML("ul.links > li > a", func(element *colly.HTMLElement) {
		megaCrawler.AppendAuthor(element, element.Text)
	})

	s.OnHTML("h1", func(element *colly.HTMLElement) {
		megaCrawler.SetTitle(element, element.Text)
	})
}

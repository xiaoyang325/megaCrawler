package diis

import (
	"github.com/gocolly/colly/v2"
	"megaCrawler/megaCrawler"
	"time"
)

func init() {
	s := megaCrawler.Register("diis", "国际问题研究所", "https://www.diis.dk/en").
		SetStartingUrls([]string{"https://www.diis.dk/en/research"}).
		SetTimeout(20 * time.Second)

	s.OnHTML("div.beta > div > a", func(element *colly.HTMLElement) {
		s.AddUrl(element.Attr("href"), time.Now())
	})

	s.OnHTML(".pager__item", func(element *colly.HTMLElement) {
		s.AddUrl(element.Attr("href"), time.Now())
	})

	s.OnHTML(".field-content", func(element *colly.HTMLElement) {
		megaCrawler.SetContent(element, element.Text)
	})

	s.OnHTML("time", func(element *colly.HTMLElement) {
		k, err := time.Parse("2006-01-02T15:04:05Z", element.Attr("datetime"))
		if err != nil {
			_ = megaCrawler.Logger.Errorf("Error when parsing %s to time: %s", element.Attr("datetime"), err.Error())
			k = time.Now()
		}
		megaCrawler.SetTime(element, k)
	})

	s.OnHTML("div.field.node-title > h2 > a", func(element *colly.HTMLElement) {
		megaCrawler.AppendAuthor(element, element.Text)
	})

	s.OnHTML("h1", func(element *colly.HTMLElement) {
		megaCrawler.SetTitle(element, element.Text)
	})
}

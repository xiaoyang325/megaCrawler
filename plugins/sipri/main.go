package sipri

import (
	"github.com/gocolly/colly/v2"
	"megaCrawler/megaCrawler"
	"time"
)

func init() {
	s := megaCrawler.Register("sipri", "斯德哥尔摩国际和平研究所", "https://www.sipri.org/").
		SetStartingUrls([]string{"https://www.sipri.org/publications/search"}).
		SetTimeout(20 * time.Second)

	s.OnHTML(".sticky-enabled > tbody > tr", func(element *colly.HTMLElement) {
		k, err := time.Parse("2006 - January", element.ChildText(".views-field-field-year-of-publication"))
		if err != nil {
			k, err = time.Parse("2006", element.ChildText(".views-field-field-year-of-publication"))
			if err != nil {
				_ = megaCrawler.Logger.Errorf("Error when parsing %s to time: %s", element.ChildText(".views-field-field-year-of-publication"), err.Error())
				k = time.Now()
			}
		}
		s.AddUrl(element.ChildAttr(".views-field-title > em > a", "href"), k)
	})

	s.OnHTML(".pager__item > a", func(element *colly.HTMLElement) {
		s.AddUrl("/publications/search"+element.Attr("href"), time.Now())
	})

	s.OnHTML(".body", func(element *colly.HTMLElement) {
		megaCrawler.SetContent(element, element.Text)
	})

	s.OnHTML(".views-field-combinedauthors > span > a", func(element *colly.HTMLElement) {
		megaCrawler.AppendAuthor(element, element.Text)
	})

	s.OnHTML("h1", func(element *colly.HTMLElement) {
		megaCrawler.SetTitle(element, element.Text)
	})
}

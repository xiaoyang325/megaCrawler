package pism

import (
	"github.com/gocolly/colly/v2"
	"megaCrawler/megaCrawler"
	"time"
)

func init() {
	s := megaCrawler.Register("pism", "国际事务研究所", "https://www.pism.pl/").
		SetStartingUrls([]string{"https://www.pism.pl/publications"}).
		SetTimeout(20 * time.Second)

	s.OnHTML(".article-preview", func(element *colly.HTMLElement) {
		date, err := time.Parse("04.02.2006", element.ChildText(".article-date"))
		if err != nil {
			_ = megaCrawler.Logger.Errorf("Error when parsing %s to time: %s", element.ChildText(".c-list-item__date"), err.Error())
			date = time.Now()
		}
		s.AddUrl(element.ChildAttr(".article-title > a", "href"), date)
	})

	s.OnHTML("meta[property=\"og:title\"]", func(element *colly.HTMLElement) {
		megaCrawler.SetTitle(element, element.Attr("content"))
	})

	s.OnHTML(".pagination > li > a", func(element *colly.HTMLElement) {
		s.AddUrl(element.Attr("href"), time.Now())
	})

	s.OnHTML(".author > a", func(element *colly.HTMLElement) {
		megaCrawler.AppendAuthor(element, element.Text)
	})

	s.OnHTML(".richtext-preview", func(element *colly.HTMLElement) {
		megaCrawler.AppendContent(element, element.Text)
	})
}

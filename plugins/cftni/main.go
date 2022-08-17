package cftni

import (
	"github.com/gocolly/colly/v2"
	"megaCrawler/megaCrawler"
	"time"
)

func init() {
	s := megaCrawler.Register("cftni", "美国国家利益中心", "https://www.cftni.org/").
		SetStartingUrls([]string{"https://cftni.org/category/publications/"})

	s.OnHTML(".holder > h2 > a", func(element *colly.HTMLElement) {
		s.AddUrl(element.Attr("href"), time.Now())
	})

	s.OnHTML(".wp-paginate > li > a", func(element *colly.HTMLElement) {
		s.AddUrl(element.Attr("href"), time.Now())
	})

	s.OnHTML(".content-holder", func(element *colly.HTMLElement) {
		megaCrawler.AppendContent(element, element.ChildText("p"))
	})

}

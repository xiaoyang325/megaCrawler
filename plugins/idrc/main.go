package idrc

import (
	"github.com/gocolly/colly/v2"
	"megaCrawler/megaCrawler"
	"time"
)

func init() {
	s := megaCrawler.Register("idrc", "加拿大国际发展研究中心", "https://www.idrc.ca//").
		SetStartingUrls([]string{"https://www.idrc.ca/en/search?search_api_fulltext=&facets_query="})

	s.OnHTML(".search-result-item >h2 > a", func(element *colly.HTMLElement) {
		megaCrawler.SetTitle(element, element.Text)
		s.AddUrl(element.Attr("href"), time.Now())
	})

	s.OnHTML(".page-link", func(element *colly.HTMLElement) {
		s.AddUrl(element.Attr("href"), time.Now())
	})

	s.OnHTML("h1", func(element *colly.HTMLElement) {
		megaCrawler.SetTitle(element, element.Text)
	})

	s.OnHTML(".paragraph", func(element *colly.HTMLElement) {
		megaCrawler.AppendContent(element, element.ChildText("p"))
	})

	s.OnHTML(".field--name-published-date > span", func(element *colly.HTMLElement) {
		t, err := time.Parse("January 2, 2006", element.Text)
		if err != nil {
			t = time.Now()
			return
		}
		megaCrawler.SetTime(element, t)
	})
}

package iai

import (
	"github.com/gocolly/colly/v2"
	"megaCrawler/megaCrawler"
	"time"
)

func init() {
	s := megaCrawler.Register("iai", "国际事务研究所", "https://www.iai.it/").
		SetStartingUrls([]string{"https://www.iai.it/en/pubblicazioni/lista/all/all"}).
		SetTimeout(20 * time.Second)

	s.OnHTML(".riga-p", func(element *colly.HTMLElement) {
		t, err := time.Parse("02/01/2006", element.ChildText(".date-display-single"))
		if err != nil {
			_ = megaCrawler.Logger.Errorf("Error when parsing %s to time: %s", element.ChildText(".c-list-item__date"), err.Error())
			t = time.Now()
		}
		s.AddUrl(element.ChildAttr(".tit > a", "href"), t)
	})

	s.OnHTML(".pagination > li > a", func(element *colly.HTMLElement) {
		s.AddUrl(element.Attr("href"), time.Now())
	})

	s.OnHTML("meta[property=\"og:title\"]", func(element *colly.HTMLElement) {
		megaCrawler.SetTitle(element, element.Attr("content"))
	})

	s.OnHTML(".field-pub-autori > a", func(element *colly.HTMLElement) {
		megaCrawler.AppendAuthor(element, element.Text)
	})

	s.OnHTML(".field-body", func(element *colly.HTMLElement) {
		megaCrawler.AppendContent(element, element.Text)
	})
}

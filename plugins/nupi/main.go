package nupi

import (
	"github.com/gocolly/colly/v2"
	"megaCrawler/megaCrawler"
	"time"
)

func init() {
	s := megaCrawler.Register("nupi", "国际事务研究所", "https://www.nupi.no/").
		SetStartingUrls([]string{"https://www.nupi.no/en/publications/cristin-pub"}).
		SetTimeout(20 * time.Second)

	s.OnHTML(".media-body", func(element *colly.HTMLElement) {
		s.AddUrl(element.ChildAttr(".font-weight-bold > a", "href"), time.Now())
	})

	s.OnHTML(".pagination > li > a", func(element *colly.HTMLElement) {
		s.AddUrl(element.Attr("href"), time.Now())
	})

	s.OnHTML(".pub-date", func(element *colly.HTMLElement) {
		k, err := time.Parse("January 2, 2006", element.Text)
		if err != nil {
			_ = megaCrawler.Logger.Errorf("Error when parsing %s to time: %s", element.Text, err.Error())
			k = time.Now()
		}
		megaCrawler.SetTime(element, k)
	})

	s.OnHTML("article", func(element *colly.HTMLElement) {
		megaCrawler.SetContent(element, element.Text)
	})

	s.OnHTML("meta[property=\"og:title\"]", func(element *colly.HTMLElement) {
		megaCrawler.SetTitle(element, element.Attr("content"))
	})
}

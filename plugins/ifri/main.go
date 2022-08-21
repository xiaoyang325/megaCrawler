package ifri

import (
	"github.com/gocolly/colly/v2"
	"megaCrawler/megaCrawler"
	"time"
)

func init() {
	s := megaCrawler.Register("ifri", "国际关系研究所", "https://www.ifri.org/en").
		SetStartingUrls([]string{"https://www.ifri.org/fr/publications?page=1"}).
		SetTimeout(20 * time.Second)

	s.OnHTML(".col-md-8", func(element *colly.HTMLElement) {
		q := element.Request.URL.Query()
		if !q.Has("page") {
			return
		}
		k, err := time.Parse("02/04/2006", element.ChildText(".date-vignette"))
		if err != nil {
			_ = megaCrawler.Logger.Errorf("Error when parsing %s to time: %s", element.ChildText(".date-vignette > span"), err.Error())
			k = time.Now()
		}
		s.AddUrl(element.ChildAttr(".title-vignette-search > a", "href"), k)
	})

	s.OnHTML(".pagination > li > a", func(element *colly.HTMLElement) {
		s.AddUrl(element.Attr("href"), time.Now())
	})

	s.OnHTML(".field", func(element *colly.HTMLElement) {
		megaCrawler.SetContent(element, element.Text)
	})

	s.OnHTML("meta[property=\"og:title\"]", func(element *colly.HTMLElement) {
		megaCrawler.SetTitle(element, element.Attr("content"))
	})
}

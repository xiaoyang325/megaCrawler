package csis

import (
	"github.com/gocolly/colly/v2"
	"megaCrawler/megaCrawler"
	"time"
)

func init() {
	s := megaCrawler.Register("csis", "https://www.csis.org/").
		SetStartingUrls([]string{"https://www.csis.org/analysis"}).
		SetTimeout(20 * time.Second)

	s.OnHTML(".ds-right", func(element *colly.HTMLElement) {
		t, err := time.Parse("2006-01-02T15:04:05-07:00", element.ChildAttr(".date-display-single", "content"))
		if megaCrawler.Debug && err != nil {
			_ = megaCrawler.Logger.Error("Could not parse time" + err.Error())
			t = time.Now()
		}
		s.AddUrl(element.ChildAttr(".teaser__title > a", "href"), t)
	})

	s.OnHTML("meta[property=\"og:title\"]", func(element *colly.HTMLElement) {
		megaCrawler.SetTitle(element, element.Attr("content"))
	})

	s.OnHTML("article", func(element *colly.HTMLElement) {
		megaCrawler.SetContent(element, element.Text)
	})

	s.OnHTML(".pane--csis-publication-written-by", func(element *colly.HTMLElement) {
		megaCrawler.SetAuthor(element, element.ChildText(".teaser__title"))
	})
}

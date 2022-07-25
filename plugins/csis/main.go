package csis

import (
	"github.com/gocolly/colly/v2"
	"megaCrawler/megaCrawler"
	"time"
)

func init() {
	s := megaCrawler.Register("csis", "战略与国际研究中心", "https://www.csis.org/").
		SetStartingUrls([]string{"https://www.csis.org/analysis?&field_categories_field_regions%5B0%5D=797"}).
		SetTimeout(20 * time.Second)

	s.OnHTML(".ds-right", func(element *colly.HTMLElement) {
		t, err := time.Parse("2006-01-02T15:04:05-07:00", element.ChildAttr(".date-display-single", "content"))
		if megaCrawler.Debug && err != nil {
			_ = megaCrawler.Logger.Error("Could not parse time" + err.Error())
			t = time.Now()
		}
		s.AddUrl(element.ChildAttr(".teaser__title > a", "href"), t)
	})

	s.OnHTML(".pager__link", func(element *colly.HTMLElement) {
		s.AddUrl(element.Attr("href"), time.Now())
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

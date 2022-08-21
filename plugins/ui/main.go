package ui

import (
	"github.com/gocolly/colly/v2"
	"megaCrawler/megaCrawler"
	"time"
)

func init() {
	s := megaCrawler.Register("ui", "国际事务研究所", "https://www.ui.se/").
		SetStartingUrls([]string{"https://www.ui.se/butiken/uis-publikationer/ui-report/", "https://www.ui.se/butiken/uis-publikationer/ui-paper/", "https://www.ui.se/butiken/uis-publikationer/ui-brief/"}).
		SetTimeout(20 * time.Second)

	s.OnHTML(".block-link", func(element *colly.HTMLElement) {
		s.AddUrl(element.Attr("href"), time.Now())
	})

	s.OnHTML(".preamble", func(element *colly.HTMLElement) {
		megaCrawler.SetTitle(element, element.Text)
	})
	s.OnHTML("meta[property=\"og:title\"]", func(element *colly.HTMLElement) {
		megaCrawler.SetTitle(element, element.Attr("content"))
	})
}

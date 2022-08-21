package ceu

import (
	"github.com/gocolly/colly/v2"
	"megaCrawler/megaCrawler"
	"time"
)

func init() {
	s := megaCrawler.Register("ceu", "卡内基欧洲中心", "https://carnegieeurope.eu/").
		SetStartingUrls([]string{"https://carnegieeurope.eu/publications/"}).
		SetTimeout(20 * time.Second)

	s.OnHTML(".clearfix", func(element *colly.HTMLElement) {
		k, err := time.Parse("January 02, 2006", element.ChildText(".inline-block"))
		if err != nil {
			_ = megaCrawler.Logger.Errorf("Error when parsing %s to time: %s", element.ChildText(".inline-block"), err.Error())
			k = time.Now()
		}
		s.AddUrl(element.ChildAttr(".no-margin > a", "href"), k)
	})

	s.OnHTML(".page-link", func(element *colly.HTMLElement) {
		s.AddUrl(element.Attr("href"), time.Now())
	})

	s.OnHTML(".article-body", func(element *colly.HTMLElement) {
		megaCrawler.SetContent(element, element.Text)
	})

	s.OnHTML("meta[property=\"og:title\"]", func(element *colly.HTMLElement) {
		megaCrawler.SetTitle(element, element.Attr("content"))
	})

	s.OnHTML(".em", func(element *colly.HTMLElement) {
		megaCrawler.AppendAuthor(element, element.Text)
	})
}

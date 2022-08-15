package cna

import (
	"github.com/gocolly/colly/v2"
	"megaCrawler/megaCrawler"
	"regexp"
	"time"
)

func init() {
	s := megaCrawler.Register("cna", "海军分析中心", "https://www.cna.org/").
		SetStartingUrls([]string{"https://www.cna.org/our-research/explore-all", "https://www.cna.org/our-media/indepth"}).
		SetTimeout(20 * time.Second)

	s.OnHTML(".resource", func(element *colly.HTMLElement) {
		s.AddUrl(element.Attr("href"), time.Now())
	})

	s.OnHTML(".article", func(element *colly.HTMLElement) {
		s.AddUrl(element.Attr("href"), time.Now())
	})

	s.OnHTML(".paginator > a", func(element *colly.HTMLElement) {
		s.AddUrl(element.Attr("href"), time.Now())
	})

	s.OnHTML("meta[property=\"og:title\"]", func(element *colly.HTMLElement) {
		megaCrawler.SetTitle(element, element.Attr("content"))
	})

	s.OnHTML(".article-text", func(element *colly.HTMLElement) {
		megaCrawler.SetContent(element, element.Text)
	})

	s.OnHTML(".executive-summary", func(element *colly.HTMLElement) {
		megaCrawler.SetContent(element, element.Text)
	})

	s.OnHTML(".author", func(element *colly.HTMLElement) {
		megaCrawler.AppendAuthor(element, element.Text)
	})

	s.OnHTML(".dateline", func(element *colly.HTMLElement) {
		m, _ := regexp.Compile("\\w+ \\d+, \\d+")
		t, err := time.Parse("January 2, 2006", m.FindString(element.Text))
		if err != nil {
			t = time.Now()
		}
		megaCrawler.SetTime(element, t)
	})
}

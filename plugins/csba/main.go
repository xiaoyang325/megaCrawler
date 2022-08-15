package csba

import (
	"github.com/gocolly/colly/v2"
	"megaCrawler/megaCrawler"
	"strings"
	"time"
)

func init() {
	s := megaCrawler.Register("csba", "战略与预算评估中心", "https://csbaonline.org/").
		SetStartingUrls([]string{"https://csbaonline.org/research/publications"}).
		SetTimeout(20 * time.Second)

	s.OnHTML(".article", func(element *colly.HTMLElement) {
		t, err := time.Parse("Jan 2, 2006", element.ChildText("time"))
		if err != nil {
			t = time.Now()
		}
		s.AddUrl(element.ChildAttr("a", "href"), t)
	})

	s.OnHTML(".pagination > li > a", func(element *colly.HTMLElement) {
		s.AddUrl(element.Attr("href"), time.Now())
	})

	s.OnHTML("meta[property=\"og:title\"]", func(element *colly.HTMLElement) {
		megaCrawler.SetTitle(element, element.Attr("content"))
	})

	s.OnHTML(".article-content", func(element *colly.HTMLElement) {
		megaCrawler.SetContent(element, element.Text)
	})

	s.OnHTML(".article-meta > a", func(element *colly.HTMLElement) {
		if strings.Contains(element.Attr("href"), "people") {
			megaCrawler.AppendAuthor(element, element.Text)
		}
	})

	s.OnHTML(".article-single > .article-meta > .font-weight-bold", func(element *colly.HTMLElement) {
		megaCrawler.AppendAuthor(element, element.Text)
	})
}

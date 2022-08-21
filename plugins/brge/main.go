package brge

import (
	"github.com/gocolly/colly/v2"
	"megaCrawler/megaCrawler"
	"time"
)

func init() {
	s := megaCrawler.Register("brge", "布鲁盖尔研究所", "https://www.bruegel.org/").
		SetStartingUrls([]string{"https://www.bruegel.org/topics", "https://www.bruegel.org/publications", "https://www.bruegel.org/bruegel-blog"}).
		SetTimeout(20 * time.Second)

	s.OnHTML("article", func(element *colly.HTMLElement) {
		date, err := time.Parse("02 January 2006", element.ChildText(".c-list-item__date"))
		if err != nil {
			_ = megaCrawler.Logger.Errorf("Error when parsing %s to time: %s", element.ChildText(".c-list-item__date"), err.Error())
			date = time.Now()
		}
		s.AddUrl(element.Attr("about"), date)
	})

	s.OnHTML("meta[property=\"og:title\"]", func(element *colly.HTMLElement) {
		megaCrawler.SetTitle(element, element.Attr("content"))
	})

	s.OnHTML(".c-pager__button-link", func(element *colly.HTMLElement) {
		u, _ := element.Request.URL.Parse(element.Attr("href"))
		s.AddUrl(u.String(), time.Now())
	})

	s.OnHTML("a.c-single-header__meta-term", func(element *colly.HTMLElement) {
		megaCrawler.AppendAuthor(element, element.Text)
	})

	s.OnHTML(".o-content-from-editor", func(element *colly.HTMLElement) {
		megaCrawler.AppendContent(element, element.Text)
	})
}

package brooking

import (
	"github.com/gocolly/colly/v2"
	"megaCrawler/megaCrawler"
	"time"
)

func init() {
	s := megaCrawler.Register("brooking", "布鲁金斯学会", "https://www.brookings.edu/").
		FromRobotTxt("https://www.brookings.edu/robots.txt").
		SetTimeout(20 * time.Second)

	s.OnXML("//urlset/url", func(e *colly.XMLElement) {
		t, _ := time.Parse("2006-01-02T15:04:05-07:00", e.ChildText("lastmod"))
		s.AddUrl(e.ChildText("loc"), t)
	})

	s.OnXML("//sitemapindex/sitemap", func(e *colly.XMLElement) {
		s.AddUrl(e.ChildText("loc"), time.Now())
	})

	s.OnHTML("meta[property=\"og:title\"]", func(element *colly.HTMLElement) {
		megaCrawler.SetTitle(element, element.Attr("content"))
	})

	s.OnHTML("div[itemprop=\"articleBody\"]", func(element *colly.HTMLElement) {
		megaCrawler.SetContent(element, element.Text)
	})

	s.OnHTML(".article-header__link", func(element *colly.HTMLElement) {
		megaCrawler.SetAuthor(element, element.Text)
	})
}

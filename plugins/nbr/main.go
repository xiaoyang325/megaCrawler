package nbr

import (
	"github.com/gocolly/colly/v2"
	"megaCrawler/megaCrawler"
	"regexp"
	"time"
)

func init() {
	s := megaCrawler.Register("nbr", "国家亚洲研究局海洋意识项目", "https://map.nbr.org/").
		SetStartingUrls([]string{"https://map.nbr.org/sitemap.xml"}).
		SetTimeout(20 * time.Second)

	s.OnXML("//urlset/url", func(e *colly.XMLElement) {
		s.AddUrl(e.ChildText("loc"), time.Now())
	})

	s.OnXML("//sitemapindex/sitemap", func(e *colly.XMLElement) {
		s.AddUrl(e.ChildText("loc"), time.Now())
	})

	s.OnHTML("title", func(element *colly.HTMLElement) {
		megaCrawler.SetTitle(element, element.Text)
	})

	s.OnHTML(".elementor-widget-theme-post-content", func(element *colly.HTMLElement) {
		megaCrawler.SetContent(element, element.Text)
	})

	s.OnHTML(".elementor-post-info", func(element *colly.HTMLElement) {
		m1, _ := regexp.Compile("(.+) / (.*)")
		match := m1.FindStringSubmatch(megaCrawler.StandardizeSpaces(element.Text))
		if match == nil {
			return
		}

		megaCrawler.AppendAuthor(element, match[1])

		t, err := time.Parse("January 2, 2006", match[2])
		if err == nil {
			t = time.Now()
		}
		megaCrawler.SetTime(element, t)
	})
}

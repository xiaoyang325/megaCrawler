package eliamep

import (
	"github.com/gocolly/colly/v2"
	"megaCrawler/megaCrawler"
	"regexp"
	"time"
)

func init() {
	s := megaCrawler.Register("eliamep", "欧洲和外交政策基金会", "https://www.eliamep.gr").
		SetStartingUrls([]string{"https://www.eliamep.gr/en/publications/"}).
		SetTimeout(20 * time.Second)

	s.OnHTML(".medium-9", func(element *colly.HTMLElement) {
		m1, _ := regexp.Compile("\\w+ \\d+, \\d+")
		dateString := m1.FindString(element.ChildText(".date"))
		k, err := time.Parse("January 2, 2006", dateString)
		if err != nil {
			_ = megaCrawler.Logger.Errorf("Error when parsing %s to time: %s", element.ChildText(".views-field-field-year-of-publication"), err.Error())
			k = time.Now()
		}
		s.AddUrl(element.ChildAttr(".title", "href"), k)
	})

	s.OnHTML(".page-numbers", func(element *colly.HTMLElement) {
		s.AddUrl(element.Attr("href"), time.Now())
	})

	s.OnHTML(".articleBody", func(element *colly.HTMLElement) {
		megaCrawler.SetContent(element, element.Text)
	})

	s.OnHTML(".experts > a", func(element *colly.HTMLElement) {
		megaCrawler.AppendAuthor(element, element.Text)
	})

	s.OnHTML("h1", func(element *colly.HTMLElement) {
		megaCrawler.SetTitle(element, element.Text)
	})
}

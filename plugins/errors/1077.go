package errors

import (
	"github.com/gocolly/colly/v2"
	"megaCrawler/Crawler"
	"megaCrawler/Extractors"
	"strings"
)

func init() {
	engine := Crawler.Register("1077", "欧洲安全与合作组织", "https://www.osce.org")

	engine.SetStartingUrls([]string{"https://www.osce.org/sitemap.xml"})

	extractorConfig := Extractors.Config{
		Author:       true,
		Image:        true,
		Language:     true,
		PublishDate:  true,
		Tags:         true,
		Text:         true,
		Title:        true,
		TextLanguage: "",
	}

	extractorConfig.Apply(engine)

	engine.OnXML("//loc", func(element *colly.XMLElement, ctx *Crawler.Context) {
		if strings.Contains(element.Text, "sitemap.xml") {
			engine.Visit(element.Text, Crawler.Index)
			return
		}
		if !strings.Contains(element.Text, "?download") {
			engine.Visit(element.Text, Crawler.News)
		}
	})
}

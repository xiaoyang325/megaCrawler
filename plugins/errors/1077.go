package errors

import (
	"megaCrawler/crawlers"
	"megaCrawler/extractors"
	"strings"

	"github.com/gocolly/colly/v2"
)

func init() {
	engine := crawlers.Register("1077", "欧洲安全与合作组织", "https://www.osce.org")

	engine.SetStartingURLs([]string{"https://www.osce.org/sitemap.xml"})

	extractorConfig := extractors.Config{
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

	engine.OnXML("//loc", func(element *colly.XMLElement, ctx *crawlers.Context) {
		if strings.Contains(element.Text, "sitemap.xml") {
			engine.Visit(element.Text, crawlers.Index)
			return
		}
		if !strings.Contains(element.Text, "?download") {
			engine.Visit(element.Text, crawlers.News)
		}
	})
}

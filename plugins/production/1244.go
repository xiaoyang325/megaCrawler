package production

import (
	"strings"

	"github.com/gocolly/colly/v2"

	"megaCrawler/crawlers"
	"megaCrawler/extractors"
)

func init() {
	engine := crawlers.Register("1244", "莫斯科物理技术学院", "https://mipt.ru")

	engine.SetStartingURLs([]string{"https://mipt.ru/sitemap.xml"})

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
		switch {
		case strings.Contains(element.Text, ".xml"):
			engine.Visit(element.Text, crawlers.Index)
		case strings.Contains(element.Text, "/persons/"):
			engine.Visit(element.Text, crawlers.Expert)
		case strings.Contains(element.Text, "/news/"):
			engine.Visit(element.Text, crawlers.News)
		}
	})
}

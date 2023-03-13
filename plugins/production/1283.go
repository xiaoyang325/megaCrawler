package production

import (
	"strings"

	"github.com/gocolly/colly/v2"

	"megaCrawler/crawlers"
	"megaCrawler/extractors"
)

func init() {
	engine := crawlers.Register("1283", "哈德森研究所", "https://www.hudson.org")

	engine.SetStartingURLs([]string{})

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
		case strings.Contains(element.Text, "/experts/"):
			engine.Visit(element.Text, crawlers.Expert)
		default:
			engine.Visit(element.Text, crawlers.News)
		}
	})
}

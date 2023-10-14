package production

import (
	"strings"

	"github.com/gocolly/colly/v2"

	"megaCrawler/crawlers"
	"megaCrawler/extractors"
)

func init() {
	engine := crawlers.Register("1272", "国际应用系统分析研究所", "https://iiasa.ac.at")

	engine.SetStartingURLs([]string{"https://iiasa.ac.at/sitemap.xml"})

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
		case strings.Contains(element.Text, "/staff/"):
			engine.Visit(element.Text, crawlers.Expert)
		case strings.Contains(element.Text, "/news/"):
			engine.Visit(element.Text, crawlers.News)
		}
	})
}

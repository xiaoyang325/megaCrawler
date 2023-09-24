package production

import (
	"strings"

	"megaCrawler/crawlers"
	"megaCrawler/extractors"

	"github.com/gocolly/colly/v2"
)

func init() {
	engine := crawlers.Register("reuters", "Reuters", "https://www.reuters.com/")

	engine.SetStartingURLs([]string{"https://www.reuters.com/arc/outboundfeeds/news-sitemap-index/?outputType=xml"})

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
		if strings.Contains(element.Text, "news-sitemap") {
			engine.Visit(element.Text, crawlers.Index)
			return
		}
		engine.Visit(element.Text, crawlers.News)
	})
}

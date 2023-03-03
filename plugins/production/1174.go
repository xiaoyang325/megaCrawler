package production

import (
	"megaCrawler/crawlers"
	"megaCrawler/extractors"
	"strings"

	"github.com/gocolly/colly/v2"
)

func init() {
	engine := crawlers.Register("1174", "彼得森国际经济研究所", "https://www.piie.com")

	engine.SetStartingURLs([]string{"https://www.piie.com/sitemap.xml"})

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
		engine.Visit(element.Text, crawlers.News)
	})
}

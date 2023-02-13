package production

import (
	"megaCrawler/crawlers"
	"megaCrawler/extractors"
	"strings"

	"github.com/gocolly/colly/v2"
)

func init() {
	engine := crawlers.Register("1075", "斯德哥尔摩国际和平研究所", "https://www.sipri.org")

	engine.SetStartingURLs([]string{"https://www.sipri.org/sitemap.xml"})

	extractorConfig := extractors.Config{
		Author:      true,
		Image:       true,
		Language:    true,
		PublishDate: true,
		Tags:        true,
		Text:        true,
		Title:       true,
	}

	extractorConfig.Apply(engine)

	engine.OnXML("//loc", func(element *colly.XMLElement, ctx *crawlers.Context) {
		if strings.Contains(element.Text, "sitemap.xml") {
			engine.Visit(element.Text, crawlers.Index)
		}
		if strings.Contains(element.Text, "/about/bios/") {
			engine.Visit(element.Text, crawlers.Expert)
		}
		if strings.Contains(element.Text, "/node/") || strings.Contains(element.Text, "/news/") {
			engine.Visit(element.Text, crawlers.News)
		}
	})
}

package production

import (
	"strings"

	"github.com/gocolly/colly/v2"

	"megaCrawler/crawlers"
	"megaCrawler/extractors"
)

func init() {
	engine := crawlers.Register("1210", "伍德罗威尔逊中心", "https://www.wilsoncenter.org")

	engine.SetStartingURLs([]string{"https://www.wilsoncenter.org/sitemap.xml"})

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
		case strings.Contains(element.Text, "sitemap.xml"):
			engine.Visit(element.Text, crawlers.Index)
		case strings.Contains(element.Text, "/person/"):
			engine.Visit(element.Text, crawlers.Expert)
		case strings.Contains(element.Text, "/article/"):
			engine.Visit(element.Text, crawlers.News)
		}
	})
}

package production

import (
	"megaCrawler/crawlers"
	"megaCrawler/extractors"

	"github.com/gocolly/colly/v2"
)

func init() {
	engine := crawlers.Register("1046", "激进派右翼分析中心", "https://www.radicalrightanalysis.com")

	engine.SetStartingUrls([]string{"https://www.radicalrightanalysis.com/wp-sitemap-posts-post-1.xml"})

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
		engine.Visit(element.Text, crawlers.News)
	})
}

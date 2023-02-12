package production

import (
	"github.com/gocolly/colly/v2"
	"megaCrawler/Crawler"
	"megaCrawler/Extractors"
)

func init() {
	engine := Crawler.Register("1046", "激进派右翼分析中心", "https://www.radicalrightanalysis.com")

	engine.SetStartingUrls([]string{"https://www.radicalrightanalysis.com/wp-sitemap-posts-post-1.xml"})

	extractorConfig := Extractors.Config{
		Author:      true,
		Image:       true,
		Language:    true,
		PublishDate: true,
		Tags:        true,
		Text:        true,
		Title:       true,
	}

	extractorConfig.Apply(engine)

	engine.OnXML("//loc", func(element *colly.XMLElement, ctx *Crawler.Context) {
		engine.Visit(element.Text, Crawler.News)
	})
}

package production

import (
	"strings"

	"github.com/gocolly/colly/v2"

	"megaCrawler/crawlers"
	"megaCrawler/extractors"
)

func init() {
	engine := crawlers.Register("1260", "伊斯兰堡战略研究所", "https://issi.org.pk")

	engine.SetStartingURLs([]string{"https://issi.org.pk/wp-sitemap.xml"})

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
		case strings.Contains(ctx.URL, "wp-sitemap.xml"):
			engine.Visit(element.Text, crawlers.Index)
		case strings.Contains(ctx.URL, "wp-sitemap-posts"):
			engine.Visit(element.Text, crawlers.News)
		}
	})

	engine.OnHTML(".td-post-content", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Content = element.Text
	})
}

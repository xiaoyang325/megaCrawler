package production

import (
	"megaCrawler/crawlers"
	"megaCrawler/extractors"
	"strings"

	"github.com/gocolly/colly/v2"
)

func init() {
	engine := crawlers.Register("1141", "中美研究中心", "https://chinaus-icas.org")

	engine.SetStartingURLs([]string{"https://chinaus-icas.org/research-sitemap.xml", "https://chinaus-icas.org/icas_in_the_news-sitemap.xml"})

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
		case strings.Contains(ctx.URL, "icas_in_the_news-sitemap.xml"):
			engine.Visit(element.Text, crawlers.News)
		case strings.Contains(ctx.URL, "research-sitemap.xml"):
			engine.Visit(element.Text, crawlers.Report)
		}
	})
}

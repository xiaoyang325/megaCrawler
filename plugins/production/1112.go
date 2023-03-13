package production

import (
	"strings"

	"megaCrawler/crawlers"
	"megaCrawler/extractors"

	"github.com/gocolly/colly/v2"
)

func init() {
	engine := crawlers.Register("1112", "国家利益中心", "https://cftni.org")

	engine.SetStartingURLs([]string{"https://cftni.org/wp-sitemap-posts-post-1.xml", "https://cftni.org/wp-sitemap-posts-expert-1.xml"})

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
		case strings.Contains(ctx.URL, "expert"):
			engine.Visit(element.Text, crawlers.Expert)
		case strings.Contains(ctx.URL, "post"):
			engine.Visit(element.Text, crawlers.News)
		}
	})
}

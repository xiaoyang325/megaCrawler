package production

import (
	"strings"

	"megaCrawler/crawlers"
	"megaCrawler/extractors"

	"github.com/gocolly/colly/v2"
)

func init() {
	engine := crawlers.Register("1118", "华府公民道德责任组织", "https://www.citizensforethics.org")

	engine.SetStartingURLs([]string{"https://www.citizensforethics.org/post-sitemap.xml", "https://www.citizensforethics.org/reports-sitemap.xml"})

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
		if strings.Contains(ctx.URL, "post-sitemap.xml") {
			engine.Visit(element.Text, crawlers.News)
		} else if strings.Contains(ctx.URL, "reports-sitemap.xml") {
			engine.Visit(element.Text, crawlers.Report)
		}
	})

	engine.OnHTML(".actions-latest-update__date", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.PublicationTime = element.Text
	})
}

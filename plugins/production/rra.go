package production

import (
	"megaCrawler/crawlers"
	"megaCrawler/extractors"
	"strings"

	"github.com/gocolly/colly/v2"
)

func init() {
	w := crawlers.Register("rra", "激进派右翼分析中心", "https://www.radicalrightanalysis.com/")

	w.SetStartingUrls([]string{"https://www.radicalrightanalysis.com/wp-sitemap-posts-post-1.xml"})

	w.OnXML("//loc", func(element *colly.XMLElement, ctx *crawlers.Context) {
		w.Visit(element.Text, crawlers.News)
	})

	w.OnHTML("html", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		extractors.Titles(ctx, element)
		extractors.Authors(ctx, element)
		extractors.PublishingDate(ctx, element)
		extractors.Tags(ctx, element)
	})

	w.OnHTML(".entry-content", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Content = strings.TrimSpace(element.Text)
	})
}

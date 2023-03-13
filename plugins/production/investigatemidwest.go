package production

import (
	"strings"

	"megaCrawler/crawlers"
	"megaCrawler/extractors"

	"github.com/gocolly/colly/v2"
)

func init() {
	w := crawlers.Register("investigatemidwest", "中西部调查报道中心", "https://www.investigatemidwest.org/")

	w.SetStartingURLs([]string{"https://investigatemidwest.org/post.xml"})

	w.OnXML("//loc", func(element *colly.XMLElement, ctx *crawlers.Context) {
		w.Visit(element.Text, crawlers.News)
	})

	w.OnHTML("html", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		extractors.Titles(ctx, element)
		extractors.PublishingDate(ctx, element)
	})

	w.OnHTML(".subtitle", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.SubTitle = element.Text
	})

	w.OnHTML(".entry-content", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Content = extractors.TrimText(element.DOM)
	})

	w.OnHTML(".post-category-link", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.CategoryText = strings.TrimSpace(element.Text)
	})
}

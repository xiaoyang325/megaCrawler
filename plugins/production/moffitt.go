package production

import (
	"time"

	"megaCrawler/crawlers"
	"megaCrawler/extractors"

	"github.com/araddon/dateparse"
	"github.com/gocolly/colly/v2"
)

func init() {
	w := crawlers.Register("moffitt", "莫菲特癌症中心", "https://moffitt.org/")

	w.SetStartingURLs([]string{"https://moffitt.org/XMLsitemap"})

	w.OnXML("//loc", func(element *colly.XMLElement, ctx *crawlers.Context) {
		w.Visit(element.Text, crawlers.Index)
	})

	w.OnHTML(".article-head", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Title = element.Text
	})

	w.OnHTML(".m-article__content", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.PageType = crawlers.News
		ctx.Content = extractors.TrimText(element.DOM)
	})

	w.OnHTML("h3", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.CategoryText = element.Text
	})

	w.OnHTML(".fa-tag", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Tags = append(ctx.Tags, element.Text)
	})

	w.OnHTML(".article > .text-sm", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		t, _ := dateparse.ParseAny(element.Text)
		ctx.PublicationTime = t.Format(time.RFC3339)
	})
}

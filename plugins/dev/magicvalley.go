package dev

import (
	"megaCrawler/crawlers"

	"github.com/gocolly/colly/v2"
)

func init() {
	w := crawlers.Register("", "", "https://magicvalley.com/")

	w.SetStartingURLs([]string{
		"https://magicvalley.com/", "https://magicvalley.com/news/#tracking-source=main-nav", "https://magicvalley.com/sports/#tracking-source=main-nav", "https://jobs.magicvalley.com/#tracking-source=main-nav",
	})

	w.OnHTML(".list-inline>li>.tnt-byline", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Authors = append(ctx.Authors, element.Text)
	})
	w.OnHTML(".asset-content", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Content += element.Text
	})
	w.OnHTML("a.tnt-asset-link", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		w.Visit(element.Attr("href"), crawlers.News)
	})
	w.OnHTML(".headline", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Title += element.Text
	})

}

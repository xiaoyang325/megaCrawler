package dev

import (
	"megaCrawler/crawlers"

	"github.com/gocolly/colly/v2"
)

func init() {
	w := crawlers.Register("philstar", "philstar", "https://www.philstar.com/")

	w.SetStartingURLs([]string{
		"https://www.philstar.com/", "https://www.philstar.com/headlines", "https://www.philstar.com/opinion", "https://www.philstar.com/nation", "https://www.philstar.com/world", "https://www.philstar.com/business", "https://www.philstar.com/sports", "https://www.philstar.com/entertainment", "https://www.philstar.com/lifestyle",
	})

	w.OnHTML("div.article__credits-author-pub", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Authors = append(ctx.Authors, element.Text)
	})
	w.OnHTML("div.article__writeup", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Content += element.Text
	})
	w.OnHTML("div.article__date-published", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.PublicationTime += element.Text
	})
	w.OnHTML("div.article__title", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Title += element.Text
	})
}

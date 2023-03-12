package dev

import (
	"megaCrawler/crawlers"

	"github.com/gocolly/colly/v2"
)

func init() {
	w := crawlers.Register("bnn-news", "bnn-news", "https://bnn-news.com/")

	w.SetStartingURLs([]string{
		"https://bnn-news.com/", "https://bnn-news.com/category/baltics", "https://bnn-news.com/category/social", "https://bnn-news.com/category/politics", "https://bnn-news.com/category/money", "https://bnn-news.com/category/interviews", "https://bnn-news.com/category/opinion", "https://bnn-news.com/category/bnn-investigation", "https://bnn-news.com/category/world", "https://bnn-news.com/category/leisure", "https://bnn-news.com/category/advertising",
	})

	w.OnHTML(".tdb-author-name", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Authors = append(ctx.Authors, element.Text)
	})
	w.OnHTML(".tdb_single_content", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Content += element.Text
	})
	w.OnHTML("a.td-image-wrap", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		w.Visit(element.Attr("href"), crawlers.News)
	})
	w.OnHTML(".entry-title>a", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		w.Visit(element.Attr("href"), crawlers.News)
	})
	w.OnHTML(".entry-date", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.PublicationTime += element.Text
	})
	w.OnHTML(".tdb-title-text", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Title += element.Text
	})

}

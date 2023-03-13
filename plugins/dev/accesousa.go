package dev

import (
	"megaCrawler/crawlers"

	"github.com/gocolly/colly/v2"
)

func init() {
	w := crawlers.Register("accesousa", "accesousa", "https://www.accesousa.com/")

	w.SetStartingURLs([]string{
		"https://www.accesousa.com/noticias/#navlink=subnav", "https://www.accesousa.com/finanzas/#navlink=navbar", "https://www.accesousa.com/inmigracion/#navlink=navbar", "https://www.accesousa.com/cultura/#navlink=navbar", "https://www.accesousa.com/topics/impuestos-estados-unidos/#navlink=navbar",
	})

	w.OnHTML(".header>.bio>div>div>.byline", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Authors = append(ctx.Authors, element.Text)
	})
	w.OnHTML(".story-body", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Content += element.Text
	})
	w.OnHTML("h3>a", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		w.Visit(element.Attr("href"), crawlers.News)
	})
	w.OnHTML("a.kicker", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		w.Visit(element.Attr("href"), crawlers.News)
	})
	w.OnHTML("#update_date", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.PublicationTime += element.Text
	})
	w.OnHTML(".story-body>.header>.h1", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Title += element.Text
	})
}

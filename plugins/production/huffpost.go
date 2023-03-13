package production

import (
	"megaCrawler/crawlers"

	"github.com/gocolly/colly/v2"
)

func init() {
	w := crawlers.Register("huffpost", "huffpost", "https://www.huffpost.com/")

	w.SetStartingURLs([]string{
		"https://www.huffpost.com/",
	})

	w.OnHTML("div.entry__byline__author>div>a.js-entry-link", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Authors = append(ctx.Authors, element.Text)
	})
	w.OnHTML("div.entry__content-list-container", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Content += element.Text
	})
	w.OnHTML("div.author--bio__bio", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Description += element.Text
	})
	w.OnHTML("div.entry__byline__author>div>a.js-entry-link", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		w.Visit(element.Attr("href"), crawlers.Expert)
	})
	w.OnHTML("a.left-nav__menu__item", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		w.Visit(element.Attr("href"), crawlers.Index)
	})
	w.OnHTML("a.subnav__item", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		w.Visit(element.Attr("href"), crawlers.Index)
	})
	w.OnHTML("ul.pagination>li>a", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		w.Visit(element.Attr("href"), crawlers.Index)
	})
	w.OnHTML("div.author--bio__details>div", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Email = element.Text
	})
	w.OnHTML("h1.author--bio__name", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Name += element.Text
	})
	w.OnHTML("a.card__headline", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		w.Visit(element.Attr("href"), crawlers.News)
	})
	w.OnHTML("div.timestamp", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.PublicationTime += element.Text
	})
	w.OnHTML("h1.headline", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Title += element.Text
	})
	w.OnHTML("p.author--bio__title", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Title += element.Text
	})
}

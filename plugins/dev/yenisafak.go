package dev

import (
	"megaCrawler/crawlers"

	"github.com/gocolly/colly/v2"
)

func init() {
	w := crawlers.Register("yenisafak", "yenisafak", "https://www.yenisafak.com/")

	w.SetStartingURLs([]string{
		"https://www.yenisafak.com/",
	})

	w.OnHTML(".author-about-card-info__name", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Authors = append(ctx.Authors, element.Text)
	})
	w.OnHTML(".item>.content", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Content += element.Text
	})
	w.OnHTML(".ys-header-menu__cards__card", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		w.Visit(element.Attr("href"), crawlers.Index)
	})
	w.OnHTML(".author-about-card-social-media>a", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Link = append(ctx.Authors, element.Text)
	})
	w.OnHTML("#__layout > div > div.layout-content > div.detail-page.author-column-detail-page > div > div.author-about-card > div.author-about-card-container > div > div > div.author-about-card-content__author-info-content > div.author-about-card-social-media > div > a", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Link = append(ctx.Authors, element.Text)
	})
	w.OnHTML(".ys-link>a", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		w.Visit(element.Attr("href"), crawlers.News)
	})
	w.OnHTML(".author-about-card-info__date", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.PublicationTime += element.Text
	})
	w.OnHTML(".author-about-card-meta__title", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Title += element.Text
	})

}

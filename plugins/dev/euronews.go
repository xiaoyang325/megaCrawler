package dev

import (
	"megaCrawler/crawlers"

	"github.com/gocolly/colly/v2"
)

func init() {
	w := crawlers.Register("euronews", "euronews", "https://www.euronews.com/")

	w.SetStartingURLs([]string{
		"https://www.euronews.com/",
	})

	w.OnHTML("div.swiper-slide-active>div>div.o-article__body>div>article>header>div.c-article-contributors", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Authors = append(ctx.Authors, element.Text)
	})
	w.OnHTML("div.swiper-slide-active>div>div.o-article__body>div>article>section.c-article__container", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Content += element.Text
	})
	w.OnHTML("a.list-item__link", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		w.Visit(element.Attr("href"), crawlers.Index)
	})
	w.OnHTML("a.m-object__title__link", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		w.Visit(element.Attr("href"), crawlers.News)
	})
	w.OnHTML("div.swiper-slide-active>div>div.o-article__body>div>article>header>div.c-article-contributors>time.c-article-date", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.PublicationTime += element.Text
	})
	w.OnHTML("div.swiper-slide-active>div.o-article>div.o-article__body>div>article>header>div>h1.c-article-title", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Title += element.Text
	})
}

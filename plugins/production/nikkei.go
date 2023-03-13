package production

import (
	"megaCrawler/crawlers"

	"github.com/gocolly/colly/v2"
)

func init() {
	w := crawlers.Register("nikkei", "nikkei", "https://asia.nikkei.com/")

	w.SetStartingURLs([]string{
		"https://asia.nikkei.com/",
	})

	w.OnHTML("span.article__author", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Authors = append(ctx.Authors, element.Text)
	})
	w.OnHTML("div#article-body-preview", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Content += element.Text
	})
	w.OnHTML("a.load-more", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		w.Visit(element.Attr("href"), crawlers.Index)
	})
	w.OnHTML("a.nav__link", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		w.Visit(element.Attr("href"), crawlers.Index)
	})
	w.OnHTML("a.article-block__title", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		w.Visit(element.Attr("href"), crawlers.Index)
	})
	w.OnHTML("a.article-block__primary-tag", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		w.Visit(element.Attr("href"), crawlers.News)
	})
	w.OnHTML("ul.card-article__footer>li>a", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		w.Visit(element.Attr("href"), crawlers.News)
	})
	w.OnHTML("h4.card-article__headline", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		w.Visit(element.Attr("href"), crawlers.News)
	})
	w.OnHTML(".card-article__headline>a", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		w.Visit(element.Attr("href"), crawlers.News)
	})
	w.OnHTML("a.article-block__title", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		w.Visit(element.Attr("href"), crawlers.News)
	})
	w.OnHTML("time.timestamp__time", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.PublicationTime += element.Text
	})
	w.OnHTML("h1.article-header__title", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Title += element.Text
	})
}

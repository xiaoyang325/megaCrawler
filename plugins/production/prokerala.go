package production

import (
	"megaCrawler/crawlers"
	"strings"

	"github.com/gocolly/colly/v2"
)

func init() {
	w := crawlers.Register("prokerala", "prokerala", "https://www.prokerala.com/")

	w.SetStartingURLs([]string{
		"https://www.prokerala.com/", "https://www.prokerala.com/news/",
	})

	w.OnHTML("div.article-info > span:nth-child(1) > a", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Authors = append(ctx.Authors, element.Text)
	})
	w.OnHTML("div.article-content", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Content += element.Text
	})
	w.OnHTML("ul.wdgt-menu-list>li>a", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		w.Visit(element.Attr("href"), crawlers.Index)
	})
	w.OnHTML("div.media-card-caption>a", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		w.Visit(element.Attr("href"), crawlers.News)
	})
	w.OnHTML("div.card-caption>p>a", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		w.Visit(element.Attr("href"), crawlers.News)
	})
	w.OnHTML("div.caption-wrapper>a", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		w.Visit(element.Attr("href"), crawlers.News)
	})
	w.OnHTML("div.article-info > span:nth-child(2)", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.PublicationTime = strings.Split(element.Text, ",")[1]
	})
	w.OnHTML("header.page-header > h1", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Title += element.Text
	})
}

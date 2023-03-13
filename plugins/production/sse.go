package production

import (
	"megaCrawler/crawlers"

	"github.com/gocolly/colly/v2"
)

func init() {
	w := crawlers.Register("sse", "sse", "http://english.sse.com.cn/")

	w.SetStartingURLs([]string{
		"http://english.sse.com.cn/", "http://english.sse.com.cn/news/newsrelease/", "http://english.sse.com.cn/events/conference/",
	})

	w.OnHTML("body > div.main.mainbg > div.container.article > p:nth-child(3)", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Authors = append(ctx.Authors, element.Text)
	})
	w.OnHTML("div.article", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Content += element.Text
	})
	w.OnHTML("a.more", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		w.Visit(element.Attr("href"), crawlers.Index)
	})
	w.OnHTML("ul.list_news>li>a", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		w.Visit(element.Attr("href"), crawlers.News)
	})
	w.OnHTML("a.title_link", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		w.Visit(element.Attr("href"), crawlers.News)
	})
	w.OnHTML("div.article_opt", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.PublicationTime += element.Text
	})
	w.OnHTML("div.article>h1", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Title += element.Text
	})
}

package commoncause

import (
	"github.com/gocolly/colly/v2"
	_ "github.com/gocolly/colly/v2"
	"megaCrawler/Crawler"
)

func init() {
	w := Crawler.Register("commoncause", "Common Cause", "https://www.commoncause.org/")

	w.SetStartingUrls([]string{"https://www.commoncause.org/democracy-wire/",
		"https://www.commoncause.org/our-work/voting-and-elections/",
		"https://www.commoncause.org/our-work/gerrymandering-and-representation/",
		"https://www.commoncause.org/our-work/ethics-and-accountability/",
		"https://www.commoncause.org/our-work/money-influence/",
		"https://www.commoncause.org/our-work/media-and-democracy/",
		"https://www.commoncause.org/our-work/constitution-courts-and-democracy-issues/",
		"https://www.commoncause.org/what-we-do/",
		"https://www.commoncause.org/news/",
		"https://www.commoncause.org/resources/",
		"https://www.commoncause.org/news-clips"})

	// 从翻页器获取链接并访问
	w.OnHTML("a.page-numbers", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		w.Visit(element.Attr("href"), Crawler.Index)
	})

	// 从index访问新闻
	w.OnHTML("a.full-link", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		w.Visit(element.Attr("href"), Crawler.News)
	})
	// report.title
	w.OnHTML("h1.page-title", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		ctx.Title = element.Text
	})
	//report.publish time
	w.OnHTML("div.post-info", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		ctx.PublicationTime = element.Text
	})
	// report.author
	w.OnHTML("span.name", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		ctx.Authors = append(ctx.Authors, element.Text)
	})
	// report.link
	w.OnHTML("span.number", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		ctx.Link = append(ctx.Link, element.Attr("href"))
	})
	w.OnHTML("li.contact>a", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		ctx.Link = append(ctx.Link, element.Attr("href"))
	})
	// report .content
	w.OnHTML("main>div.content>.module", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		ctx.Content = element.Text
	})

}

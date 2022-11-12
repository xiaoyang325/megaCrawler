package spri

import (
	"github.com/gocolly/colly/v2"
	"megaCrawler/megaCrawler"
)

func init() {
	w := megaCrawler.Register("spri", "安全政策改革研究所", "https://www.securityreform.org/")
	w.SetStartingUrls([]string{"https://www.securityreform.org/news-and-analysis"})

	//访问新闻
	w.OnHTML("a.summary-title-link", func(element *colly.HTMLElement, ctx *megaCrawler.Context) {
		w.Visit(element.Attr("href"), megaCrawler.News)
	})

	//获取新闻标题
	w.OnHTML("header > h1 > a", func(element *colly.HTMLElement, ctx *megaCrawler.Context) {
		ctx.Title = element.Text
	})

	//获取新闻时间
	w.OnHTML(" header > div > span.date > a > time", func(element *colly.HTMLElement, ctx *megaCrawler.Context) {
		ctx.PublicationTime = element.Text
	})

	//获取正文
	w.OnHTML("#page>article>div>div>div>div>div>div>div>div>p", func(element *colly.HTMLElement, ctx *megaCrawler.Context) {
		ctx.Content = element.Text
	})
}

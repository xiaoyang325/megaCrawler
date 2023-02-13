package production

import (
	"megaCrawler/crawlers"

	"github.com/gocolly/colly/v2"
)

func init() {
	w := crawlers.Register("people", "人民日报", "http://en.people.cn/")
	w.SetStartingUrls([]string{"http://en.people.cn/90780/index.html",
		"http://en.people.cn/90785/index.html",
		"http://en.people.cn/90777/index.html",
		"http://en.people.cn/business/index.html",
		"http://en.people.cn/90882/index.html",
		"http://en.people.cn/90782/index.html",
		"http://en.people.cn/202936/index.html"})

	// index
	w.OnHTML("div.col-1.fl > div > a", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		w.Visit(element.Attr("href"), crawlers.Index)
	})

	// 访问新闻
	w.OnHTML("div.col-1.fl > ul > li> a", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		w.Visit(element.Attr("href"), crawlers.News)
	})

	// 标题
	w.OnHTML("div.w860.d2txtCon.cf > h1", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Title = element.Text
	})

	// 正文
	w.OnHTML(".w860.d2txtCon.cf > p", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Content += element.Text
	})
}

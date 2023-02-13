package production

import (
	"megaCrawler/crawlers"

	"github.com/gocolly/colly/v2"
)

func init() {
	w := crawlers.Register("si", "史密斯研究所", "https://www.smithinst.co.uk/")
	w.SetStartingURLs([]string{"https://www.smithinst.co.uk/insights/",
		"https://www.smithinst.co.uk/the-team/"})

	// index
	w.OnHTML("a.btn.btn-lg.btn-link", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		w.Visit(element.Attr("href"), crawlers.Index)
	})

	// 访问新闻
	w.OnHTML(".post-link", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		w.Visit(element.Attr("href"), crawlers.News)
	})

	// 获取新闻标题
	w.OnHTML("#header > div > div > div > h1", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Title = element.Text
	})

	// 获取新闻作者
	w.OnHTML("#header > div > div > div > p:nth-child(3)", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Authors = append(ctx.Authors, element.Text)
	})

	// 获取新闻时间
	w.OnHTML("#header > div > div > div > p.small.text-muted", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.PublicationTime = element.Text
	})

	// 获取正文
	w.OnHTML("div.row > div.col-md-7 > p", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Content = element.Text
	})

	// 获取人物信息
	w.OnHTML(".person", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		subCtx := ctx.CreateSubContext()
		subCtx.PageType = crawlers.Expert
		subCtx.Name = element.ChildText(".person-name")
		subCtx.Title = element.ChildText(".person-position")
		subCtx.Description = element.ChildText(".person > div > p")
	})
}

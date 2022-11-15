package si

import (
	"github.com/gocolly/colly/v2"
	"megaCrawler/Crawler"
)

func init() {
	w := Crawler.Register("si", "史密斯研究所", "https://www.smithinst.co.uk/")
	w.SetStartingUrls([]string{"https://www.smithinst.co.uk/insights/",
		"https://www.smithinst.co.uk/the-team/"})

	//index
	w.OnHTML("a.btn.btn-lg.btn-link", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		w.Visit(element.Attr("href"), Crawler.Index)
	})

	//访问新闻
	w.OnHTML("h2.post-title.h5", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		w.Visit(element.Attr("href"), Crawler.News)
	})

	//获取新闻标题
	w.OnHTML("#header > div > div > div > h1", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		ctx.Title = element.Text
	})

	//获取新闻作者
	w.OnHTML("#header > div > div > div > p:nth-child(3)", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		ctx.Authors = append(ctx.Authors, element.Text)
	})

	//获取新闻时间
	w.OnHTML("#header > div > div > div > p.small.text-muted", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		ctx.PublicationTime = element.Text
	})

	//获取正文
	w.OnHTML("  div.row > div.col-md-7 > p", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		ctx.Content = element.Text
	})

	//获取人物信息
	w.OnHTML(".person", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		subCtx := ctx.CreateSubContext()
		subCtx.PageType = Crawler.Expert
		subCtx.Name = element.ChildText(".person-name")
		subCtx.Title = element.ChildText(".person-position")
		subCtx.Description = element.ChildText(".person > div > p")
	})
}

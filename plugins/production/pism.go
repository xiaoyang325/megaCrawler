package production

import (
	"strings"

	"megaCrawler/crawlers"

	"github.com/gocolly/colly/v2"
)

func init() {
	w := crawlers.Register("pism", "国际事务研究所", "https://www.pism.pl/")

	w.SetStartingURLs([]string{"https://www.pism.pl/publikacje",
		"https://www.pism.pl/konferencje",
		"https://www.pism.pl/pism_w_mediach/media",
		"https://www.pism.pl/projekty"})

	// 尝试寻找下载pdf的按钮，并如果存在则将页面类型转换为报告
	w.OnHTML("a.button", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		if strings.Contains(element.Attr("href"), ".pdf") {
			ctx.File = append(ctx.File, element.Attr("href"))
			ctx.PageType = crawlers.Report
		}
	})

	// 从翻页器获取链接并访问
	w.OnHTML("ul.pagination>li>a", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		w.Visit(element.Attr("href"), crawlers.Index)
	})
	// 从index访问新闻
	w.OnHTML("div.article-title>a", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		w.Visit(element.Attr("href"), crawlers.News)
	})

	// report.title
	w.OnHTML("div.title", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Title = element.Text
	})
	// report.publish time
	w.OnHTML("div.date", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.PublicationTime = element.Text
	})
	// report.description
	w.OnHTML("div.lead>p", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Description = element.Text
	})
	// report.author
	w.OnHTML("div.author>a", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Authors = append(ctx.Authors, element.Text)
	})
	// report .content
	w.OnHTML("div.richtext-preview", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Content = element.Text
	})

	// 内含Expert
	w.OnHTML("div.author>a", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		w.Visit(element.Attr("href"), crawlers.Expert)
	})

	// expert.Name
	w.OnHTML("div.name", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Name = element.Text
	})
	// expert.title
	w.OnHTML("div.position", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Title = element.Text
	})

	// expert.link
	w.OnHTML(".email-twitter>a", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Link = append(ctx.Link, element.Attr("href"))
	})

	w.OnHTML("div.richtext-preview", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Content = element.Text
	})
}

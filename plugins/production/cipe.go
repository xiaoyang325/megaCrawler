package production

import (
	"strings"

	"megaCrawler/crawlers"

	"github.com/gocolly/colly/v2"
)

func init() {
	w := crawlers.Register("cipe", "国际私营企业中心", "https://www.cipe.org/")
	w.SetStartingURLs([]string{
		"https://www.cipe.org/who-we-are/leadership/",
		"https://www.cipe.org/who-we-are/board/",
		"https://www.cipe.org/resources/",
		"https://www.cipe.org/blog/",
	})

	// 访问专家
	w.OnHTML(".people--grid-teaser > a", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		w.Visit(element.Attr("href"), crawlers.Expert)
	})

	// 专家姓名
	w.OnHTML(".page__title", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Name = element.Text
	})

	// 专家头衔
	w.OnHTML(".page__subtitle", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Title = element.Text
	})

	// 专家描述,新闻正文
	w.OnHTML(".field--body", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		if ctx.PageType == crawlers.Expert {
			ctx.Description = element.Text
		} else if ctx.PageType == crawlers.News || ctx.PageType == crawlers.Report {
			ctx.Content = element.Text
		}
	})

	// index
	w.OnHTML(".custom-pagination > a", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		w.Visit(element.Attr("href"), crawlers.Index)
	})

	// 访问文章
	w.OnHTML("div.listing-teaser__group-2 > h2 > a", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		w.Visit(element.Attr("href"), crawlers.News)
	})

	// 尝试寻找下载pdf的按钮，并如果存在则将页面类型转换为报告
	w.OnHTML(".download-button", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		if strings.Contains(element.Attr("href"), ".pdf") {
			ctx.File = append(ctx.File, element.Attr("href"))
			ctx.PageType = crawlers.Report
		}
	})

	// 标题
	w.OnHTML("h1.entry-title", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Title = element.Text
	})

	// 作者
	w.OnHTML(".page__header__meta>a", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Authors = append(ctx.Authors, element.Text)
	})
}

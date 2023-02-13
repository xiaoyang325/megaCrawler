package production

import (
	"megaCrawler/crawlers"

	"github.com/gocolly/colly/v2"
)

func init() {
	w := crawlers.Register("cepr", "经济与政策研究中心", "https://cepr.net/")
	w.SetStartingUrls([]string{"https://cepr.net/world/", "https://cepr.net/united-states/", "https://cepr.net/staff-experts/"})

	// index
	w.OnHTML(".next-btn.page-number", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		w.Visit(element.Attr("href"), crawlers.Index)
	})

	// 访问专家
	w.OnHTML(".staff-name>a", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		w.Visit(element.Attr("href"), crawlers.Expert)
	})

	// 专家姓名
	w.OnHTML("h5.staff-name", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		if ctx.PageType == crawlers.Expert {
			ctx.Name = element.Text
		}
	})

	// 专家头衔
	w.OnHTML("h5.staff-job", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		if ctx.PageType == crawlers.Expert {
			ctx.Title = element.Text
		}
	})

	// 专家描述
	w.OnHTML("#viewport > section > div > div > article > div.english > p", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Description += element.Text
	})

	// 访问报告
	w.OnHTML(".tax-post-con > h2 > a", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		w.Visit(element.Attr("href"), crawlers.Report)
	})

	// 报告标题
	w.OnHTML(".article-header-left-inner > h1", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Title = element.Text
	})

	// 标签
	w.OnHTML(".art-category2", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Tags = append(ctx.Tags, element.Text)
	})

	// 作者
	w.OnHTML(".author-desc-wrap > h3 > span >a", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Authors = append(ctx.Authors, element.Text)
	})

	// 正文
	w.OnHTML(".art-body-desc", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Content = element.Text
	})
}

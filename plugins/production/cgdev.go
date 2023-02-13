package production

import (
	"megaCrawler/crawlers"
	"strings"

	"github.com/gocolly/colly/v2"
)

func init() {
	w := crawlers.Register("cgdev", "全球发展中心", "https://www.cgdev.org/")
	w.SetStartingURLs([]string{"https://www.cgdev.org/section/experts", "https://www.cgdev.org/section/publications"})

	// 访问专家
	w.OnHTML(" div.view-content > div > div > a", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		w.Visit(element.Attr("href"), crawlers.Expert)
	})

	// 专家姓名
	w.OnHTML("div > div.title-wrapper > h1", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		if ctx.PageType == crawlers.Expert {
			ctx.Name = element.Text
		}
	})

	// 专家头衔
	w.OnHTML("div > div.title-wrapper > h5", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Title = element.Text
	})

	// 专家描述
	w.OnHTML("field field--name-field-desc4 field--type-text-long field--label-hidden w3-clear w3-bar-item field__item", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Description = element.Text
	})

	// 专家领域
	w.OnHTML("#expert-topics-listing > div > div > a", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Area += element.Text + " "
	})

	// 访问index
	w.OnHTML(".w3-button.pager__item", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		w.Visit(element.Attr("href"), crawlers.Index)
	})

	// 访问新闻
	w.OnHTML(".search-result-left > a", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		w.Visit(element.Attr("href"), crawlers.News)
	})

	// 尝试寻找下载pdf的按钮，并如果存在则将页面类型转换为报告
	w.OnHTML("a.gold", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		if strings.Contains(element.Attr("href"), ".pdf") {
			ctx.File = append(ctx.File, element.Attr("href"))
			ctx.PageType = crawlers.Report
		}
	})

	// 获取标题
	w.OnHTML(" div > div > h1", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		if ctx.PageType == crawlers.News || ctx.PageType == crawlers.Report {
			ctx.Title = element.Text
		}
	})

	// 作者
	w.OnHTML(".w3-section.field.field--name-field-authors.field--type-entity-reference.field--label-hidden.field__items > div> a", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Authors = append(ctx.Authors, element.Text)
	})

	// 时间
	w.OnHTML(".node-content-left>div.published-time", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.PublicationTime = element.Text
	})

	// 正文
	w.OnHTML(".field-body-blog-post.w3-bar-item.field__item", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Content = element.Text
	})
}

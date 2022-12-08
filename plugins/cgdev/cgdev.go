package cgdev

import (
	"github.com/gocolly/colly/v2"
	"megaCrawler/Crawler"
	"strings"
)

func init() {
	w := Crawler.Register("cgdev", "全球发展中心", "https://www.cgdev.org/")
	w.SetStartingUrls([]string{"https://www.cgdev.org/section/experts", "https://www.cgdev.org/section/publications"})

	//访问专家
	w.OnHTML(" div.view-content > div > div > a", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		w.Visit(element.Attr("href"), Crawler.Expert)
	})

	//专家姓名
	w.OnHTML("div > div.title-wrapper > h1", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		if ctx.PageType == Crawler.Expert {
			ctx.Name = element.Text
		}
	})

	//专家头衔
	w.OnHTML("div > div.title-wrapper > h5", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		ctx.Title = element.Text
	})

	//专家描述
	w.OnHTML("field field--name-field-desc4 field--type-text-long field--label-hidden w3-clear w3-bar-item field__item", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		ctx.Description = element.Text
	})

	//专家领域
	w.OnHTML("#expert-topics-listing > div > div > a", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		ctx.Area += element.Text + " "
	})

	//访问index
	w.OnHTML(".w3-button.pager__item", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		w.Visit(element.Attr("href"), Crawler.Index)
	})

	//访问新闻
	w.OnHTML(".search-result-left > a", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		w.Visit(element.Attr("href"), Crawler.News)
	})

	// 尝试寻找下载pdf的按钮，并如果存在则将页面类型转换为报告
	w.OnHTML("a.gold", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		if strings.Contains(element.Attr("href"), ".pdf") {
			ctx.File = append(ctx.File, element.Attr("href"))
			ctx.PageType = Crawler.Report
		}
	})

	//获取标题
	w.OnHTML(" div > div > h1", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		if ctx.PageType == Crawler.News || ctx.PageType == Crawler.Report {
			ctx.Title = element.Text
		}
	})

	//作者
	w.OnHTML(".w3-section.field.field--name-field-authors.field--type-entity-reference.field--label-hidden.field__items > div> a", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		ctx.Authors = append(ctx.Authors, element.Text)
	})

	//时间
	w.OnHTML(".node-content-left>div.published-time", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		ctx.PublicationTime = element.Text
	})

	//正文
	w.OnHTML(".field-body-blog-post.w3-bar-item.field__item", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		ctx.Content = element.Text
	})
}

package production

import (
	"megaCrawler/crawlers"

	"github.com/gocolly/colly/v2"
)

func init() {
	w := crawlers.Register("euiss", "欧盟安全研究所", "https://www.iss.europa.eu/")
	w.SetStartingUrls([]string{"https://www.iss.europa.eu/analyst-team",
		"https://www.iss.europa.eu/publications/reports"})

	// index
	w.OnHTML("li.arrow > a", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		w.Visit(element.Attr("href"), crawlers.Index)
	})

	// 访问专家
	w.OnHTML(".field-type-ds.field-label-hidden.field-wrapper > h2 > a", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		w.Visit(element.Attr("href"), crawlers.Expert)
	})

	// 访问报告
	w.OnHTML(".field-wrapper > span > a", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		if ctx.PageType != crawlers.Expert {
			w.Visit(element.Attr("href"), crawlers.Report)
		}
	})

	// 专家姓名,报告标题
	w.OnHTML("h1#page-title.title", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		if ctx.PageType == crawlers.Expert {
			ctx.Name = element.Text
		} else if ctx.PageType == crawlers.Report {
			ctx.Title = element.Text
		}
	})

	// 专家介绍
	w.OnHTML("views-field views-field-description", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Description = element.Text
	})

	// 作者
	w.OnHTML(".field-name-field-author > ul > li > a", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Authors = append(ctx.Authors, element.Text)
	})

	// 时间
	w.OnHTML("date-display-single", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.PublicationTime = element.Text
	})

	// 报告标签
	w.OnHTML(" main > div > div > div > div > div.publication-info > div > ul > li", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Tags = append(ctx.Tags, element.Text)
	})

	// 报告正文
	w.OnHTML("body field", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		if ctx.PageType == crawlers.Report {
			ctx.Content = element.Text
		}
	})

	// pdf
	w.OnHTML("span.file>a", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.File = append(ctx.File, element.Attr("href"))
	})
}

package euiss

import (
	"github.com/gocolly/colly/v2"
	"megaCrawler/Crawler"
)

func init() {
	w := Crawler.Register("euiss", "欧盟安全研究所", "https://www.iss.europa.eu/")
	w.SetStartingUrls([]string{"https://www.iss.europa.eu/analyst-team",
		"https://www.iss.europa.eu/publications/reports"})

	//index
	w.OnHTML("li.arrow > a", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		w.Visit(element.Attr("href"), Crawler.Index)
	})

	//访问专家
	w.OnHTML(".field-type-ds.field-label-hidden.field-wrapper > h2 > a", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		w.Visit(element.Attr("href"), Crawler.Expert)
	})

	//访问报告
	w.OnHTML(".field-wrapper > span > a", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		if ctx.PageType != Crawler.Expert {
			w.Visit(element.Attr("href"), Crawler.Report)
		}
	})

	//专家姓名,报告标题
	w.OnHTML("h1#page-title.title", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		if ctx.PageType == Crawler.Expert {
			ctx.Name = element.Text
		} else if ctx.PageType == Crawler.Report {
			ctx.Title = element.Text
		}
	})

	//专家介绍
	w.OnHTML("views-field views-field-description", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		ctx.Description = element.Text
	})

	//作者
	w.OnHTML(".field-name-field-author > ul > li > a", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		ctx.Authors = append(ctx.Authors, element.Text)
	})

	//时间
	w.OnHTML("date-display-single", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		ctx.PublicationTime = element.Text
	})

	//报告标签
	w.OnHTML(" main > div > div > div > div > div.publication-info > div > ul > li", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		ctx.Tags = append(ctx.Tags, element.Text)
	})

	//报告正文
	w.OnHTML("body field", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		if ctx.PageType == Crawler.Report {
			ctx.Content = element.Text
		}
	})

	//pdf
	w.OnHTML("span.file>a", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		ctx.File = append(ctx.File, element.Attr("href"))
	})
}

package production

import (
	"megaCrawler/crawlers"
	"strings"

	"github.com/gocolly/colly/v2"
)

func init() {
	w := crawlers.Register("oxfordenergy", "牛津能源研究所", "https://www.oxfordenergy.org/")
	w.SetStartingUrls([]string{"https://www.oxfordenergy.org/about/staff/", "https://www.oxfordenergy.org/"})

	// index
	w.OnHTML("#menu-item-44156 > ul > li > a", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		w.Visit(element.Attr("href"), crawlers.Index)
	})
	w.OnHTML("#publications-results > span > nav > a", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		w.Visit(element.Attr("href"), crawlers.Index)
	})

	// 访问报告
	w.OnHTML(" ul > li > div > div > div > div.row > div > a", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		w.Visit(element.Attr("href"), crawlers.Report)
	})

	// 获取报告标题,人物姓名
	w.OnHTML("div.small-12.medium-8.columns.side-border-right > h1", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		if ctx.PageType == crawlers.Report {
			ctx.Title = element.Text
		} else if ctx.PageType == crawlers.Expert {
			ctx.Name = element.Text
		}
	})

	// 获取报告分类
	w.OnHTML(" div > div.small-10.columns > p:nth-child(2) > a", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.CategoryText += element.Text + ", "
	})

	// 获取报告标签
	w.OnHTML(" div > div.small-10.columns > p:nth-child(4) > a", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Tags = append(ctx.Tags, element.Text)
	})

	// 获取报告作者
	w.OnHTML(" div.small-12.medium-8.columns.side-border-right > p.authors > a", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Authors = append(ctx.Authors, element.Text)
	})

	// 获取报告正文,人物描述
	w.OnHTML(" div > div.small-12.medium-8.columns.side-border-right > p", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		if ctx.PageType == crawlers.Report {
			ctx.Content = element.Text
		} else if ctx.PageType == crawlers.Expert {
			ctx.Description = element.Text
		}
	})

	// 获取pdf
	w.OnHTML(" div > div.medium-4.columns > p > a", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		if strings.Contains(element.Attr("href"), ".pdf") {
			ctx.File = append(ctx.File, element.Attr("href"))
		}
	})

	// 访问人物
	w.OnHTML(" div > div.small-12.medium-9.columns.side-border-right.page-padding > div> div > ul > li > a", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		w.Visit(element.Attr("href"), crawlers.Expert)
	})

	// 获取人物头衔
	w.OnHTML(" div.small-12.side-border-right.medium-8.columns > h3", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Title = element.Text
	})

	// 获取人物联系方式
	w.OnHTML("body > div.off-canvas-wrap > div > div.page-content > div > div.small-12.medium-4.columns > ul > li > a", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Email = element.Text
	})
}

package plugins

import (
	"github.com/gocolly/colly/v2"
	"megaCrawler/Crawler"
	"regexp"
	"strings"
)

func init() {
	w := Crawler.Register("thecic", "国际理事会", "https://thecic.org/")

	w.SetStartingUrls([]string{"https://thecic.org/research-publications/behind-the-headlines/"})

	// 尝试寻找下载pdf的按钮，并如果存在则将页面类型转换为报告
	w.OnHTML("a.button", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		if strings.Contains(element.Attr("href"), ".pdf") {
			ctx.File = append(ctx.File, element.Attr("href"))
			ctx.PageType = Crawler.Report
		}
	})
	// 从index访问新闻
	w.OnHTML("h2.entry-title>a", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		w.Visit(element.Attr("href"), Crawler.News)
	})

	// report.title
	w.OnHTML(" div.et_pb_title_container > h1", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		ctx.Title = element.Text
	})
	//report.publish time
	w.OnHTML("span.updated", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		ctx.PublicationTime = element.Text
	})

	// report.author
	w.OnHTML("div.et_pb_text_inner>p", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		authorName := element.Text
		authorRegex, _ := regexp.Compile("By: ([\\w ]+)")
		authorMatch := authorRegex.FindStringSubmatch(element.Text)
		if len(authorMatch) == 2 {
			authorName = authorMatch[1]
		}

		ctx.Authors = append(ctx.Authors, authorName)
	})

	// report .content
	w.OnHTML("div.et_pb_row_1", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		ctx.Content = element.Text
	})

}

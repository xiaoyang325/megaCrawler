package csis_or_id

import (
	"github.com/gocolly/colly/v2"
	"megaCrawler/Crawler"
	"strings"
)

func init() {
	w := Crawler.Register("csis_or_id", "战略与国际研究中心", "https://csis.or.id/")

	w.SetStartingUrls([]string{
		"https://csis.or.id/projects/",
		"https://csis.or.id/publications/books/",
		"https://csis.or.id/publications/commentaries/",
		"https://csis.or.id/publications/policy-paper-series/",
		"https://csis.or.id/publications/press-release/",
		"https://csis.or.id/publications/research-report/",
		"https://csis.or.id/publications/working-paper/",
	})

	// 访问 Report 从 Index
	w.OnHTML(`div.post-image > a`, func(element *colly.HTMLElement, ctx *Crawler.Context) {
		w.Visit(element.Attr("href"), Crawler.Report)
	})

	// 获取 Title
	w.OnHTML(`[class="page-title text-black text-left"] > h2`, func(element *colly.HTMLElement, ctx *Crawler.Context) {
		ctx.Title = strings.TrimSpace(element.Text)
	})

	// 获取 CategoryText
	w.OnHTML(`[class="page-title text-black text-left"] > h5`, func(element *colly.HTMLElement, ctx *Crawler.Context) {
		ctx.CategoryText = strings.TrimSpace(element.Text)
	})

	// 获取 Authors
	w.OnHTML(`[class="page-title text-black text-left"] > .row > div > a > h4`, func(element *colly.HTMLElement, ctx *Crawler.Context) {
		ctx.Authors = append(ctx.Authors, strings.TrimSpace(element.Text))
	})

	// 获取 Content
	w.OnHTML(`.post-item .text-justify`, func(element *colly.HTMLElement, ctx *Crawler.Context) {
		ctx.Content = strings.TrimSpace(element.Text)
	})

	// 获取 File
	w.OnHTML(`.btn-csis > a`, func(element *colly.HTMLElement, ctx *Crawler.Context) {
		if strings.Contains(element.Attr("href"), ".pdf") {
			ctx.File = append(ctx.File, element.Attr("href"))
		}
	})
}

package grc

import (
	"github.com/gocolly/colly/v2"
	"megaCrawler/Crawler"
	"strconv"
	"strings"
)

func init() {
	w := Crawler.Register("grc", "海湾研究中心", "https://www.grc.net/")

	w.SetStartingUrls([]string{
		"https://www.grc.net/commentary-and-analysis",
		"https://www.grc.net/country-updates",
	})

	// 访问 Report 从 Index 通过 SubContext
	w.OnHTML(`.gdlr-core-course-item-info > .gdlr-core-tail`,
		func(element *colly.HTMLElement, ctx *Crawler.Context) {
			sub_ctx := ctx.CreateSubContext()

			sub_ctx.PageType = Crawler.Report
			sub_ctx.CategoryText = "Country Updates"
			sub_ctx.PublicationTime = strings.Replace(element.ChildText("span:nth-child(1)"), "Download PDF", "", 1)
			sub_ctx.Title = element.ChildText(".gdlr-core-head")
			sub_ctx.Description = element.ChildText("div:nth-child(5)")

			raw_str := element.ChildText(".analysis_authors")
			str := strings.Replace(raw_str, element.ChildText(".analysis_authors > strong"), "", 1)
			str = strings.Replace(str, "*", "", 1)
			sub_ctx.Authors = append(sub_ctx.Authors, strings.TrimSpace(str))

			sub_ctx.File = append(sub_ctx.File, element.ChildAttr(`a[target="_blank"]`, "href"))
		})

	// 访问 Report 从 Index
	w.OnHTML(`.gdlr-core-head > a`,
		func(element *colly.HTMLElement, ctx *Crawler.Context) {
			w.Visit(element.Attr("href"), Crawler.Report)
		})

	// 获取 Title
	w.OnHTML(`[class="gdlr-core-title-item-title gdlr-core-skin-title "]`,
		func(element *colly.HTMLElement, ctx *Crawler.Context) {
			ctx.Title = strings.TrimSpace(element.Text)
		})

	// 获取 PublicationTime
	w.OnHTML(`.gdlr-core-pbf-element > div > span`,
		func(element *colly.HTMLElement, ctx *Crawler.Context) {
			ctx.PublicationTime = strings.TrimSpace(element.Text)
		})

	// 获取 CategoryText
	w.OnHTML(`h1.kingster-page-title`,
		func(element *colly.HTMLElement, ctx *Crawler.Context) {
			ctx.CategoryText = strings.TrimSpace(element.Text)
		})

	// 获取 Authors
	w.OnHTML(`.gdlr-core-pbf-element > div > .analysis_authors`,
		func(element *colly.HTMLElement, ctx *Crawler.Context) {
			str := strings.Replace(element.Text, element.ChildText("strong"), "", 1)
			str = strings.Replace(str, "*", "", 1)
			ctx.Authors = append(ctx.Authors, strings.TrimSpace(str))
		})

	// 获取 Content
	w.OnHTML(`div.gdlr-core-text-box-item-content`,
		func(element *colly.HTMLElement, ctx *Crawler.Context) {
			ctx.Content = strings.TrimSpace(element.Text)
		})

	// 获取 File
	w.OnHTML(`a[class="gdlr-core-button  gdlr-core-button-solid gdlr-core-button-no-border"][target="_blank"]`,
		func(element *colly.HTMLElement, ctx *Crawler.Context) {
			ctx.File = append(ctx.File, element.Attr("href"))
		})
}

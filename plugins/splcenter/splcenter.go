package splcenter

import (
	"github.com/gocolly/colly/v2"
	"megaCrawler/Crawler"
	"strings"
)

func init() {
	w := Crawler.Register("splcenter", "南方贫困法律中心", "https://www.splcenter.org/")

	w.SetStartingUrls([]string{
		"https://www.splcenter.org/features-stories/",
		"https://www.splcenter.org/fighting-hate/extremist-files/",
		"https://www.splcenter.org/hatewatch/",
		"https://www.splcenter.org/year-hate-extremism-2021",
		"https://www.splcenter.org/issues/hate-and-extremism",
		"https://www.splcenter.org/issues/childrens-rights",
		"https://www.splcenter.org/issues/immigrant-justice",
		"https://www.splcenter.org/issues/lgbtq-rights",
		"https://www.splcenter.org/our-issues/economic-justice",
		"https://www.splcenter.org/issues/mass-incarceration",
		"https://www.splcenter.org/our-issues/voting-rights",
	})

	// 从频道入口访问更多信息的 Index
	w.OnHTML(".more-link a",
		func(element *colly.HTMLElement, ctx *Crawler.Context) {
			w.Visit(element.Attr("href"), Crawler.Index)
		})

	// 从翻页器获取下一页 Index 并访问
	w.OnHTML(".pager-next > a",
		func(element *colly.HTMLElement, ctx *Crawler.Context) {
			w.Visit(element.Attr("href"), Crawler.Index)
		})

	// 从 Index 访问 News
	w.OnHTML(`#main-content div[class="field-item even "] > h1 > a`,
		func(element *colly.HTMLElement, ctx *Crawler.Context) {
			w.Visit(element.Attr("href"), Crawler.News)
		})

	// 获取 Title
	w.OnHTML(".group-header h1",
		func(element *colly.HTMLElement, ctx *Crawler.Context) {
			ctx.Title = strings.TrimSpace(element.Text)
		})

	// 获取 Publication Time
	w.OnHTML(".date-display-single",
		func(element *colly.HTMLElement, ctx *Crawler.Context) {
			ctx.PublicationTime = strings.TrimSpace(element.Text)
		})

	// 获取 Authors（情况一）
	w.OnHTML(`div[class="field field-name-field-person field-type-entityreference field-label-hidden"] .field-items a`,
		func(element *colly.HTMLElement, ctx *Crawler.Context) {
			ctx.Authors = append(ctx.Authors, strings.TrimSpace(element.Text))
		})

	// 获取 Authors（情况二）
	w.OnHTML(`div[class="field field-name-field-byline field-type-text field-label-hidden"] .field-items div`,
		func(element *colly.HTMLElement, ctx *Crawler.Context) {
			ctx.Authors = append(ctx.Authors, strings.TrimSpace(element.Text))
		})

	// 获取 Authors（情况三）
	w.OnHTML(`div[class="field field-name-title-field field-type-text field-label-hidden"] .field-items a`,
		func(element *colly.HTMLElement, ctx *Crawler.Context) {
			ctx.Authors = append(ctx.Authors, strings.TrimSpace(element.Text))
		})

	// 获取 Content（情况一）
	w.OnHTML("#group-content-container .field-items> div",
		func(element *colly.HTMLElement, ctx *Crawler.Context) {
			ctx.Content = element.Text
		})

	// 获取 Content（情况二）
	w.OnHTML(".group-content .field-items> div",
		func(element *colly.HTMLElement, ctx *Crawler.Context) {
			ctx.Content = element.Text
		})

	// 获取 Related Documents 中的 File（/case-docket）
	w.OnHTML(".file > a",
		func(element *colly.HTMLElement, ctx *Crawler.Context) {
			ctx.File = append(ctx.File, strings.TrimSpace(element.Text))
		})
}

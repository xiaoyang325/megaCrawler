package ids

import (
	"github.com/gocolly/colly/v2"
	"megaCrawler/Crawler"
	"strings"
)

func init() {
	w := Crawler.Register("ids", "发展研究所", "https://www.ids.ac.uk/")
	
	w.SetStartingUrls([]string{
		"https://www.ids.ac.uk/about/clusters-and-teams/",
		"https://www.ids.ac.uk/events/",
		"https://www.ids.ac.uk/news-and-opinion/",
		"https://www.ids.ac.uk/learn-at-ids/learning-for-development/",
		"https://www.ids.ac.uk/learn-at-ids/",
	})

	// 访问 See all 从频道入口
	w.OnHTML(`a[title="See all"]`,
		func(element *colly.HTMLElement, ctx *Crawler.Context) {
			w.Visit(element.Attr("href"), Crawler.Index)
		})

	// 访问下一页 Index
	w.OnHTML(`a[title="Next page"]`,
		func(element *colly.HTMLElement, ctx *Crawler.Context) {
			w.Visit(element.Attr("href"), Crawler.Index)
		})

	// 访问 Report 从 Index
	w.OnHTML(`h4[class="c-content-item__heading ts-heading-4"] > a`,
		func(element *colly.HTMLElement, ctx *Crawler.Context) {
			w.Visit(element.Attr("href"), Crawler.Report)
		})

	// 访问 Report 从 Index (Type 2)
	w.OnHTML(`.c-featured-pages__items > a`,
		func(element *colly.HTMLElement, ctx *Crawler.Context) {
			w.Visit(element.Attr("href"), Crawler.Report)
		})

	// 获取 Title
	w.OnHTML(`h1[class="c-single-header__heading ts-heading-2"]`,
		func(element *colly.HTMLElement, ctx *Crawler.Context) {
			ctx.Title = strings.TrimSpace(element.Text)
		})

	// 获取 Title (Type 2)
	w.OnHTML(`h1.ts-heading-1`,
		func(element *colly.HTMLElement, ctx *Crawler.Context) {
			ctx.Title = strings.TrimSpace(element.Text)
		})

	// 获取 Title (Type 3)
	w.OnHTML(`h1[class="c-learn-banner__degree ts-heading-4"]`,
		func(element *colly.HTMLElement, ctx *Crawler.Context) {
			ctx.Title = strings.TrimSpace(element.Text)
		})

	// 获取 Title (Type 4)
	w.OnHTML(`h1[class="c-single-bio__heading ts-heading-3"]`,
		func(element *colly.HTMLElement, ctx *Crawler.Context) {
			ctx.Title = strings.TrimSpace(element.Text)
		})

	// 获取 Description
	w.OnHTML(`div.entry-summary > p`,
		func(element *colly.HTMLElement, ctx *Crawler.Context) {
			ctx.Description = strings.TrimSpace(element.Text)
		})

	// 获取 PublicationTime
	w.OnHTML(`.c-date__heading`,
		func(element *colly.HTMLElement, ctx *Crawler.Context) {
			ctx.PublicationTime = strings.TrimSpace(element.Text)
		})

	// 获取 PublicationTime (Type 2)
	w.OnHTML(`p[class="c-basic-single-meta__item ts-heading-5"]`,
		func(element *colly.HTMLElement, ctx *Crawler.Context) {
			time_str := strings.Replace(element.Text, "Published on", "", 1)
			time_str = strings.TrimSpace(time_str)
			ctx.PublicationTime = strings.TrimSpace(time_str)
		})

	// 获取 CategoryText
	w.OnHTML(`a[class="c-single-header__sub-heading c-single-header__sub-heading--link ts-heading-3"]`,
		func(element *colly.HTMLElement, ctx *Crawler.Context) {
			ctx.CategoryText = strings.TrimSpace(element.Text)
		})

	// 获取 CategoryText (Type 2)
	w.OnHTML(`h2[class="c-meta-block__heading ts-heading-6"]`,
		func(element *colly.HTMLElement, ctx *Crawler.Context) {
			ctx.CategoryText = strings.TrimSpace(element.Text)
		})

	// 获取 CategoryText (Type 3)
	w.OnHTML(`[class="c-breadcrumbs__link-text ts-heading-5"]`,
		func(element *colly.HTMLElement, ctx *Crawler.Context) {
			ctx.CategoryText = strings.TrimSpace(element.Text)
		})

	// 获取 Authors
	w.OnHTML(`div[class="c-basic-single-meta "] .c-person__link`,
		func(element *colly.HTMLElement, ctx *Crawler.Context) {
			ctx.Authors = append(ctx.Authors, strings.TrimSpace(element.Text))
		})

	// 获取 Content
	w.OnHTML(`.o-content-from-editor`,
		func(element *colly.HTMLElement, ctx *Crawler.Context) {
			if (!strings.Contains(ctx.url, "/learn-at-ids/")) {
				ctx.Content = strings.TrimSpace(element.Text)
			}
		})

	// 获取 Content (Type 2)
	w.OnHTML(`.o-layout .u-container:nth-child(1) .o-layout__eight-col`,
		func(element *colly.HTMLElement, ctx *Crawler.Context) {
			if (strings.Contains(ctx.url, "/learn-at-ids/")) {
				ctx.Content = strings.TrimSpace(element.Text)
			}
		})

	// 获取 Tags
	w.OnHTML(`div[class="c-meta-block__meta-item "] > a`,
		func(element *colly.HTMLElement, ctx *Crawler.Context) {
			ctx.Tags = append(ctx.Tags, strings.TrimSpace(element.Text))
		})
}

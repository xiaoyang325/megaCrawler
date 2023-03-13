package production

import (
	"strings"

	"megaCrawler/crawlers"
	"megaCrawler/extractors"

	"github.com/gocolly/colly/v2"
)

func init() {
	w := crawlers.Register("ids", "发展研究所", "https://www.ids.ac.uk/")

	w.SetStartingURLs([]string{
		"https://www.ids.ac.uk/about/clusters-and-teams/",
		"https://www.ids.ac.uk/events/",
		"https://www.ids.ac.uk/news-and-opinion/",
		"https://www.ids.ac.uk/learn-at-ids/learning-for-development/",
		"https://www.ids.ac.uk/learn-at-ids/",
	})

	// 访问 See all 从频道入口
	w.OnHTML(`a[title="See all"]`, func(element *colly.HTMLElement, ctx *crawlers.Context) {
		w.Visit(element.Attr("href"), crawlers.Index)
	})

	// 访问下一页 Index
	w.OnHTML(`a[title="Next page"]`, func(element *colly.HTMLElement, ctx *crawlers.Context) {
		w.Visit(element.Attr("href"), crawlers.Index)
	})

	// 访问 Report 从 Index
	w.OnHTML(`h4[class="c-content-item__heading ts-heading-4"] > a`, func(element *colly.HTMLElement, ctx *crawlers.Context) {
		w.Visit(element.Attr("href"), crawlers.Report)
	})

	// 访问 Report 从 Index (Type 2)
	w.OnHTML(`.c-featured-pages__items > a`, func(element *colly.HTMLElement, ctx *crawlers.Context) {
		w.Visit(element.Attr("href"), crawlers.Report)
	})

	w.OnHTML("html", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		extractors.Titles(ctx, element)
	})

	// 获取 Description
	w.OnHTML(`div.entry-summary > p`, func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Description = strings.TrimSpace(element.Text)
	})

	// 获取 PublicationTime
	w.OnHTML(`.c-date__heading`, func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.PublicationTime = strings.TrimSpace(element.Text)
	})

	// 获取 PublicationTime (Type 2)
	w.OnHTML(`p[class="c-basic-single-meta__item ts-heading-5"]`, func(element *colly.HTMLElement, ctx *crawlers.Context) {
		timeStr := strings.Replace(element.Text, "Published on", "", 1)
		timeStr = strings.TrimSpace(timeStr)
		ctx.PublicationTime = strings.TrimSpace(timeStr)
	})

	// 获取 CategoryText
	w.OnHTML(`a[class="c-single-header__sub-heading c-single-header__sub-heading--link ts-heading-3"]`, func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.CategoryText = strings.TrimSpace(element.Text)
	})

	// 获取 CategoryText (Type 2)
	w.OnHTML(`h2[class="c-meta-block__heading ts-heading-6"]`, func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.CategoryText = strings.TrimSpace(element.Text)
	})

	// 获取 CategoryText (Type 3)
	w.OnHTML(`[class="c-breadcrumbs__link-text ts-heading-5"]`, func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.CategoryText = strings.TrimSpace(element.Text)
	})

	// 获取 Authors
	w.OnHTML(`.c-basic-single-meta__author .c-person__link`, func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Authors = append(ctx.Authors, strings.TrimSpace(element.Text))
	})
	w.OnHTML(`aside .c-person__link`, func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Authors = append(ctx.Authors, strings.TrimSpace(element.Text))
	})

	// 获取 Content
	w.OnHTML(`.o-content-from-editor`, func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Content = strings.TrimSpace(element.Text)
	})

	// 获取 Tags
	w.OnHTML(`div[class="c-meta-block__meta-item "] > a`, func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Tags = append(ctx.Tags, strings.TrimSpace(element.Text))
	})
}

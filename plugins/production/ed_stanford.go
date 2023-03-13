package production

import (
	"strings"

	"megaCrawler/crawlers"

	"github.com/gocolly/colly/v2"
)

func partEd(w *crawlers.WebsiteEngine) {
	// 从频道入口获取 Index 并访问
	w.OnHTML(".text-center > a", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		w.Visit(element.Attr("href"), crawlers.Index)
	})

	// 从翻页器获取下一页 Index 并访问
	w.OnHTML(`a[title="Go to next page"]`, func(element *colly.HTMLElement, ctx *crawlers.Context) {
		w.Visit(element.Attr("href"), crawlers.Index)
	})

	// 从 Index 访问 Report（/news）
	w.OnHTML(".news-card .nc-title > a", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		w.Visit(element.Attr("href"), crawlers.Report)
	})

	// 从 Index 访问 Report（/podcast）
	w.OnHTML(".podcast-card .pc-title > a", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		w.Visit(element.Attr("href"), crawlers.Report)
	})

	// 获取 Title
	w.OnHTML(".nh-title", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Title = element.Text
	})

	// 获取 SubTitle
	w.OnHTML(".nh-subtitle .field-items > div", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.SubTitle = element.Text
	})

	// 获取 Publication Time
	w.OnHTML(".nh-byline .date-display-single", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.PublicationTime = element.Text
	})

	// 获取 Authors
	w.OnHTML(`div[class="field field-name-field-news-source field-type-text field-label-hidden"] div div`, func(element *colly.HTMLElement, ctx *crawlers.Context) {
		// "By Adam Hadhazy" -> "Adam had"
		rawString := strings.ReplaceAll(element.Text, "By", "")
		rawString = strings.TrimSpace(rawString)
		ctx.Authors = append(ctx.Authors, rawString)
	})

	// 获取 Content
	w.OnHTML(`div[class="content sof45"]`, func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Content = element.Text
	})

	// 获取 Tags
	w.OnHTML(".nh-tag .field-items a", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Tags = append(ctx.Tags, strings.TrimSpace(element.Text))
	})
}

package production

import (
	"strings"

	"megaCrawler/crawlers"

	"github.com/gocolly/colly/v2"
)

func partNews(w *crawlers.WebsiteEngine) {
	// 从翻页器获取更多（« Older stories）并访问
	w.OnHTML(`div[class="btn btn-su-alert"] a`, func(element *colly.HTMLElement, ctx *crawlers.Context) {
		w.Visit(element.Attr("href"), crawlers.Index)
	})

	// 从 Index 访问 Report
	w.OnHTML(".card-content h3 a", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		w.Visit(element.Attr("href"), crawlers.Report)
	})

	// 获取 Title
	w.OnHTML("#story-head div h1", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Title = element.Text
	})

	// 获取 Description
	w.OnHTML(".excerpt", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Description = element.Text
	})

	// 获取 Publication Time
	w.OnHTML("#story-head div time", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.PublicationTime = element.Text
	})

	// 获取 Authors
	w.OnHTML(".byline", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		// "By Adam Hadhazy" -> "Adam had"
		rawString := strings.ReplaceAll(element.Text, "By", "")
		rawString = strings.TrimSpace(rawString)
		ctx.Authors = append(ctx.Authors, rawString)
	})

	// 获取 Content
	w.OnHTML("#story-content", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Content = element.ChildText("p")
	})

	// 获取 Tags
	w.OnHTML(`div[class="btn btn-category"] a`, func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Tags = append(ctx.Tags, strings.TrimSpace(element.Text))
	})
}

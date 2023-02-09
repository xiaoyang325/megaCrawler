package production

import (
	"github.com/gocolly/colly/v2"
	"megaCrawler/Crawler"
	"strings"
)

func partNews(w *Crawler.WebsiteEngine) {
	// 从翻页器获取更多（« Older stories）并访问
	w.OnHTML(`div[class="btn btn-su-alert"] a`, func(element *colly.HTMLElement, ctx *Crawler.Context) {
		w.Visit(element.Attr("href"), Crawler.Index)
	})

	// 从 Index 访问 Report
	w.OnHTML(".card-content h3 a", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		w.Visit(element.Attr("href"), Crawler.Report)
	})

	// 获取 Title
	w.OnHTML("#story-head div h1", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		ctx.Title = element.Text
	})

	// 获取 Description
	w.OnHTML(".excerpt", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		ctx.Description = element.Text
	})

	// 获取 Publication Time
	w.OnHTML("#story-head div time", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		ctx.PublicationTime = element.Text
	})

	// 获取 Authors
	w.OnHTML(".byline", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		// "By Adam Hadhazy" -> "Adam had"
		raw_string := strings.ReplaceAll(element.Text, "By", "")
		raw_string = strings.TrimSpace(raw_string)
		ctx.Authors = append(ctx.Authors, raw_string)
	})

	// 获取 Content
	w.OnHTML("#story-content", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		ctx.Content = element.ChildText("p")
	})

	// 获取 Tags
	w.OnHTML(`div[class="btn btn-category"] a`, func(element *colly.HTMLElement, ctx *Crawler.Context) {
		ctx.Tags = append(ctx.Tags, strings.TrimSpace(element.Text))
	})
}

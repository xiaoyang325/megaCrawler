package production

import (
	"strings"

	"megaCrawler/crawlers"

	"github.com/gocolly/colly/v2"
)

func init() {
	w := crawlers.Register("iidnet", "国际对话倡议", "https://iidnet.org/")

	w.SetStartingURLs([]string{
		"https://iidnet.org/press-releases-archive/",
		"https://iidnet.org/resources-archive/",
		"https://iidnet.org/news-archive/",
	})

	// 访问下一页 Index
	w.OnHTML(`[class="pagination clearfix"] > .alignleft > a`, func(element *colly.HTMLElement, ctx *crawlers.Context) {
		w.Visit(element.Attr("href"), crawlers.Index)
	})

	// 访问 News 从 Index
	w.OnHTML(`.entry-title > a`, func(element *colly.HTMLElement, ctx *crawlers.Context) {
		w.Visit(element.Attr("href"), crawlers.News)
	})

	// 获取 Title
	w.OnHTML(`h1.entry-title`, func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Title = strings.TrimSpace(element.Text)
	})

	// 获取 PublicationTime
	w.OnHTML(`span.published`, func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.PublicationTime = strings.TrimSpace(element.Text)
	})

	// 获取 Authors
	w.OnHTML(`[class="author vcard"] > a`, func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Authors = append(ctx.Authors, strings.TrimSpace(element.Text))
	})

	// 获取 Content
	w.OnHTML(`[class="et_pb_row et_pb_row_0"] div.et_pb_text_inner`, func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Content = strings.TrimSpace(element.Text)
	})

	// 获取 Content
	w.OnHTML(`#content-area .entry-content`, func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Content = strings.TrimSpace(element.Text)
	})

	// 获取 Tags
	w.OnHTML(`.post-meta > a[rel="category tag"]`, func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Tags = append(ctx.Tags, strings.TrimSpace(element.Text))
	})
}

package production

import (
	"regexp"
	"strings"

	"megaCrawler/crawlers"

	"github.com/gocolly/colly/v2"
)

func init() {
	w := crawlers.Register("ait", "美国在台协会", "https://www.ait.org.tw/")

	w.SetStartingURLs([]string{
		"https://www.ait.org.tw/news/",
		"https://www.ait.org.tw/category/speeches/",
		"https://www.ait.org.tw/category/news/",
	})

	// 从 频道 访问 Read All
	w.OnHTML(`.custom_button > a`, func(element *colly.HTMLElement, ctx *crawlers.Context) {
		w.Visit(element.Attr("href"), crawlers.Index)
	})

	// 从 Index 访问 News
	w.OnHTML(`.entry-title > a`, func(element *colly.HTMLElement, ctx *crawlers.Context) {
		w.Visit(element.Attr("href"), crawlers.News)
	})

	// 从 Index 访问 Next Page
	w.OnHTML(`a[class="next page-numbers"]`, func(element *colly.HTMLElement, ctx *crawlers.Context) {
		w.Visit(element.Attr("href"), crawlers.Index)
	})

	// 获取 Title
	w.OnHTML(`div[class="internal-title-header medium"]`, func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Title = strings.TrimSpace(element.Text)
	})

	// 获取 Publication Time
	w.OnHTML(".main-content-wrapper", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		// 尝试从 .datemeta 获取，如果为空则使用正则处理另一个标签
		meta := element.ChildText(".datemeta")
		if meta == "" {
			str := element.ChildText(`p[style="text-align: right;"]`)
			reg := regexp.MustCompile(`[A-Z]{2}-\d+`)
			str = reg.ReplaceAllString(str, "")
			ctx.PublicationTime = strings.TrimSpace(str)
		} else {
			ctx.PublicationTime = meta
		}
	})

	// 获取 Authors
	w.OnHTML(`.authormeta > a`, func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Authors = append(ctx.Authors, strings.TrimSpace(element.Text))
	})

	// 获取 Content
	w.OnHTML(`div[class="paragraph alignwide container"] > div`, func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Content = element.Text
	})

	// 获取 Tags
	w.OnHTML(`.tagsmeta a[rel="tag"]`, func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Tags = append(ctx.Tags, element.Text)
	})

	// 获取 Location
	w.OnHTML(`.locationmeta > a`, func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Location = element.Text
	})

	// 获取 CategoryText
	w.OnHTML(`.catsmeta > a`, func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.CategoryText = element.Text
	})
}

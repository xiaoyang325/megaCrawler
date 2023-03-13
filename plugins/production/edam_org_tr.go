package production

import (
	"strings"

	"megaCrawler/crawlers"

	"github.com/gocolly/colly/v2"
)

func init() {
	w := crawlers.Register("edam_org_tr", "Center for Economics and Foreign Policy Studies",
		"https://edam.org.tr/")

	w.SetStartingURLs([]string{
		"https://edam.org.tr/en/kategori/foreign-policy-security/",
		"https://edam.org.tr/en/kategori/cyber-policy/",
		"https://edam.org.tr/en/kategori/economics-globalization/",
		"https://edam.org.tr/en/kategori/edam-blog-en/",
		"https://edam.org.tr/en/kategori/defense-intelligence-sentinel/",
	})

	// 访问下一页 Index
	w.OnHTML(`.main-pagination > a[class="next page-numbers"]`, func(element *colly.HTMLElement, ctx *crawlers.Context) {
		w.Visit(element.Attr("href"), crawlers.Index)
	})

	// 访问 News 从 Index
	w.OnHTML(`a.post-title`, func(element *colly.HTMLElement, ctx *crawlers.Context) {
		w.Visit(element.Attr("href"), crawlers.News)
	})

	// 获取 Title
	w.OnHTML(`h1.post-title`, func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Title = strings.TrimSpace(element.Text)
	})

	// 获取 PublicationTime
	w.OnHTML(`time.post-date`, func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.PublicationTime = strings.TrimSpace(element.Text)
	})

	// 获取 CategoryText
	w.OnHTML(`[class="category cf"] a`, func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.CategoryText = strings.TrimSpace(element.Text)
	})

	// 获取 Authors
	w.OnHTML(`.posted-by > a`, func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Authors = append(ctx.Authors, strings.TrimSpace(element.Text))
	})

	// 获取 Content
	w.OnHTML(`div[class="post-content post-dynamic"]`, func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Content = strings.TrimSpace(element.Text)
	})

	// 获取 Tags
	w.OnHTML(`.tagcloud a[rel="tag"]`, func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Tags = append(ctx.Tags, strings.TrimSpace(element.Text))
	})

	// 获取 File
	w.OnHTML(`[class="post-content post-dynamic"] p[style="text-align: center;"] > a[target="_blank"]`, func(element *colly.HTMLElement, ctx *crawlers.Context) {
		if strings.Contains(element.Attr("href"), ".pdf") {
			ctx.File = append(ctx.File, element.Attr("href"))
			ctx.PageType = crawlers.Report
		}
	})
}

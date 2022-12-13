package edam_org_tr

import (
	"github.com/gocolly/colly/v2"
	"megaCrawler/Crawler"
	"strings"
)

func init() {
	w := Crawler.Register("edam_org_tr", "Center for Economics and Foreign Policy Studies",
		"https://edam.org.tr/")

	w.SetStartingUrls([]string{
		"https://edam.org.tr/en/kategori/foreign-policy-security/",
		"https://edam.org.tr/en/kategori/cyber-policy/",
		"https://edam.org.tr/en/kategori/economics-globalization/",
		"https://edam.org.tr/en/kategori/edam-blog-en/",
		"https://edam.org.tr/en/kategori/defense-intelligence-sentinel/",
	})

	// 访问下一页 Index
	w.OnHTML(`.main-pagination > a[class="next page-numbers"]`, func(element *colly.HTMLElement, ctx *Crawler.Context) {
		w.Visit(element.Attr("href"), Crawler.Index)
	})

	// 访问 News 从 Index
	w.OnHTML(`a.post-title`, func(element *colly.HTMLElement, ctx *Crawler.Context) {
		w.Visit(element.Attr("href"), Crawler.News)
	})

	// 获取 Title
	w.OnHTML(`h1.post-title`, func(element *colly.HTMLElement, ctx *Crawler.Context) {
		ctx.Title = strings.TrimSpace(element.Text)
	})

	// 获取 PublicationTime
	w.OnHTML(`time.post-date`, func(element *colly.HTMLElement, ctx *Crawler.Context) {
		ctx.PublicationTime = strings.TrimSpace(element.Text)
	})

	// 获取 CategoryText
	w.OnHTML(`[class="category cf"] a`, func(element *colly.HTMLElement, ctx *Crawler.Context) {
		ctx.CategoryText = strings.TrimSpace(element.Text)
	})

	// 获取 Authors
	w.OnHTML(`.posted-by > a`, func(element *colly.HTMLElement, ctx *Crawler.Context) {
		ctx.Authors = append(ctx.Authors, strings.TrimSpace(element.Text))
	})

	// 获取 Content
	w.OnHTML(`div[class="post-content post-dynamic"]`, func(element *colly.HTMLElement, ctx *Crawler.Context) {
		ctx.Content = strings.TrimSpace(element.Text)
	})

	// 获取 Tags
	w.OnHTML(`.tagcloud a[rel="tag"]`, func(element *colly.HTMLElement, ctx *Crawler.Context) {
		ctx.Tags = append(ctx.Tags, strings.TrimSpace(element.Text))
	})

	// 获取 File
	w.OnHTML(`[class="post-content post-dynamic"] p[style="text-align: center;"] > a[target="_blank"]`, func(element *colly.HTMLElement, ctx *Crawler.Context) {
		if strings.Contains(element.Attr("href"), ".pdf") {
			ctx.File = append(ctx.File, element.Attr("href"))
			ctx.PageType = Crawler.Report
		}
	})
}

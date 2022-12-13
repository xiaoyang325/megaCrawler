package centerforsecuritypolicy

import (
	"github.com/gocolly/colly/v2"
	"megaCrawler/Crawler"
	"strings"
)

func init() {
	w := Crawler.Register("centerforsecuritypolicy", "安全政策中心",
		"https://centerforsecuritypolicy.org/")

	w.SetStartingUrls([]string{
		"https://centerforsecuritypolicy.org/category/articles/",
		"https://centerforsecuritypolicy.org/category/books-and-reports/",
		"https://centerforsecuritypolicy.org/category/decision-briefs/",
		"https://centerforsecuritypolicy.org/category/in-the-news/",
		"https://centerforsecuritypolicy.org/category/press-release/",
		"https://centerforsecuritypolicy.org/category/situation-report/",
	})

	// 访问下一页 Index
	w.OnHTML(`a[class="next page-numbers"]`, func(element *colly.HTMLElement, ctx *Crawler.Context) {
		w.Visit(element.Attr("href"), Crawler.Index)
	})

	// 访问 News & Report 从 Index
	w.OnHTML(`[class="article-title article-title-1"] > a`, func(element *colly.HTMLElement, ctx *Crawler.Context) {
		url := element.Attr("href")
		if strings.Contains(ctx.Url, "/in-the-news/") || strings.Contains(ctx.Url, "/press-release/") {
			w.Visit(url, Crawler.News)
		} else {
			w.Visit(url, Crawler.Report)
		}
	})

	// 获取 Title
	w.OnHTML(`.entry-header h1.entry-title`, func(element *colly.HTMLElement, ctx *Crawler.Context) {
		ctx.Title = strings.TrimSpace(element.Text)
	})

	// 获取 PublicationTime
	w.OnHTML(`[class="item-metadata posts-date"] > i`, func(element *colly.HTMLElement, ctx *Crawler.Context) {
		ctx.PublicationTime = strings.TrimSpace(element.Text)
	})

	// 获取 Authors
	w.OnHTML(`[class="item-metadata posts-author"] > a`, func(element *colly.HTMLElement, ctx *Crawler.Context) {
		ctx.Authors = append(ctx.Authors, strings.TrimSpace(element.Text))
	})

	// 获取 Content
	w.OnHTML(`.pf-content > article, .entry-content`, func(element *colly.HTMLElement, ctx *Crawler.Context) {
		ctx.Content = strings.TrimSpace(element.Text)
	})

	// 获取 Tags
	w.OnHTML(`.cat-links > .meta-category > a`, func(element *colly.HTMLElement, ctx *Crawler.Context) {
		ctx.Tags = append(ctx.Tags, strings.TrimSpace(element.Text))
	})

	// 获取 File
	w.OnHTML(`.pf-content a`, func(element *colly.HTMLElement, ctx *Crawler.Context) {
		fileUrl := element.Attr("href")
		if strings.Contains(fileUrl, ".pdf") {
			ctx.File = append(ctx.File, element.Attr("href"))
		}
	})
}

package wipc

import (
	"github.com/gocolly/colly/v2"
	"megaCrawler/Crawler"
	"strings"
)

func init() {
	w := Crawler.Register("wipc", "Women's International Peace Center", "https://wipc.org/")

	w.SetStartingUrls([]string{
		"https://wipc.org/blog/latest-news/",
		"https://wipc.org/blog/featured-stories/",
		"https://wipc.org/resources/reports/",
		"https://wipc.org/resources/policy-briefs/",
	})

	// 访问 Report 从 Index ***
	w.OnHTML(`.post-inner > .post-thumbnail > a`, func(element *colly.HTMLElement, ctx *Crawler.Context) {
		w.Visit(element.Attr("href"), Crawler.Report)
	})

	// 访问 Report 从 Index ***
	w.OnHTML(`div[class="eael-entry-overlay fade-in"] > a`, func(element *colly.HTMLElement, ctx *Crawler.Context) {
		w.Visit(element.Attr("href"), Crawler.Report)
	})

	// 获取 Title ***
	w.OnHTML(`h1[class="page-title typo-heading"]`, func(element *colly.HTMLElement, ctx *Crawler.Context) {
		ctx.Title = strings.TrimSpace(element.Text)
	})

	// 获取 CategoryText ***
	w.OnHTML(`span:nth-child(3) a[property="item"]`, func(element *colly.HTMLElement, ctx *Crawler.Context) {
		ctx.CategoryText = strings.TrimSpace(element.Text)
	})

	// 获取 Content ***
	w.OnHTML(`div.has-content-area`, func(element *colly.HTMLElement, ctx *Crawler.Context) {
		ctx.Content = strings.TrimSpace(element.ChildText("p, h2, h3"))
	})

	// 获取 Tags ***
	w.OnHTML(`.tags-links > a[rel="tag"]`, func(element *colly.HTMLElement, ctx *Crawler.Context) {
		tag := strings.Replace(element.Text, "#", "", -1)
		ctx.Tags = append(ctx.Tags, strings.TrimSpace(tag))
	})

	// 获取 File ***
	w.OnHTML(`a[download]`, func(element *colly.HTMLElement, ctx *Crawler.Context) {
		ctx.File = append(ctx.File, element.Attr("href"))
	})
}

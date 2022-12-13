package pipvtr_wordpress

import (
	"github.com/gocolly/colly/v2"
	"megaCrawler/Crawler"
	"strings"
)

func init() {
	w := Crawler.Register("pipvtr_wordpress", "安全与国际研究所", "https://pipvtr.wordpress.com/")

	w.SetStartingUrls([]string{
		"https://pipvtr.wordpress.com/sitemap.xml",
	})

	// 访问 Report 从 sitemap
	w.OnXML(`//loc`, func(element *colly.XMLElement, ctx *Crawler.Context) {
		w.Visit(element.Text, Crawler.Report)
	})

	// 获取 Title
	w.OnHTML(`h1.entry-title`, func(element *colly.HTMLElement, ctx *Crawler.Context) {
		ctx.Title = strings.TrimSpace(element.Text)
	})

	// 获取 Title
	w.OnHTML(`h2.entry-title > a`, func(element *colly.HTMLElement, ctx *Crawler.Context) {
		ctx.Title = strings.TrimSpace(element.Text)
	})

	// 获取 PublicationTime
	w.OnHTML(`.entry-meta time[class="entry-date published"]`, func(element *colly.HTMLElement, ctx *Crawler.Context) {
		ctx.PublicationTime = strings.TrimSpace(element.Text)
	})

	// 获取 CategoryText
	w.OnHTML(`.entry-meta a[rel="category tag"]`, func(element *colly.HTMLElement, ctx *Crawler.Context) {
		ctx.CategoryText = strings.TrimSpace(element.Text)
	})

	// 获取 Authors
	w.OnHTML(`.entry-meta span[class="author vcard"] > a`, func(element *colly.HTMLElement, ctx *Crawler.Context) {
		ctx.Authors = append(ctx.Authors, strings.TrimSpace(element.Text))
	})

	// 获取 Content
	w.OnHTML(`div.entry-content`, func(element *colly.HTMLElement, ctx *Crawler.Context) {
		ctx.Content = strings.TrimSpace(element.Text)
	})
}

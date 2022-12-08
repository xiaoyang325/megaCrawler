package razumkov_org_ua

import (
	"github.com/gocolly/colly/v2"
	"megaCrawler/Crawler"
	"strings"
)

func init() {
	w := Crawler.Register("razumkov_org_ua", "拉祖姆科夫中心", "https://razumkov.org.ua/")

	w.SetStartingUrls([]string{
		"https://razumkov.org.ua/en/research-areas/economy",
		"https://razumkov.org.ua/en/research-areas/security",
		"https://razumkov.org.ua/en/research-areas/energy",
		"https://razumkov.org.ua/en/research-areas/foreign-policy",
		"https://razumkov.org.ua/en/research-areas/domestic-and-legal-policy",
		"https://razumkov.org.ua/en/research-areas/sotsialna-polityka",
		"https://razumkov.org.ua/en/sociology/press-releases",
	})

	// 访问下一页 Index
	w.OnHTML(`.pagination > li > a.next`, func(element *colly.HTMLElement, ctx *Crawler.Context) {
		w.Visit(element.Attr("href"), Crawler.Index)
	})

	// 访问 Report 从 Index
	w.OnHTML(`.tagItemView .tagItemTitle > a`, func(element *colly.HTMLElement, ctx *Crawler.Context) {
		w.Visit(element.Attr("href"), Crawler.Report)
	})

	// 获取 Title
	w.OnHTML(`.itemHeader > h2.itemTitle`, func(element *colly.HTMLElement, ctx *Crawler.Context) {
		ctx.Title = strings.TrimSpace(element.Text)
	})

	// 获取 PublicationTime
	w.OnHTML(`.itemHeader  span.itemDateCreated`, func(element *colly.HTMLElement, ctx *Crawler.Context) {
		ctx.PublicationTime = strings.TrimSpace(element.Text)
	})

	// 获取 CategoryText
	w.OnHTML(`.itemHeader .itemCategory > a`, func(element *colly.HTMLElement, ctx *Crawler.Context) {
		ctx.CategoryText = strings.TrimSpace(element.Text)
	})

	// 获取 Authors
	w.OnHTML(`.itemAuthor > a[rel="author"]`, func(element *colly.HTMLElement, ctx *Crawler.Context) {
		ctx.Authors = append(ctx.Authors, strings.TrimSpace(element.Text))
	})

	// 获取 Content
	w.OnHTML(`.itemBody > .itemFullText`, func(element *colly.HTMLElement, ctx *Crawler.Context) {
		ctx.Content = strings.TrimSpace(element.Text)
	})

	// 获取 Tags
	w.OnHTML(`ul.itemTags > li > a`, func(element *colly.HTMLElement, ctx *Crawler.Context) {
		ctx.Tags = append(ctx.Tags, strings.TrimSpace(element.Text))
	})

	// 获取 File
	w.OnHTML(`a[class="btn download"]`, func(element *colly.HTMLElement, ctx *Crawler.Context) {
		if strings.Contains(element.Attr("href"), ".pdf") {
			url := "https://razumkov.org.ua" + element.Attr("href")
			ctx.File = append(ctx.File, url)
		}
	})
}

package fau

import (
	"github.com/gocolly/colly/v2"
	"megaCrawler/Crawler"
	"strings"
)

func init() {
	w := Crawler.Register("fau", "皮尔斯堡大西洋大学", "https://www.fau.edu/")

	w.SetStartingUrls([]string{
		"https://www.fau.edu/hboi/newsroom/news/",
	})

	// 访问下一页 Index
	w.OnHTML(`.inline-nav > li:nth-last-child(1) > a.fau-pager__arrows`, func(element *colly.HTMLElement, ctx *Crawler.Context) {
		w.Visit(element.Attr("href"), Crawler.Index)
	})

	// 访问 News 从 Index
	w.OnHTML(`.blog-post > .title > a`, func(element *colly.HTMLElement, ctx *Crawler.Context) {
		w.Visit(element.Attr("href"), Crawler.News)
	})

	// 获取 Title
	w.OnHTML(`h1.mb-1`, func(element *colly.HTMLElement, ctx *Crawler.Context) {
		ctx.Title = strings.TrimSpace(element.Text)
	})

	// 获取 Title
	w.OnHTML(`.col-sm-8 > h1`, func(element *colly.HTMLElement, ctx *Crawler.Context) {
		ctx.Title = strings.TrimSpace(element.Text)
	})

	// 获取 PublicationTime
	w.OnHTML(`div.pb-4 > .toUpperCase:nth-child(3)`, func(element *colly.HTMLElement, ctx *Crawler.Context) {
		ctx.PublicationTime = strings.TrimSpace(element.Text)
	})

	// 获取 PublicationTime
	w.OnHTML(`div.row.spacer--bottom--20 > span`, func(element *colly.HTMLElement, ctx *Crawler.Context) {
		raw_str := element.Text
		raw_str = strings.Replace(raw_str, element.ChildText("a"), "", 1)
		raw_str = strings.Replace(raw_str, "By", "", 1)
		raw_str = strings.Replace(raw_str, "|", "", 1)
		ctx.PublicationTime = strings.TrimSpace(raw_str)
	})

	// 获取 Authors
	w.OnHTML(`div.pb-4 > .toUpperCase:nth-child(2)`, func(element *colly.HTMLElement, ctx *Crawler.Context) {
		name := strings.Replace(element.Text, "by", "", 1)
		name = strings.Replace(element.Text, "|", "", 1)
		ctx.Authors = append(ctx.Authors, strings.TrimSpace(name))
	})

	// 获取 Authors
	w.OnHTML(`div.row.spacer--bottom--20 > span > a`, func(element *colly.HTMLElement, ctx *Crawler.Context) {
		ctx.Authors = append(ctx.Authors, strings.TrimSpace(element.Text))
	})

	// 获取 Content
	w.OnHTML(`.section-m > section#section`, func(element *colly.HTMLElement, ctx *Crawler.Context) {
		ctx.Content = strings.TrimSpace(element.Text)
	})

	// 获取 Content
	w.OnHTML(`div.col-sm-8 `, func(element *colly.HTMLElement, ctx *Crawler.Context) {
		ctx.Content = strings.TrimSpace(element.ChildText("p"))
	})

	// 获取 Tags
	w.OnHTML(`div.col-sm-4 > p > a`, func(element *colly.HTMLElement, ctx *Crawler.Context) {
		ctx.Tags = append(ctx.Tags, strings.TrimSpace(element.Text))
	})
}

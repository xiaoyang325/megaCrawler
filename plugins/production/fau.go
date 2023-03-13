package production

import (
	"strings"

	"megaCrawler/crawlers"

	"github.com/gocolly/colly/v2"
)

func init() {
	w := crawlers.Register("fau", "皮尔斯堡大西洋大学", "https://www.fau.edu/")

	w.SetStartingURLs([]string{
		"https://www.fau.edu/hboi/newsroom/news/",
	})

	// 访问下一页 Index
	w.OnHTML(`.inline-nav > li:nth-last-child(1) > a.fau-pager__arrows`, func(element *colly.HTMLElement, ctx *crawlers.Context) {
		w.Visit(element.Attr("href"), crawlers.Index)
	})

	// 访问 News 从 Index
	w.OnHTML(`.blog-post > .title > a`, func(element *colly.HTMLElement, ctx *crawlers.Context) {
		w.Visit(element.Attr("href"), crawlers.News)
	})

	// 获取 Title
	w.OnHTML(`h1.mb-1`, func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Title = strings.TrimSpace(element.Text)
	})

	// 获取 Title
	w.OnHTML(`.col-sm-8 > h1`, func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Title = strings.TrimSpace(element.Text)
	})

	// 获取 PublicationTime
	w.OnHTML(`div.pb-4 > .toUpperCase:nth-child(3)`, func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.PublicationTime = strings.TrimSpace(element.Text)
	})

	// 获取 PublicationTime
	w.OnHTML(`div.row.spacer--bottom--20 > span`, func(element *colly.HTMLElement, ctx *crawlers.Context) {
		rawStr := element.Text
		rawStr = strings.Replace(rawStr, element.ChildText("a"), "", 1)
		rawStr = strings.Replace(rawStr, "By", "", 1)
		rawStr = strings.Replace(rawStr, "|", "", 1)
		ctx.PublicationTime = strings.TrimSpace(rawStr)
	})

	// 获取 Authors
	w.OnHTML(`div.pb-4 > .toUpperCase:nth-child(2)`, func(element *colly.HTMLElement, ctx *crawlers.Context) {
		name := strings.Replace(element.Text, "by", "", 1)
		name = strings.Replace(name, "|", "", 1)
		ctx.Authors = append(ctx.Authors, strings.TrimSpace(name))
	})

	// 获取 Authors
	w.OnHTML(`div.row.spacer--bottom--20 > span > a`, func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Authors = append(ctx.Authors, strings.TrimSpace(element.Text))
	})

	// 获取 Content
	w.OnHTML(`.section-m > section#section`, func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Content = strings.TrimSpace(element.Text)
	})

	// 获取 Content
	w.OnHTML(`div.col-sm-8 `, func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Content = strings.TrimSpace(element.ChildText("p"))
	})

	// 获取 Tags
	w.OnHTML(`div.col-sm-4 > p > a`, func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Tags = append(ctx.Tags, strings.TrimSpace(element.Text))
	})
}

package production

import (
	"megaCrawler/crawlers"
	"strings"

	"github.com/gocolly/colly/v2"
)

func init() {
	w := crawlers.Register("cejil", "国际法与司法中心", "https://cejil.org/")

	w.SetStartingUrls([]string{
		"https://cejil.org/en/press-releases/",
		"https://cejil.org/en/publications/",
		"https://cejil.org/en/blog/",
	})

	// 访问下一页 Index
	w.OnHTML(`a[class="next page-numbers"]`, func(element *colly.HTMLElement, ctx *crawlers.Context) {
		w.Visit(element.Attr("href"), crawlers.Index)
	})

	// 访问 News & Report 从 Index
	w.OnHTML(`.brd_btm > div.txt-rel > h3 > a`, func(element *colly.HTMLElement, ctx *crawlers.Context) {
		if strings.Contains(ctx.Url, "/press-releases/") {
			w.Visit(element.Attr("href"), crawlers.News)
		} else {
			w.Visit(element.Attr("href"), crawlers.Report)
		}
	})

	// 获取 Title
	w.OnHTML(`article > div.box-content > h1`, func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Title = strings.TrimSpace(element.Text)
	})

	// 获取 PublicationTime
	w.OnHTML(`article > div.box-content > .date`, func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.PublicationTime = strings.TrimSpace(element.Text)
	})

	// 获取 Content
	w.OnHTML(`article > div.box-content`, func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Content = strings.TrimSpace(element.ChildText("p"))
	})

	// 获取 Tags
	w.OnHTML(`article > div.box-content > .tags`, func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Tags = append(ctx.Tags, strings.TrimSpace(element.Text))
	})

	// 获取 File
	w.OnHTML(`[class="box-widget download"] a`, func(element *colly.HTMLElement, ctx *crawlers.Context) {
		if strings.Contains(element.Attr("href"), ".pdf") {
			ctx.File = append(ctx.File, element.Attr("href"))
		}
	})
}

package production

import (
	"strings"

	"megaCrawler/crawlers"

	"github.com/gocolly/colly/v2"
)

func init() {
	w := crawlers.Register("demandprogress", "Demand Progress",
		"https://demandprogress.org/")

	w.SetStartingURLs([]string{
		"https://demandprogress.org/media/press-releases/",
		"https://demandprogress.org/policy-work/",
	})

	// 访问下一页 Index
	w.OnHTML(`a.nextpostslink`, func(element *colly.HTMLElement, ctx *crawlers.Context) {
		w.Visit(element.Attr("href"), crawlers.Index)
	})

	// 访问 News 从 Index
	w.OnHTML(`article > div > div > header > h2 > a`, func(element *colly.HTMLElement, ctx *crawlers.Context) {
		if !strings.Contains(element.Attr("href"), ".pdf") {
			w.Visit(element.Attr("href"), crawlers.News)
		}
	})

	// 添加 Report 从 Index 通过 SubContext
	w.OnHTML(`div.item-content`, func(element *colly.HTMLElement, ctx *crawlers.Context) {
		url := element.ChildAttr(`h2[class="entry-title nudge-hover"] > a`, "href")

		if strings.Contains(url, ".pdf") {
			subCtx := ctx.CreateSubContext()
			subCtx.PageType = crawlers.Report
			subCtx.Title = strings.TrimSpace(element.ChildText(`h2[class="entry-title nudge-hover"] > a`))
			subCtx.File = append(subCtx.File, url)
			subCtx.CategoryText = strings.TrimSpace(element.ChildText(".post-type-name  > a"))
		}
	})

	// 获取 Title
	w.OnHTML(`h1.entry-title`, func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Title = strings.TrimSpace(element.Text)
	})

	// 获取 PublicationTime
	w.OnHTML(`.header-content > time`, func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.PublicationTime = strings.TrimSpace(element.Text)
	})

	// 获取 CategoryText
	w.OnHTML(`.post-type-name > span`, func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.CategoryText = strings.TrimSpace(element.Text)
	})

	// 获取 Content
	w.OnHTML(`div[class="entry-content stop-gradient"]`, func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Content = strings.TrimSpace(element.ChildText("p, td"))
	})
}

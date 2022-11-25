package demandprogress

import (
	"github.com/gocolly/colly/v2"
	"megaCrawler/Crawler"
	"strings"
)

func init() {
	w := Crawler.Register("demandprogress", "Demand Progress",
			"https://demandprogress.org/")
	
	w.SetStartingUrls([]string{
		"https://demandprogress.org/media/press-releases/",
		"https://demandprogress.org/policy-work/",
	})

	// 访问下一页 Index 
	w.OnHTML(`a.nextpostslink`,
		func(element *colly.HTMLElement, ctx *Crawler.Context) {
			w.Visit(element.Attr("href"), Crawler.Index)
		})

	// 访问 News 从 Index 
	w.OnHTML(`article > div > div > header > h2 > a`,
		func(element *colly.HTMLElement, ctx *Crawler.Context) {
			if !strings.Contains(element.Attr("href"), ".pdf") {
				w.Visit(element.Attr("href"), Crawler.News)
			}
		})

	// 添加 Report 从 Index 通过 SubContext
	w.OnHTML(`div.item-content`,
		func(element *colly.HTMLElement, ctx *Crawler.Context) {
			url := element.ChildAttr(`h2[class="entry-title nudge-hover"] > a`, "href")

			if strings.Contains(url, ".pdf") {
				sub_ctx := ctx.CreateSubContext()
				sub_ctx.PageType = Crawler.Report
				sub_ctx.Title = strings.TrimSpace(element.ChildText(`h2[class="entry-title nudge-hover"] > a`))
				sub_ctx.File = append(sub_ctx.File, url)
				sub_ctx.CategoryText = strings.TrimSpace(element.ChildText(".post-type-name  > a"))
			}
		})

	// 获取 Title 
	w.OnHTML(`h1.entry-title`,
		func(element *colly.HTMLElement, ctx *Crawler.Context) {
			ctx.Title = strings.TrimSpace(element.Text)
		})

	// 获取 PublicationTime 
	w.OnHTML(`.header-content > time`,
		func(element *colly.HTMLElement, ctx *Crawler.Context) {
			ctx.PublicationTime = strings.TrimSpace(element.Text)
		})

	// 获取 CategoryText 
	w.OnHTML(`.post-type-name > span`,
		func(element *colly.HTMLElement, ctx *Crawler.Context) {
			ctx.CategoryText = strings.TrimSpace(element.Text)
		})

	// 获取 Content 
	w.OnHTML(`div[class="entry-content stop-gradient"]`,
		func(element *colly.HTMLElement, ctx *Crawler.Context) {
			ctx.Content = strings.TrimSpace(element.ChildText("p"))
		})
}

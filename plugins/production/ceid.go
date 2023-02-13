package production

import (
	"megaCrawler/crawlers"

	"github.com/gocolly/colly/v2"
)

func init() {
	w := crawlers.Register("ceid", "新型传染病政策与研究中心", "https://www.bu.edu/ceid/")
	w.SetStartingUrls([]string{"https://www.bu.edu/ceid/news-events/in-the-media/"})

	// index
	w.OnHTML(".nav-next > a", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		w.Visit(element.Attr("href"), crawlers.Index)
	})

	// 访问新闻
	w.OnHTML("a.more", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		w.Visit(element.Attr("href"), crawlers.News)
	})

	// 新闻标题
	w.OnHTML("article > h2", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Title = element.Text
	})

	// 新闻正文
	w.OnHTML("article > p", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Content += element.Text
	})
}

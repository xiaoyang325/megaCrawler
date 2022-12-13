package impact

import (
	"github.com/gocolly/colly/v2"
	"megaCrawler/Crawler"
)

func init() {
	w := Crawler.Register("impact", "影响新闻", "https://impactnews.uk/")
	w.SetStartingUrls([]string{"https://impactnews.uk/news/"})

	//index
	w.OnHTML("div.alignleft > a", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		w.Visit(element.Attr("href"), Crawler.Index)
	})

	//访问新闻
	w.OnHTML("h2.entry-title > a", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		w.Visit(element.Attr("href"), Crawler.News)
	})

	//标题
	w.OnHTML("h1.et_pb_module_header", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		if ctx.PageType == Crawler.News {
			ctx.Title = element.Text
		}
	})

	//正文
	w.OnHTML("div > p", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		if ctx.PageType == Crawler.News {
			ctx.Content += element.Text
		}
	})
}

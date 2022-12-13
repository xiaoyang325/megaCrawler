package globaltimes

import (
	"github.com/gocolly/colly/v2"
	"megaCrawler/Crawler"
)

func init() {
	w := Crawler.Register("globaltimes", "环球时报", "https://www.globaltimes.cn/")
	w.SetStartingUrls([]string{"https://www.globaltimes.cn/china/index.html",
		"https://www.globaltimes.cn/source/index.html",
		"https://www.globaltimes.cn/opinion/index.html",
		"https://www.globaltimes.cn/In-depth/index.html",
		"https://www.globaltimes.cn/world/index.html",
		"https://www.globaltimes.cn/life/index.html",
		"https://www.globaltimes.cn/sport/index.html"})

	//index
	w.OnHTML("body > div:nth-child(7) > div > div > nav > ul > li > a", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		w.Visit(element.Attr("href"), Crawler.Index)
	})

	//news
	w.OnHTML("a.new_title_ms", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		w.Visit(element.Attr("href"), Crawler.News)
	})

	//标题
	w.OnHTML("div.article_title", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		ctx.Title = element.Text
	})

	//正文
	w.OnHTML("div.article_content", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		ctx.Content = element.Text
	})
}

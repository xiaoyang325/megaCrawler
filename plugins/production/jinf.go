package production

import (
	"megaCrawler/crawlers"

	"github.com/gocolly/colly/v2"
)

func init() {
	w := crawlers.Register("jinf", "国家基本问题研究所", "https://jinf.jp/")
	w.SetStartingUrls([]string{"https://jinf.jp/meeting", "https://jinf.jp/weekly", "https://jinf.jp/suggestion"})

	// index
	w.OnHTML("#pageNavWrapper > ul > li> a", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		w.Visit(element.Attr("href"), crawlers.Index)
	})

	// 访问news
	w.OnHTML("#articleAreaInner > div > div.title > h3 > a", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		w.Visit(element.Attr("href"), crawlers.News)
	})

	// 访问report
	w.OnHTML("#articleAreaInner > div> div.titleArea > div > div > h3 > a", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		w.Visit(element.Attr("href"), crawlers.Report)
	})

	// 获取标题
	w.OnHTML(" div > div> div > div.title > h3", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Title = element.Text
	})

	// 获取正文
	w.OnHTML(".text", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Content = element.Text
	})

	// 获取pdf
	w.OnHTML(" div > div.clearfix > div > div.text > div.box5 > p > a", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.File = append(ctx.File, element.Attr("href"))
	})
}

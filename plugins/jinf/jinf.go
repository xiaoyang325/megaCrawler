package jinf

import (
	"github.com/gocolly/colly/v2"
	"megaCrawler/Crawler"
)

func init() {
	w := Crawler.Register("jinf", "国家基本问题研究所", "https://jinf.jp/")
	w.SetStartingUrls([]string{"https://jinf.jp/meeting", "https://jinf.jp/weekly", "https://jinf.jp/suggestion"})

	//index
	w.OnHTML("#pageNavWrapper > ul > li> a", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		w.Visit(element.Attr("href"), Crawler.Index)
	})

	//访问news
	w.OnHTML("#articleAreaInner > div > div.title > h3 > a", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		w.Visit(element.Attr("href"), Crawler.News)
	})

	//访问report
	w.OnHTML("#articleAreaInner > div> div.titleArea > div > div > h3 > a", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		w.Visit(element.Attr("href"), Crawler.Report)
	})

	//获取标题
	w.OnHTML(" div > div> div > div.title > h3", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		ctx.Title = element.Text
	})

	//获取正文
	w.OnHTML(".text", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		ctx.Content = element.Text
	})

	//获取pdf
	w.OnHTML(" div > div.clearfix > div > div.text > div.box5 > p > a", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		ctx.File = append(ctx.File, element.Attr("href"))
	})
}

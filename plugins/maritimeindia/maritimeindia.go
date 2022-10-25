package maritimeindia

import (
	"github.com/gocolly/colly/v2"
	"megaCrawler/megaCrawler"
)

func init() {
	w := megaCrawler.Register("maritimeindia", "国家海事基金会", "https://maritimeindia.org/")

	w.SetStartingUrls([]string{"https://maritimeindia.org/category/articles-nmf/",
		"https://maritimeindia.org/category/holistic-maritime-security-thematic/",
		"https://maritimeindia.org/category/maritime_technology/",
		"https://maritimeindia.org/category/blue-economy-climate-change/",
		"https://maritimeindia.org/category/oceanic-resources/",
		"https://maritimeindia.org/category/maritime-law/",
		"https://maritimeindia.org/category/environment-issues/",
		"https://maritimeindia.org/category/maritime-trade-connectivity/",
		"https://maritimeindia.org/category/maritime-energy/",
		"https://maritimeindia.org/category/maritime-history-and-culture/",
		"https://maritimeindia.org/category/maritime-safety/",
		"https://maritimeindia.org/previous-events/",
		"https://maritimeindia.org/making-waves/"})

	//index
	w.OnHTML("a.inactive", func(element *colly.HTMLElement, ctx *megaCrawler.Context) {
		w.Visit(element.Attr("href"), megaCrawler.Index)
	})

	//访问报告
	w.OnHTML("h3.slide-entry-title.entry-title", func(element *colly.HTMLElement, ctx *megaCrawler.Context) {
		w.Visit(element.Attr("href"), megaCrawler.Report)
	})

	w.OnHTML("figure > div ", func(element *colly.HTMLElement, ctx *megaCrawler.Context) {
		w.Visit(element.Attr("href"), megaCrawler.Report)
	})

	//获取报告标题
	w.OnHTML("h1 > a", func(element *colly.HTMLElement, ctx *megaCrawler.Context) {
		ctx.Title = element.Text
	})

	//获取作者
	w.OnHTML("div.editor-right", func(element *colly.HTMLElement, ctx *megaCrawler.Context) {
		ctx.Authors = append(ctx.Authors, element.Text)
	})

	//报告关键词
	w.OnHTML(" time > span", func(element *colly.HTMLElement, ctx *megaCrawler.Context) {
		ctx.Keywords = append(ctx.Keywords, element.Text)
	})

	//报告正文
	w.OnHTML("div.entry-content", func(element *colly.HTMLElement, ctx *megaCrawler.Context) {
		ctx.Content = element.Text
	})

	//pdf
	w.OnHTML(".avia-builder-el-no-sibling > a", func(element *colly.HTMLElement, ctx *megaCrawler.Context) {
		ctx.File = append(ctx.File, element.Attr("href"))
	})

	//访问新闻
	w.OnHTML(" div.mec-wrap > div > article", func(element *colly.HTMLElement, ctx *megaCrawler.Context) {
		w.Visit(element.Attr("href"), megaCrawler.News)
	})

	w.OnHTML(" div.mec-event-content > h1", func(element *colly.HTMLElement, ctx *megaCrawler.Context) {
		ctx.Title = element.Text
	})

	w.OnHTML(".col-md-8 > div.mec-event-content > div", func(element *colly.HTMLElement, ctx *megaCrawler.Context) {
		ctx.Content = element.Text
	})

}

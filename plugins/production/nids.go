package production

import (
	"strings"

	"megaCrawler/crawlers"

	"github.com/gocolly/colly/v2"
)

func init() {
	w := crawlers.Register("nids", "防卫研究所", "http://www.nids.mod.go.jp/index.html")
	w.SetStartingURLs([]string{"http://www.nids.mod.go.jp/research/profile/index.html",
		"http://www.nids.mod.go.jp/publication/index.html"})

	// index
	w.OnHTML("#profile > ul > li > a", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		url, err := element.Request.URL.Parse(element.Attr("href"))
		if err != nil {
			crawlers.Sugar.Error(err.Error())
			return
		}
		w.Visit(url.String(), crawlers.Index)
	})

	// 访问人物
	w.OnHTML("div> table > tbody > tr > td> a", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		url, err := element.Request.URL.Parse(element.Attr("href"))
		if err != nil {
			crawlers.Sugar.Error(err.Error(), element.Request.URL.String())
			return
		}
		w.Visit(url.String(), crawlers.Expert)
	})

	// 获取人物姓名
	w.OnHTML("p.name.mtx", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Name = element.Text
	})

	// 人物领域
	w.OnHTML("#content > div:nth-child(4) > p:nth-child(2)", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Area = element.Text
	})

	// 人物简介
	w.OnHTML("#content > div:nth-child(3) > div > div > table> tbody > tr", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Description = element.Text
	})

	// index(report)
	w.OnHTML("#localnavi > ul > li> a", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		url, err := element.Request.URL.Parse(element.Attr("href"))
		if err != nil {
			crawlers.Sugar.Error(err.Error())
			return
		}
		w.Visit(url.String(), crawlers.Index)
	})
	w.OnHTML("#content > div.section > ul > li> a", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		url, err := element.Request.URL.Parse(element.Attr("href"))
		if err != nil {
			crawlers.Sugar.Error(err.Error())
			return
		}
		w.Visit(url.String(), crawlers.Index)
	})
	w.OnHTML("#content > div.section.mtx > div > div > ul > li> a", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		url, err := element.Request.URL.Parse(element.Attr("href"))
		if err != nil {
			crawlers.Sugar.Error(err.Error())
			return
		}
		w.Visit(url.String(), crawlers.Index)
	})

	// 尝试寻找下载pdf的按钮，并如果存在则将页面类型转换为报告
	w.OnHTML("a", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		if strings.Contains(element.Attr("href"), ".pdf") && (ctx.PageType != crawlers.Expert) {
			ctx.File = append(ctx.File, element.Attr("href"))
			ctx.PageType = crawlers.Report
		}
	})

	w.OnHTML("h1", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Title = element.Text
	})
}

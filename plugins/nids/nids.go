package nids

import (
	"github.com/gocolly/colly/v2"
	"megaCrawler/Crawler"
	"strings"
)

func init() {
	w := Crawler.Register("nids", "防卫研究所", "http://www.nids.mod.go.jp/index.html")
	w.SetStartingUrls([]string{"http://www.nids.mod.go.jp/research/profile/index.html",
		"http://www.nids.mod.go.jp/publication/index.html"})

	//index
	w.OnHTML("#profile > ul > li > a", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		url, _ := element.Request.URL.Parse(element.Attr("href"))
		w.Visit(url.String(), Crawler.Index)
	})

	//访问人物
	w.OnHTML(" div> table > tbody > tr > td> a", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		url, _ := element.Request.URL.Parse(element.Attr("href"))
		w.Visit(url.String(), Crawler.Expert)
	})

	//获取人物姓名
	w.OnHTML("p.name.mtx", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		ctx.Name = element.Text
	})

	//人物领域
	w.OnHTML("#content > div:nth-child(4) > p:nth-child(2)", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		ctx.Area = element.Text
	})

	//人物简介
	w.OnHTML("#content > div:nth-child(3) > div > div > table> tbody > tr", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		ctx.Description = element.Text
	})

	//index(report)
	w.OnHTML("#localnavi > ul > li> a", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		url, _ := element.Request.URL.Parse(element.Attr("href"))
		w.Visit(url.String(), Crawler.Index)
	})
	w.OnHTML("#content > div.section > ul > li> a", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		url, _ := element.Request.URL.Parse(element.Attr("href"))
		w.Visit(url.String(), Crawler.Index)
	})
	w.OnHTML("#content > div.section.mtx > div > div > ul > li> a", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		url, _ := element.Request.URL.Parse(element.Attr("href"))
		w.Visit(url.String(), Crawler.Index)
	})

	// 尝试寻找下载pdf的按钮，并如果存在则将页面类型转换为报告
	w.OnHTML("li > a", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		if strings.Contains(element.Attr("href"), ".pdf") && (ctx.PageType != Crawler.Expert) {
			ctx.File = append(ctx.File, element.Attr("href"))
			ctx.PageType = Crawler.Report
		}
	})
	w.OnHTML("div > a", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		if strings.Contains(element.Attr("href"), ".pdf") && (ctx.PageType != Crawler.Expert) {
			ctx.File = append(ctx.File, element.Attr("href"))
			ctx.PageType = Crawler.Report
		}
	})
	w.OnHTML("p > a", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		if strings.Contains(element.Attr("href"), ".pdf") && (ctx.PageType != Crawler.Expert) {
			ctx.File = append(ctx.File, element.Attr("href"))
			ctx.PageType = Crawler.Report
		}
	})
}

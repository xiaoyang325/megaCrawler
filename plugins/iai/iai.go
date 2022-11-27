package iai

import (
	"github.com/gocolly/colly/v2"
	"megaCrawler/Crawler"
	"strings"
)

func init() {
	w := Crawler.Register("iai", "国际事务研究所", "https://www.iai.it/")

	w.SetStartingUrls([]string{"https://www.iai.it/en/tema/africa",
		"https://www.iai.it/en/tema/399",
		"https://www.iai.it/en/tema/13",
		"https://www.iai.it/en/tema/9",
		"https://www.iai.it/en/tema/299",
		"https://www.iai.it/en/tema/11",
		"https://www.iai.it/en/tema/15",
		"https://www.iai.it/en/tema/12",
		"https://www.iai.it/en/tema/500",
		"https://www.iai.it/en/tema/8",
		"https://www.iai.it/en/tema/455",
		"https://www.iai.it/en/tema/445",
		"https://www.iai.it/en/tema/14",
		"https://www.iai.it/en/tema/462",
		"https://www.iai.it/en/pubblicazioni/lista/all/all"})

	// 尝试寻找下载pdf的按钮，并如果存在则将页面类型转换为报告
	w.OnHTML("a.button", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		if strings.Contains(element.Attr("href"), ".pdf") {
			ctx.File = append(ctx.File, element.Attr("href"))
			ctx.PageType = Crawler.Report
		}
	})

	// 从翻页器获取链接并访问
	w.OnHTML("div.more-link>a", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		w.Visit(element.Attr("href"), Crawler.Index)
	})
	// 从翻页器获取链接并访问
	w.OnHTML("ul.pagination>li>a", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		w.Visit(element.Attr("href"), Crawler.Index)
	})

	// 从index访问新闻
	w.OnHTML("div.field-title-field>a", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		w.Visit(element.Attr("href"), Crawler.News)
	})
	w.OnHTML(".tit>a", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		w.Visit(element.Attr("href"), Crawler.News)
	})
	//内含Expert
	w.OnHTML("div.esperti>ul>li>div>h3>a", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		w.Visit(element.Attr("href"), Crawler.Expert)
	})

	// report.title
	w.OnHTML("h1.page-title", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		ctx.Title = element.Text
	})

	// report.author
	w.OnHTML(".field-pub-autori>a", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		ctx.Authors = append(ctx.Authors, element.Text)
	})
	w.OnHTML("div.field-ricerca-autore>a", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		ctx.Authors = append(ctx.Authors, element.Text)
	})
	//内含Expert
	w.OnHTML(".field-pub-autori>a", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		w.Visit(element.Attr("href"), Crawler.Expert)
	})
	w.OnHTML("div.field-ricerca-autore>a", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		w.Visit(element.Attr("href"), Crawler.Expert)
	})

	//report.publish time
	w.OnHTML("span.date-display-single", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		ctx.PublicationTime = element.Text
	})
	w.OnHTML("div.data-r", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		ctx.PublicationTime = element.Text
	})
	// report .content
	w.OnHTML("div.field-body", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		ctx.Content = element.Text
	})
	// expert.Name
	w.OnHTML(" h1.page-header", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		ctx.Name = element.Text
	})
	// expert.title
	w.OnHTML("div.field-autore-qualifica", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		ctx.Title = element.Text
	})
	// expert.description
	w.OnHTML("div.field-body", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		ctx.Content = element.Text
	})
	// expert.link
	w.OnHTML("div.riga-social>span>a", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		ctx.Link = append(ctx.Link, element.Attr("href"))
	})
	// expert.area
	w.OnHTML("div.field-pub-keywords", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		ctx.Area = ctx.Area + "," + element.Text
	})

}

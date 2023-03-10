package production

import (
	"github.com/gocolly/colly/v2"
	"megaCrawler/crawlers"
	"strings"
)

func init() {
	w := crawlers.Register("bruegel", "布鲁盖尔研究所", "https://www.bruegel.org/")

	w.SetStartingURLs([]string{"https://www.bruegel.org/topics/banking-and-capital-markets",
		"https://www.bruegel.org/topics/digital-economy-and-innovation",
		"https://www.bruegel.org/topics/european-governance",
		"https://www.bruegel.org/topics/global-economy-and-trade",
		"https://www.bruegel.org/topics/green-economy",
		"https://www.bruegel.org/topics/inclusive-economy",
		"https://www.bruegel.org/topics/macroeconomic-policies",
		"https://www.bruegel.org/publications/datasets",
		"https://www.bruegel.org/publications/testimonies",
		"https://www.bruegel.org/publications/external-publications",
		"https://www.bruegel.org/publications/external-opinion",
		"https://www.bruegel.org/bruegel-blog",
		"https://www.bruegel.org/events/past-events",
		"https://www.bruegel.org/podcast-series/future-work",
		"https://www.bruegel.org/event-series/energy-crisis",
		"https://www.bruegel.org/event-series/ukraine",
		"https://www.bruegel.org/event-series/china"})

	// 尝试寻找下载pdf的按钮，并如果存在则将页面类型转换为报告
	w.OnHTML("a.button", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		if strings.Contains(element.Attr("href"), ".pdf") {
			ctx.File = append(ctx.File, element.Attr("href"))
			ctx.PageType = crawlers.Report
		}
	})

	// 从翻页器获取链接并访问
	w.OnHTML("a.o-link", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		w.Visit(element.Attr("href"), crawlers.Index)
	})
	w.OnHTML("a.c-pager__button-link", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		w.Visit(element.Attr("href"), crawlers.Index)
	})

	// 从index访问新闻
	w.OnHTML("h3.c-card__title>a", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		w.Visit(element.Attr("href"), crawlers.News)
	})
	w.OnHTML("h2.c-list-item__title>a", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		w.Visit(element.Attr("href"), crawlers.News)
	})

	// report.title
	w.OnHTML("#main-content > article > header > div > div > h1 > span > font > font", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Title = element.Text
	})
	w.OnHTML("h1.c-single-header__title>span", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Title = element.Text
	})

	// report.description
	w.OnHTML("p.c-single-header__description>font>font", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Description = element.Text
	})
	w.OnHTML("p.c-single-header__description", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Description = element.Text
	})
	// report.publish time
	w.OnHTML(" dd.c-single-header__meta-term > font > font", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.PublicationTime = element.Text
	})
	w.OnHTML(" #main-content > article > header > div > div > dl > div:nth-child(1) > dd ", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.PublicationTime = element.Text
	})

	// report.author
	w.OnHTML("dd.c-single-header__meta-term>a>font>font", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Authors = append(ctx.Authors, element.Text)
	})
	w.OnHTML("#main-content > article > header > div > div > dl > div:nth-child(2) > dd > a", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Authors = append(ctx.Authors, element.Text)
	})

	// report .content
	w.OnHTML("div.c-wysiwyg__inner", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Content = element.Text
	})
	// report.category
	w.OnHTML("small.c-single-header__label", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.CategoryText = element.Text
	})
}

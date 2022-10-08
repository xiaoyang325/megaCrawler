package cato

import (
	"github.com/gocolly/colly/v2"
	"megaCrawler/megaCrawler"
	"strings"
)

var w *megaCrawler.WebsiteEngine

func handler(element *colly.HTMLElement, ctx *megaCrawler.Context) {
	w.Visit(element.Attr("href"), megaCrawler.Expert)
}

func init() {
	w := megaCrawler.Register("cato", "卡托研究所", "https://www.csis.org/")

	w.SetStartingUrls([]string{"https://www.csis.org/experts"})

	w.OnResponse(func(response *colly.Response, ctx *megaCrawler.Context) {
		println(response.Request.URL.String())
	})

	w.OnHTML("a.btn", func(element *colly.HTMLElement, ctx *megaCrawler.Context) {
		if strings.Contains(element.Attr("href"), ".pdf") {
			ctx.File = append(ctx.File, element.Attr("href"))
		}
	})

	w.OnHTML(".pager__link", func(element *colly.HTMLElement, ctx *megaCrawler.Context) {
		w.Visit(element.Attr("href"), megaCrawler.Index)
	})

	w.OnHTML(".teaser--portrait-image > .teaser__title > a", handler)

	w.OnHTML("h1", func(element *colly.HTMLElement, ctx *megaCrawler.Context) {
		if ctx.PageType == megaCrawler.Expert {
			ctx.Name = element.Text
		} else if ctx.PageType == megaCrawler.Report || ctx.PageType == megaCrawler.News {
			ctx.Title = element.Text
		}
	})

	w.OnHTML("#main > div > div.layout-constrain > div.layout-focus-page__main > div.subtitle", func(element *colly.HTMLElement, ctx *megaCrawler.Context) {
		ctx.Title = megaCrawler.StandardizeSpaces(element.Text)
	})

	w.OnHTML("#main > div > div.layout-constrain > div.layout-focus-page__main > p", func(element *colly.HTMLElement, ctx *megaCrawler.Context) {
		ctx.Description = element.Text
	})

	w.OnHTML(".layout-focus-page__main > .field field--inline field--spaced > a", func(element *colly.HTMLElement, ctx *megaCrawler.Context) {
		ctx.Area += element.Text + " "
	})
}
